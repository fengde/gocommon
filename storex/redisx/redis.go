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
func NewClient(addr string, defaultDB int, password string) (*Client, error) {
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

	if err := c.PingCheck(); err != nil {
		return nil, err
	}

	return &c, nil
}

func (p *Client) getCtx() context.Context {
	return context.Background()
}

// Set 设置key, value
func (p *Client) Set(key string, value interface{}) error {
	status := p.client.Set(p.getCtx(), key, value, 0)
	return status.Err()
}

// Set 设置key, value 带超时
func (p *Client) SetWithExpire(key string, value interface{}, expiration time.Duration) error {
	return p.client.Set(p.getCtx(), key, value, expiration).Err()
}

// GetString
func (p *Client) GetString(key string) (string, error) {
	return p.client.Get(p.getCtx(), key).Result()
}

// GetFloat64
func (p *Client) GetFloat64(key string) (float64, error) {
	return p.client.Get(p.getCtx(), key).Float64()
}

// GetInt64
func (p *Client) GetInt64(key string) (int64, error) {
	return p.client.Get(p.getCtx(), key).Int64()
}

// GetUint64
func (p *Client) GetUint64(key string) (uint64, error) {
	return p.client.Get(p.getCtx(), key).Uint64()
}

// HMSet
func (p *Client) HMSet(key string, fields map[string]interface{}) error {
	return p.client.HMSet(p.getCtx(), key, fields).Err()
}

// HSet
func (p *Client) HSet(key string, field string, value interface{}) error {
	return p.client.HSet(p.getCtx(), key, field, value).Err()
}

// HGetString
func (p *Client) HGetString(key string, field string) (string, error) {
	return p.client.HGet(p.getCtx(), key, field).Result()
}

// HGetFloat64
func (p *Client) HGetFloat64(key string, field string) (float64, error) {
	return p.client.HGet(p.getCtx(), key, field).Float64()
}

// HGetInt64
func (p *Client) HGetInt64(key string, field string) (int64, error) {
	return p.client.HGet(p.getCtx(), key, field).Int64()
}

// HGetUint64
func (p *Client) HGetUint64(key string, field string) (uint64, error) {
	return p.client.HGet(p.getCtx(), key, field).Uint64()
}

// HGetAll
func (p *Client) HGetAll(key string) (map[string]string, error) {
	return p.client.HGetAll(p.getCtx(), key).Result()
}

// HDel 删除hash内field
func (p *Client) HDel(key string, fields ...string) error {
	return p.client.HDel(p.getCtx(), key, fields...).Err()
}

// HExist 判断hash内field是否存在
func (p *Client) HExist(key, field string) (bool, error) {
	cmd := p.client.HExists(p.getCtx(), key, field)
	return cmd.Val(), cmd.Err()
}

// Exipre 设置过期时间
func (p *Client) Exipre(key string, expiration time.Duration) error {
	return p.client.Expire(p.getCtx(), key, expiration).Err()
}

// Del 删除key
func (p *Client) Del(keys ...string) error {
	return p.client.Del(p.getCtx(), keys...).Err()
}

// Exist 检查key是否存在
func (p *Client) Exist(key string) (bool, error) {
	cmd := p.client.Exists(p.getCtx(), key)
	return cmd.Val() == 1, cmd.Err()
}

func (p *Client) Do(args ...interface{}) *redis.Cmd {
	return p.client.Do(p.getCtx(), args...)
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
func (p *Client) PingCheck() error {
	_, err := p.client.Ping(p.getCtx()).Result()
	return err
}

// IsRedisNil 判断err是否为redis nil
func IsRedisNil(err error) bool {
	return err == redis.Nil
}
