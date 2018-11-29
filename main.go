package main

import (
	"fmt"
	"gin_mysql/routers"
)

func main() {
	//路由入口
	fmt.Println("服务成功开启：～～～～～")
	routers.RegisterAPIRouter()
}
