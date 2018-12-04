package routers

import (
	. "gin_mysql/controllers"
	. "gin_mysql/models/mongo_handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRouter() {
	//路由调用函数
	gin.SetMode(gin.DebugMode) //调试模式
	router := gin.Default()    // 获得路由实例

	// 监听data  source  runserver  的get和post请求，对应方法：runserver
	routerDatasource := router.Group("/v1")
	{
		routerDatasource.GET("/mongo", Conntct_mongo)
	}
	//分组示例
	myselfpractice := router.Group("/v2")
	{
		myselfpractice.GET("/mysql", MyselfAPI1)
	}

	router.POST("/test",Test)
	// 监听端口
	http.ListenAndServe(":9769", router)
}
