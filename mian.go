package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Content，封装了request和response
	r.GET("/login/:name", func(c *gin.Context) {
		//获取api参数
		name := c.Param("name")
		c.String(http.StatusOK, name)

	})
	r.POST("/getParam", func(c *gin.Context) {
		//获取url参数,?后面的参数name、id
		name := c.DefaultQuery("name", "xiaoli")
		id := c.Query("id")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name) + id)
	})
	//获取表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.JSON(http.StatusOK, fmt.Sprintf("username: %s, password: %s, types : %s", username, password, types))
	})
	//上传单个文件
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/onloadfile", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500,"上传出错")
		}
		dst := `E:\work\temp\`+file.Filename
		//需要使用绝对路径
		err1 := c.SaveUploadedFile(file, dst)
		if err1 != nil {
			log.Fatal(err1)
		}
		c.String(http.StatusOK, file.Filename)
	})
	//3.监听端口，默认8080
	//Run("")不指定端口号，默认8080
	r.Run(":9090")
}
