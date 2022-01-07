package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"job-portal/app/config"
	"os"
)

func LoadConfig(filenames ...string)  {
	err := godotenv.Load(filenames...)
	fmt.Println("ERROR : ", err)
	if err != nil {
		os.Setenv("MONGO_DBNAME",config.MONGO_DBNAME)
		os.Setenv("MONGO_PORT",config.MONGO_PORT)
		os.Setenv("MONGO_HOST",config.MONGO_HOST)
		os.Setenv("MONGO_USERNAME",config.MONGO_USERNAME)
		os.Setenv("MONGO_PASSWORD",config.MONGO_PASSWORD)
	}
}
