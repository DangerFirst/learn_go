package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"moments/frinterface"
	"moments/pkg/apis"
	"net/http"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:123qwe@tcp(localhost)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func main() {
	conn := connectDb()
	var rankServer frinterface.ServeInterface = NewDbRank(conn, NewFatRateRank())
	var momentsServer frinterface.MomentsServerInterface = NewDbMoments(conn, NewMoments())
	if initRank, ok := rankServer.(frinterface.RankInitInterface); ok {
		if err := initRank.Init(); err != nil {
			log.Fatal("初始化失败", err)
		}
	}

	r := gin.Default()
	pprof.Register(r)
	//发朋友圈
	r.POST("/publish", func(c *gin.Context) {
		var ps *apis.PersonalShow
		if err := c.BindJSON(&ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析发布信息" + err.Error(),
			})
			return
		}
		if err := momentsServer.Publish(ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "发布失败" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	//删朋友圈
	r.PUT("/delete", func(c *gin.Context) {
		var ps *apis.PersonalShow
		if err := c.BindJSON(&ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法更新信息" + err.Error(),
			})
			return
		}
		if err := momentsServer.Delete(ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新失败" + err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}
	})
	//刷朋友圈
	r.GET("/moments", func(c *gin.Context) {
		if mo, err := momentsServer.Moments(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取失败" + err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, mo)
		}
	})

	r.POST("/register", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息" + err.Error(),
			})
			return
		}
		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
	})

	r.PUT("/personalinfo", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法更新信息",
			})
			return
		}
		if resp, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, resp)
		}
	})
	r.GET("/rank/:name", func(c *gin.Context) {
		name := c.Param("name")
		if fr, err := rankServer.GetFatRate(name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取个人排行失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, fr)
		}
	})
	r.GET("/ranktop", func(c *gin.Context) {
		if frTop, err := rankServer.GetTop(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取排行失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, frTop)
		}
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
