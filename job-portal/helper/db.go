package helper

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func init() {
	LoadConfig(".env")
}

func Connection() *mongo.Database {
	dbname := os.Getenv("MONGO_DBNAME")
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	port := os.Getenv("MONGO_PORT")
	host := os.Getenv("MONGO_HOST")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",username,password,host,port)
	opts := options.Client().ApplyURI(uri).SetMaxPoolSize(50).SetMinPoolSize(10).SetMaxConnIdleTime(1 * time.Minute)
	client,err := mongo.Connect(context.TODO(),opts)
	Panic(err)
	db := client.Database(dbname)
	return db
}

func Redis() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Username:           "root",
		Password:           "root",
		DB:                 0,
		MinIdleConns:       10,
	})
	return client
}