package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

type Temperature struct {
	TimeStamp int64 `json:"timestamp"`
	Degree    int   `json:"degree"`
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Go Redis Hello World!")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	Panic(err)

	temperatureInBytes, err := json.Marshal(Temperature{TimeStamp: 100, Degree: 25})
	Panic(err)

	err = client.Set(ctx, "temperature", temperatureInBytes, 0).Err()
	Panic(err)

	val, err := client.Get(ctx, "temperature").Result()
	Panic(err)

	fmt.Println("got", val)
}
