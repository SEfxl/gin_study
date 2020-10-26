package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//engine := gin.Default()
	//
	//engine.GET("/hello", func(context *gin.Context) {
	//	fmt.Println("请求路径",context.FullPath())
	//	context.Writer.Write([]byte("hello gin\n"))
	//})
	//
	//if err := engine.Run(":8090"); err != nil {
	//	log.Fatal(err.Error())
	//}

	//GetAndPostInterface()

	RouterMain()
}

func GetAndPostInterface() {

	engine := gin.Default()

	//get
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "default name")
		fmt.Println(name)

		//输出
		context.Writer.Write([]byte("Hello " + name + "\n"))
	})

	//post
	engine.Handle("POST", "/login", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		username := context.PostForm("username")
		password := context.PostForm("password")

		fmt.Println("username:", username, "password:", password)

		context.Writer.Write([]byte(username + "登录" + "\n"))
	})

	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}

}

func RouterMain() {
	engine := gin.Default()

	//get
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		name := context.Query("name")
		fmt.Println(name)
		context.Writer.Write([]byte("hello " + name + "\n"))
	})

	//post
	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println("username:", username)
		}

		password, exist := context.GetPostForm("password")
		if exist {
			fmt.Println("username:", password)
		}

		context.Writer.Write([]byte("hello, " + username))
	})

	//delete
	engine.DELETE("/user/:id", func(context *gin.Context) {
		uId := context.Param("id")
		fmt.Println(uId)
		context.Writer.Write([]byte("delete 用户Id:" + uId + "\n"))
	})

	engine.Run()
}
