package routers

import (
	"github.com/gin-gonic/gin"
	. "gin_mod/controllers"
	"net/http"
)


func RegisterAPIRouter() {
	//路由调用函数
	gin.SetMode(gin.DebugMode) //调试模式
	router := gin.Default()   // 获得路由实例

	// 监听data  source  runserver  的get和post请求，对应方法：runserver
	routerDatasource := router.Group("/data/source")
	{
		routerDatasource.GET("/practice", Runserver)
	}

	myselfpractice := router.Group("v2")
	{
		myselfpractice.GET("/slice", MyselfAPI1)
	}

	// 监听端口
	http.ListenAndServe(":9876",router)
}