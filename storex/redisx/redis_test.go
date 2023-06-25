package redisx

import (
	"context"
	"testing"
	"time"
)

var client, _ = NewClient(context.Background(), "127.0.0.1:6379", 0, "")

func TestClient_Del(t *testing.T) {
	client.Set(context.Background(), "test", "hello world")
	value, err := client.GetString(context.Background(), "test:del")
	t.Log(value, err)
	client.Del(context.Background(), "test:del")
	value, err = client.GetString(context.Background(), "test:del")
	t.Log(value, err)
}

func TestClient_Exipre(t *testing.T) {
	client.SetWithExpire(context.Background(), "test:exipre", "hello world", 10*time.Second)
}

func TestClient_GetFloat64(t *testing.T) {
	client.Set(context.Background(), "test:getfloat64", 10.99)
	t.Log(client.GetFloat64(context.Background(), "test:getfloat64"))
}

func TestClient_GetInt64(t *testing.T) {
	client.Set(context.Background(), "test:getint64", 8)
	t.Log(client.GetInt64(context.Background(), "test:getint64"))
}

func TestClient_GetString(t *testing.T) {
	client.Set(context.Background(), "test:getstring", "hello world")
	t.Log(client.GetString(context.Background(), "test:getstring"))
}

func TestClient_GetUInt64(t *testing.T) {
	client.Set(context.Background(), "test:getuint64", 8999999999999)
	t.Log(client.GetUint64(context.Background(), "test:getuint64"))
}

func TestClient_HGetAll(t *testing.T) {
	client.HSet(context.Background(), "test:hgetall", "float64", 1.102)
	client.HSet(context.Background(), "test:hgetall", "int64", -100)
	client.HSet(context.Background(), "test:hgetall", "string", "hello world")
	client.HSet(context.Background(), "test:hgetall", "uint64", 100)
	t.Log(client.HGetAll(context.Background(), "test:hgetall"))

}

func TestClient_HGetFloat64(t *testing.T) {
	t.Log(client.HGetFloat64(context.Background(), "test:hgetall", "float64"))
}

func TestClient_HGetInt64(t *testing.T) {
	t.Log(client.HGetInt64(context.Background(), "test:hgetall", "int64"))
}

func TestClient_HGetString(t *testing.T) {
	t.Log(client.HGetString(context.Background(), "test:hgetall", "string"))
}

func TestClient_HGetUint64(t *testing.T) {
	t.Log(client.HGetUint64(context.Background(), "test:hgetall", "uint64"))
}

func TestClient_HMSet(t *testing.T) {
	client.HMSet(context.Background(), "test:hmset", map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 11.1,
		"d": "d",
	})
	t.Log(client.HGetAll(context.Background(), "test:hmset"))
}

func TestClient_HSet(t *testing.T) {
	client.HSet(context.Background(), "test:hgetall", "float64", 1.102)
	client.HSet(context.Background(), "test:hgetall", "int64", -100)
	client.HSet(context.Background(), "test:hgetall", "string", "hello world")
	client.HSet(context.Background(), "test:hgetall", "uint64", 100)
	t.Log(client.HGetAll(context.Background(), "test:hgetall"))
}

func TestClient_Set(t *testing.T) {
	client.Set(context.Background(), "test", "123")
	t.Log(client.GetString(context.Background(), "test"))
}

func TestClient_SetWithExpire(t *testing.T) {
	t.Log(client.SetWithExpire(context.Background(), "test", "123456", 100*time.Second))
}

func TestClient_HDel(t *testing.T) {
	client.HMSet(context.Background(), "test:hdel", map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	t.Log(client.HDel(context.Background(), "test:hdel", "a", "b", "c"))
}

func TestClient_Exist(t *testing.T) {
	t.Log(client.Exist(context.Background(), "test:hgetall"))
	t.Log(client.Exist(context.Background(), "test:hgetall22222"))
}

func TestClient_HExist(t *testing.T) {
	t.Log(client.HExist(context.Background(), "tst:hgetall", "abc"))
	t.Log(client.HExist(context.Background(), "tst:hgetall222", "abc"))
}
