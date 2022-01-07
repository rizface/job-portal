package myredis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"job-portal/helper"
	"os"
	"time"
)

var Redis *redis.Client = helper.Redis()

func RedisPublish(channel string, msg string,insertId interface{}) {
	ctx,cancel := context.WithTimeout(context.Background(),10* time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%s/%s",os.Getenv("APP_HOST"),os.Getenv("APP_PORT"),insertId)
	payload := map[string]string {
		"email": msg,
		"url": url,
	}
	payloadJson,_ := json.Marshal(payload)
	Redis.Publish(ctx,channel,string(payloadJson))
}
