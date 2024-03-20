package main

import (
	"time"

	_ "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}

}
