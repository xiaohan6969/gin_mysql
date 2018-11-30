package mongo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Create_mongo(c *gin.Context) {
	err := a.Insert(
		&Person{Name: "陈耀灿", Place: "上海", Gender: "男", Age: 23, Phone: "130-7881-7881"},
		&Person{Name: "彭帆", Place: "成都", Gender: "男", Age: 24, Phone: "130-7881-5721"})
	//出错判断
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Phone: %d , mongo数据库添加成功", result.Phone)
}
