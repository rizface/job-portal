package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"os"
	"strconv"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolSize: 100,
		Username: "root",
		Password: "root",
	})
	return client
}

func sendMail(msg string) {

	var data = make(map[string]string)
	json.Unmarshal([]byte(msg),&data)
	port,_ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	server := mail.NewSMTPClient()
	server.Host = os.Getenv("SMTP_HOST")
	server.Port = port
	server.Username = os.Getenv("EMAIL")
	server.Password = os.Getenv("PASSWORD")
	server.Encryption = mail.EncryptionTLS

	smtpClient,err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.SetFrom(server.Username)
	email.AddTo(data["email"])
	email.SetBody(mail.TextHTML,"Copy & Paste Link Ini Untuk Aktivasi Akun : " + data["url"])
	email.SetSubject("Aktivasi Akun")
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("email berhasil dikirim")
	}
}

func loadConfig(fileNames ...string) {
	err := godotenv.Load(fileNames...)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadConfig(".env")
	ctx := context.Background()
	redisClient := RedisClient()
	for {
		func() {
			pubsub := redisClient.Subscribe(ctx,"job-portal-email")
			msg,err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Fatal(err)
			}
			go sendMail(msg.Payload)
		}()
	}

}
