package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Test(c *gin.Context) {
	type StuRead struct {
		Qqq interface{} `json:"Qqq"`
		Www interface{} `json:"Www"`
		Aww interface{} `json:"qw"`
	}
	//获取前段法来的json
	data, _ := ioutil.ReadAll(c.Request.Body)
	str := []byte(data)
	//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象StuRead{}
	//2.可以直接stu:=new(StuRead),此时的stu自身就是指针
	stu := StuRead{}
	err := json.Unmarshal(str, &stu)
	//res, _ := json.Marshal(data)
	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, stu)
	fmt.Println(string(data))
	//c.BindJSON(string(res))
}
