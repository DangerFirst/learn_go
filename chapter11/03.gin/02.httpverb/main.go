package main

import (
	"encoding/base64"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"learn.go/pkg/apis"
)

func main() {
	r := gin.Default()
	pprof.Register(r)
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte(`你好，gin！`))
	})
	r.GET("/list", func(c *gin.Context) {
		c.JSON(200, []apis.PersonalInformation{
			{
				Name:   "小黑",
				Sex:    "男",
				Tall:   1.8,
				Weight: 77,
				Age:    29,
			},
			{
				Name:   "小和",
				Sex:    "男",
				Tall:   1.9,
				Weight: 72,
				Age:    22,
			},
		})
	})
	r.POST("/register", func(c *gin.Context) {
		pi := &apis.PersonalInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(400, gin.H{
				"message": "无法读取个人注册信息",
			})
			return
		}
		//todo 注册到排行榜
		c.JSON(200, nil)
	})
	r.GET("/getwithquery", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	r.GET("/getwithparam/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	r.GET("/getwithparam/xiaohei", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"special": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	r.Run(":8081")
}
