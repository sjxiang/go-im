package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
}