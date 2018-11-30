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
	}

	data, _ := ioutil.ReadAll(c.Request.Body)
	str := []byte(data)
	stu := StuRead{}
	err := json.Unmarshal(str, &stu)
	//res, _ := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, stu)
	fmt.Println(string(data))
	//c.BindJSON(string(res))
}
