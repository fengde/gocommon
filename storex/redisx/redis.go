package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

// NewClient 创建集群对象，同时适用单实例
func NewClient(ctx context.Context, addr string, defaultDB int, password string) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           defaultDB,
		DialTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	c := Client{
		client: client,
	}

	if err := c.PingCheck(ctx); err != nil {
		return nil, err
	}

	return &c, nil
}

// Set 设置key, value
func (p *Client) Set(ctx context.Context, key string, value interface{}) error {
	status := p.client.Set(ctx, key, value, 0)
	return status.Err()
}

// Set 设置key, value 带超时
func (p *Client) SetWithExpire(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return p.client.Set(ctx, key, value, expiration).Err()
}

// GetString
func (p *Client) GetString(ctx context.Context, key string) (string, error) {
	return p.client.Get(ctx, key).Result()
}

// GetFloat64
func (p *Client) GetFloat64(ctx context.Context, key string) (float64, error) {
	return p.client.Get(ctx, key).Float64()
}

// GetInt64
func (p *Client) GetInt64(ctx context.Context, key string) (int64, error) {
	return p.client.Get(ctx, key).Int64()
}

// GetUint64
func (p *Client) GetUint64(ctx context.Context, key string) (uint64, error) {
	return p.client.Get(ctx, key).Uint64()
}

// HMSet
func (p *Client) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	return p.client.HMSet(ctx, key, fields).Err()
}

// HSet
func (p *Client) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return p.client.HSet(ctx, key, field, value).Err()
}

// HGetString
func (p *Client) HGetString(ctx context.Context, key string, field string) (string, error) {
	return p.client.HGet(ctx, key, field).Result()
}

// HGetFloat64
func (p *Client) HGetFloat64(ctx context.Context, key string, field string) (float64, error) {
	return p.client.HGet(ctx, key, field).Float64()
}

// HGetInt64
func (p *Client) HGetInt64(ctx context.Context, key string, field string) (int64, error) {
	return p.client.HGet(ctx, key, field).Int64()
}

// HGetUint64
func (p *Client) HGetUint64(ctx context.Context, key string, field string) (uint64, error) {
	return p.client.HGet(ctx, key, field).Uint64()
}

// HGetAll
func (p *Client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return p.client.HGetAll(ctx, key).Result()
}

// HDel 删除hash内field
func (p *Client) HDel(ctx context.Context, key string, fields ...string) error {
	return p.client.HDel(ctx, key, fields...).Err()
}

// HExist 判断hash内field是否存在
func (p *Client) HExist(ctx context.Context, key, field string) (bool, error) {
	cmd := p.client.HExists(ctx, key, field)
	return cmd.Val(), cmd.Err()
}

// Exipre 设置过期时间
func (p *Client) Exipre(ctx context.Context, key string, expiration time.Duration) error {
	return p.client.Expire(ctx, key, expiration).Err()
}

// Del 删除key
func (p *Client) Del(ctx context.Context, keys ...string) error {
	return p.client.Del(ctx, keys...).Err()
}

// Exist 检查key是否存在
func (p *Client) Exist(ctx context.Context, key string) (bool, error) {
	cmd := p.client.Exists(ctx, key)
	return cmd.Val() == 1, cmd.Err()
}

func (p *Client) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	return p.client.Do(ctx, args...)
}

// NewLocker 创建分布式锁
func (p *Client) NewLocker(sourceID string, autoUnlockSecond int64) *Locker {
	return NewLocker(p, sourceID, autoUnlockSecond)
}

// NewList 创建分布式队列
func (p *Client) NewList(listName string) *List {
	return NewList(p, listName)
}

// GetClient获取原生client
func (p *Client) GetClient() *redis.Client {
	return p.client
}

// PingCheck 检查连接是否可用
func (p *Client) PingCheck(ctx context.Context) error {
	_, err := p.client.Ping(ctx).Result()
	return err
}

// IsRedisNil 判断err是否为redis nil
func IsRedisNil(err error) bool {
	return err == redis.Nil
}
