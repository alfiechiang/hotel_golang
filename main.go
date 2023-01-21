package main

import (
	"hotel/conf"
	"hotel/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	r := server.NewRouter()
	r.Run(":6000")
}
