package lock

import (
	"context"
	"shmiloveu.fun/startingGo/demo/redis/db"
	"time"
)

// todo可重入锁、未考虑集群安全性
func Lock(key, val string, exp time.Duration) bool {
	result, err := db.RedisCli.SetNX(context.TODO(), key, val, exp).Result()
	if err != nil {
		panic(err)
	}
	return result
}

//需要判断加锁人是不是本机 使用lua脚本实现原子性
func Unlock(key, val string) {
	luaScript := "if redis.call('get',KEYS[1]) == ARGV[1] then " +
		"return redis.call('del',KEYS[1]) else return 0 end"
	_, err := db.RedisCli.Eval(context.TODO(), luaScript, []string{key}, val).Result()
	if err != nil {
		panic(err)
	}
}
