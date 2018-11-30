package controllers

import (
	"gin_mysql/models/mysql_handle"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func MyselfAPI1(c *gin.Context) {
	mysql_handle.Connect_mysql(c)

}
