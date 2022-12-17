package helper

import (
	"time"
	"math/rand"
)


// RandStringRunes 生成长度为 length 随机数字字符串
func RandStringRunes(length int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")  

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	
	return string(b)
}


func GenerateRandomNum() string {
	length := 6
	return RandStringRunes(length)
}