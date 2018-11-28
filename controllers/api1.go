package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

/**
 * [测试连接数据库]
 * @Author   CaiYu
 * @DateTime 2017-01-11T14:40:34+0800
 * @param dbname        数据库名
 * @param dbtype        数据库类型
 * @param user          用户名
 * @param password      密码
 * @param port          端口
 * @param host          主机
*/
func MyselfAPI1(c *gin.Context)  {
	var(
		status int
		desc string
	)
	dbtype := c.Query("dbtype")
	dbname := c.Query("dbname")
	user   := c.Query("user")
	password := c.Query("password")
	host := c.Query("host")
	port := c.Query("port")

	constr := user+":"+password+"@tcp("+host+":"+port+")/"+dbname

	db, err := sql.Open(dbtype, constr)
	err = db.Ping()  //sql.Open无法断定数据库是否正常连接，所以调用db.Ping()进行判定
	if err != nil {
		status = 300
		desc = "数据库连接失败"
	}else{
		status = 200
		desc = "数据库连接成功"
	}
	c.JSON(200, gin.H{"status": status,"msg": desc})

}
