package lib

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	config := LoadServerConfig()

	RedisPool = &redis.Pool{
		MaxIdle:     5,                 // idle的列表长度, 空闲的线程数
		MaxActive:   0,                 // 线程池的最大连接数， 0表示没有限制
		Wait:        true,              // 当连接数已满，是否要阻塞等待获取连接。false表示不等待，直接返回错误。
		IdleTimeout: 200 * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) { // 创建链接
			c, err := redis.Dial("tcp", config.RedisHost)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", config.RedisIndex); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //一个测试链接可用性
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	fmt.Println("Redis init on port ", config.RedisHost)
}

// get
func GetKey(key string) (string, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.String(rds.Do("GET", key))
}

// set expires为0时，表示永久性存储
func SetKey(key, value interface{}, expires int) error {
	rds := RedisPool.Get()
	defer rds.Close()
	if expires == 0 {
		_, err := rds.Do("SET", key, value)
		return err
	} else {
		_, err := rds.Do("SETEX", key, expires, value)
		return err
	}
}

// del
func DelKey(key string) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_, err := rds.Do("DEL", key)
	return err
}

// lrange
func LRange(key string, start, stop int64) ([]string, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Strings(rds.Do("LRANGE", key, start, stop))
}

// lpop
func LPop(key string) (string, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.String(rds.Do("LPOP", key))
}

// LPushAndTrimKey
func LPushAndTrimKey(key, value interface{}, size int64) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_ = rds.Send("MULTI")
	_ = rds.Send("LPUSH", key, value)
	_ = rds.Send("LTRIM", key, size-2*size, -1)
	_, err := rds.Do("EXEC")
	return err
}

// RPushAndTrimKey
func RPushAndTrimKey(key, value interface{}, size int64) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_ = rds.Send("MULTI")
	_ = rds.Send("RPUSH", key, value)
	_ = rds.Send("LTRIM", key, size-2*size, -1)
	_, err := rds.Do("EXEC")
	return err

}

// ExistsKey
func ExistsKey(key string) (bool, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Bool(rds.Do("EXISTS", key))
}

// ttl 返回剩余时间
func TTLKey(key string) (int64, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("TTL", key))
}

// incr 自增
func Incr(key string) (int64, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("INCR", key))
}

// Decr 自减
func Decr(key string) (int64, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("DECR", key))
}

// mset 批量写入 rds.Do("MSET", "ket1", "value1", "key2","value2")
func MsetKey(key_value ...interface{}) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_, err := rds.Do("MSET", key_value...)
	return err
}

// mget  批量读取 mget key1, key2, 返回map结构
func MgetKey(keys ...interface{}) map[interface{}]string {
	rds := RedisPool.Get()
	defer rds.Close()
	values, _ := redis.Strings(rds.Do("MGET", keys...))
	resultMap := map[interface{}]string{}
	keyLen := len(keys)
	for i := 0; i < keyLen; i++ {
		resultMap[keys[i]] = values[i]
	}
	return resultMap
}