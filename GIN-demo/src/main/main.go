package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// RESTful 路由
	router.GET("/hello", helloWorldGet)
	router.POST("/hello", helloWorldPost)

	//不支持正则路由
	//提取path中的参数
	router.GET("param/:id", fetchId)

	// 组路由
	group1 := router.Group("/g1")
	{
		group1.GET("/action1", action1)
		group1.GET("/action2", action2)
		group1.GET("/action3", action3)
	}

	router.Run("127.0.0.1:8082")
}

func action1(context *gin.Context) {
	context.String(http.StatusOK, "action 1")
}
func action2(context *gin.Context) {
	context.String(http.StatusOK, "action 2")
}
func action3(context *gin.Context) {
	context.String(http.StatusOK, "action 3")
}

func fetchId(context *gin.Context) {
	id := context.Param("id")
	context.String(http.StatusOK, fmt.Sprintf("id is : %s ", id))
}

func helloWorldGet(context *gin.Context) {
	context.String(http.StatusOK, "[gin]Hello , World in GET!")
}

//POST函数
func helloWorldPost(context *gin.Context) {
	context.String(http.StatusOK, "[gin]Hello , World in POST!")
}
