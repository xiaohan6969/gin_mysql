package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//定义 Person 结构，字段须为首字母大写
func Conntct_mongo(c *gin.Context) {

	type Person struct {
		Name   string
		Place  string
		Gender string
		Age    int
		Phone  string
	}

	//可本地可远程，不指定协议时默认为http协议访问，此时需要设置 mongodb 的
	//nohttpinterface=false来打开httpinterface。
	//也可以指定mongodb协议，如 "mongodb://127.0.0.1:27017"
	var MOGODB_URI = "127.0.0.1:27017"
	//连接
	session, err := mgo.Dial(MOGODB_URI)
	//连接失败时终止
	if err != nil {
		panic(err)
	}
	//延迟关闭，释放资源
	defer session.Close()
	//设置模式
	session.SetMode(mgo.Monotonic, true)
	//选择数据库与集合
	a := session.DB("stu").C("haha")
	//插入文档
	err = a.Insert(
		&Person{Name: "陈耀灿", Place: "上海", Gender: "男", Age: 23, Phone: "130-7881-7881"},
		&Person{Name: "彭帆", Place: "成都", Gender: "男", Age: 24, Phone: "130-7881-5721"})
	//出错判断
	if err != nil {
		log.Fatal(err)
	}
	//查询文档
	result := Person{}
	//注意mongodb存储后的字段大小写问题
	err = a.Find(bson.M{"name": "陈耀灿"}).One(&result)
	//出错判断
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Phone: %d , mongo数据库添加成功", result.Phone)

}
