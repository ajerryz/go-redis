package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pong := rdb.Ping(context.Background()).Val()
	fmt.Printf("pong: %s\n", pong)
}
