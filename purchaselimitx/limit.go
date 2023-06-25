package purchaselimitx

import (
	"context"
	"fmt"
	"time"

	"github.com/fengde/gocommon/errorx"
	"github.com/fengde/gocommon/jsonx"
	"github.com/fengde/gocommon/logx"
	"github.com/fengde/gocommon/storex/redisx"
)

type PurchaseLimitType uint64

const (
	PurchaseLimitTypeNone     PurchaseLimitType = 1
	PurchaseLimitTypeDay      PurchaseLimitType = 2
	PurchaseLimitTypeWeek     PurchaseLimitType = 3
	PurchaseLimitTypeLasting  PurchaseLimitType = 4
	PurchaseLimitTypeDuration PurchaseLimitType = 5
)

type IPurchaseLimit interface {
	Check(ctx context.Context, actityId string) error
	GetUseCount(ctx context.Context, actityId string) (uint64, error)
}

type PurchaseLimitCache struct {
	Start time.Time `json:"start"`
	Count uint64    `json:"count"`
}

func (p *PurchaseLimitCache) Map() map[string]interface{} {
	return map[string]interface{}{
		"start": p.Start,
		"count": p.Count,
	}
}

// 场景
// 用户按天限购，指定时间点刷新额度（比如指定每天6点刷新）
// 用户按周限购，指定时间点刷新额度（比如指定每周一6点刷新）
// 用户时长限购，时长过期到点刷新
type PurchaseLimit struct {
	limit     uint64
	typ       PurchaseLimitType
	at        string
	durateion int64
	expire    int64
	cli       *redisx.Client
}

// NewPurchaseLimitDay 按天限购， at指定刷新时间点：如 "06:00:00"
func NewPurchaseLimitDay(limit uint64, at string, cli *redisx.Client) IPurchaseLimit {
	if len(at) != 8 {
		panic("at is invalid")
	}
	return &PurchaseLimit{
		limit:  limit,
		typ:    PurchaseLimitTypeDay,
		at:     at,
		expire: 86400 * 2,
		cli:    cli,
	}
}

// NewPurchaseLimitWeek 按周限购， at指定周一刷新时间点：如 "06:00:00"
func NewPurchaseLimitWeek(limit uint64, at string, cli *redisx.Client) IPurchaseLimit {
	if len(at) != 8 {
		panic("at is invalid")
	}
	return &PurchaseLimit{
		limit:  limit,
		typ:    PurchaseLimitTypeWeek,
		at:     at,
		expire: 86400 * 14,
		cli:    cli,
	}
}

// NewPurchaseLimitDuration 按时长限购，second指定多久之后刷新
func NewPurchaseLimitDuration(limit uint64, second int64, cli *redisx.Client) IPurchaseLimit {
	return &PurchaseLimit{
		limit:     limit,
		typ:       PurchaseLimitTypeDuration,
		durateion: second,
		expire:    second * 2,
		cli:       cli,
	}
}

// NewPurchaseLimitLasting 固定限制次数，永不过期
func NewPurchaseLimitLasting(limit uint64, cli *redisx.Client) *PurchaseLimit {
	return &PurchaseLimit{
		limit:  limit,
		typ:    PurchaseLimitTypeLasting,
		expire: 0,
		cli:    cli,
	}
}

// 限购检查
func (p *PurchaseLimit) Check(ctx context.Context, actityId string) error {
	now := time.Now()

	key := p.key(ctx, actityId)
	pc, err := p.getCache(ctx, key)
	if err != nil {
		return err
	}
	if pc == nil {
		return p.flushCache(ctx, key, PurchaseLimitCache{
			Start: now,
			Count: 1,
		})
	}

	logx.DebugWithCtx(ctx, actityId, pc)

	if p.judgeWindow(pc.Start, now) {
		// 是否限购，不限购情况下，同时消费次数
		if pc.Count >= p.limit {
			return errorx.New("限购")
		}

		return p.flushCache(ctx, key, PurchaseLimitCache{
			Start: pc.Start,
			Count: pc.Count + 1,
		})
	}

	// 刷新
	return p.flushCache(ctx, key, PurchaseLimitCache{
		Start: now,
		Count: 1,
	})
}

func (p *PurchaseLimit) flushCache(ctx context.Context, key string, cache PurchaseLimitCache) error {

	if err := p.cli.Set(ctx, key, jsonx.MarshalToStringNoErr(cache)); err != nil {
		return err
	}

	if p.expire > 0 {
		if err := p.cli.Exipre(ctx, key, time.Duration(p.expire)*time.Second); err != nil {
			return err
		}
	}

	return nil
}

func (p *PurchaseLimit) judgeWindow(start time.Time, now time.Time) bool {
	switch p.typ {
	case PurchaseLimitTypeDay:
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", start.Format("2006-01-02 ")+p.at, time.Local)
		t2, _ := time.ParseInLocation("2006-01-02 15:04:05", start.AddDate(0, 0, 1).Format("2006-01-02 ")+p.at, time.Local)

		return now.After(start) && start.After(t1) && now.Before(t2)
	case PurchaseLimitTypeWeek:
		reduce := int(start.Weekday()) - 1
		if reduce == -1 {
			reduce = 6
		}
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", start.AddDate(0, 0, -reduce).Format("2006-01-02 ")+p.at, time.Local)
		t2, _ := time.ParseInLocation("2006-01-02 15:04:05", start.AddDate(0, 0, 7-reduce).Format("2006-01-02 ")+p.at, time.Local)

		return now.After(start) && start.After(t1) && now.Before(t2)
	case PurchaseLimitTypeDuration:

		return now.After(start) && now.Unix()-start.Unix() < p.durateion
	case PurchaseLimitTypeLasting:

		return true
	}

	return false
}

// GetCount 查询额度
func (p *PurchaseLimit) GetUseCount(ctx context.Context, actityId string) (uint64, error) {
	key := p.key(ctx, actityId)
	pc, err := p.getCache(ctx, key)
	if err != nil {
		return 0, err
	}

	if pc == nil {
		return 0, nil
	}

	if p.judgeWindow(pc.Start, time.Now()) {
		return pc.Count, nil
	}

	return 0, nil
}

// todo
func (p *PurchaseLimit) getCache(ctx context.Context, key string) (*PurchaseLimitCache, error) {
	str, err := p.cli.GetString(ctx, key)
	if err != nil {
		return nil, err
	}

	if str == "" {
		return nil, nil
	}

	var pc PurchaseLimitCache

	if err := jsonx.UnmarshalString(str, &pc); err != nil {
		return nil, err
	}

	return &pc, nil
}

func (p *PurchaseLimit) key(ctx context.Context, actityId string) string {
	return fmt.Sprintf("purchaselimitx:%v:%v:%v", actityId, p.typ, p.at)
}
