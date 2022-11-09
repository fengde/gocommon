package redisx

import (
	"testing"
	"time"
)

var client, _ = NewClient("127.0.0.1:6379", 0, "")

func TestClient_Del(t *testing.T) {
	client.Set("test", "hello world")
	value, err := client.GetString("test:del")
	t.Log(value, err)
	client.Del("test:del")
	value, err = client.GetString("test:del")
	t.Log(value, err)
}

func TestClient_Exipre(t *testing.T) {
	client.SetWithExpire("test:exipre", "hello world", 10*time.Second)
}

func TestClient_GetFloat64(t *testing.T) {
	client.Set("test:getfloat64", 10.99)
	t.Log(client.GetFloat64("test:getfloat64"))
}

func TestClient_GetInt64(t *testing.T) {
	client.Set("test:getint64", 8)
	t.Log(client.GetInt64("test:getint64"))
}

func TestClient_GetString(t *testing.T) {
	client.Set("test:getstring", "hello world")
	t.Log(client.GetString("test:getstring"))
}

func TestClient_GetUInt64(t *testing.T) {
	client.Set("test:getuint64", 8999999999999)
	t.Log(client.GetUint64("test:getuint64"))
}

func TestClient_HGetAll(t *testing.T) {
	client.HSet("test:hgetall", "float64", 1.102)
	client.HSet("test:hgetall", "int64", -100)
	client.HSet("test:hgetall", "string", "hello world")
	client.HSet("test:hgetall", "uint64", 100)
	t.Log(client.HGetAll("test:hgetall"))

}

func TestClient_HGetFloat64(t *testing.T) {
	t.Log(client.HGetFloat64("test:hgetall", "float64"))
}

func TestClient_HGetInt64(t *testing.T) {
	t.Log(client.HGetInt64("test:hgetall", "int64"))
}

func TestClient_HGetString(t *testing.T) {
	t.Log(client.HGetString("test:hgetall", "string"))
}

func TestClient_HGetUint64(t *testing.T) {
	t.Log(client.HGetUint64("test:hgetall", "uint64"))
}

func TestClient_HMSet(t *testing.T) {
	client.HMSet("test:hmset", map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 11.1,
		"d": "d",
	})
	t.Log(client.HGetAll("test:hmset"))
}

func TestClient_HSet(t *testing.T) {
	client.HSet("test:hgetall", "float64", 1.102)
	client.HSet("test:hgetall", "int64", -100)
	client.HSet("test:hgetall", "string", "hello world")
	client.HSet("test:hgetall", "uint64", 100)
	t.Log(client.HGetAll("test:hgetall"))
}

func TestClient_Set(t *testing.T) {
	client.Set("test", "123")
	t.Log(client.GetString("test"))
}

func TestClient_SetWithExpire(t *testing.T) {
	t.Log(client.SetWithExpire("test", "123456", 100*time.Second))
}

func TestClient_HDel(t *testing.T) {
	client.HMSet("test:hdel", map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	t.Log(client.HDel("test:hdel", "a", "b", "c"))
}

func TestClient_Exist(t *testing.T) {
	t.Log(client.Exist("test:hgetall"))
	t.Log(client.Exist("test:hgetall22222"))
}

func TestClient_HExist(t *testing.T) {
	t.Log(client.HExist("tst:hgetall", "abc"))
	t.Log(client.HExist("tst:hgetall222", "abc"))
}
