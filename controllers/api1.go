package controllers

import (
	"gin_mysql/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func MyselfAPI1(c *gin.Context) {
	models.Connect_mysql(c)

}
