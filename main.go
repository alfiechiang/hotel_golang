package main

import (
	"hotel/conf"
	"hotel/server"
	"net/http"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	r := server.NewRouter()
	r.StaticFS("/img", http.Dir("./img"))
	r.Run(":3000")

}
