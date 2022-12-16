package main

import (
	"github.com/sjxiang/go-im/conf"
	"github.com/sjxiang/go-im/router"
)


func main() {

	// 加载配置
	conf.Init()

	// 加载路由
	r := router.Setup()
	r.Run(":8080")
}