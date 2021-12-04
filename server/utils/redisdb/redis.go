package redisdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/wenknow/gin-vue-blog/server/global"
	"time"
)

var ctx = context.Background()

func Set(key string, value interface{}, expire int64) {
	err := global.REDIS.Set(ctx, key, value, time.Duration(expire)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) (string, bool) {
	val, err := global.REDIS.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		panic(err)
	} else {
		return val, true
	}
}
