package main

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"blogger/controller"
	"blogger/db"
	"blogger/log"
	"blogger/conf"

)

func main() {
	r:=gin.Default()
	// 加载配置
	err:= conf.InitConfig()
	if err != nil {
		return
	}
	dsn:=conf.GetMysqlDSN()
	redisHost,redisDb := conf.GetRedisConf()

	fmt.Println(dsn,redisHost,redisDb)
	err = db.InitDB(dsn)
	if err != nil {
		fmt.Println("mysql connect failed",err)
		return
	}
	err = db.InitRedisClient(redisHost,redisDb)
	if err != nil {
		fmt.Println("redis connect failed",err)
		return
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200,gin.H{"msg":"hello"})
	})

	// 设置日志输出的
	err =log.Initlog()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 初始化zap 日志
	err = log.InitSagurLog()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 加载全局中间件
	r.Use(controller.VisitTimesLimit())
	// url 缓存中间件
	//r.Use(controller.CacheUrl(60))


	r.LoadHTMLGlob("./views/*") // 加载模板目录
	//加载静态文件系统
	r.Static("/static","./static")
	// 404
	r.NoRoute(func(c *gin.Context) {
		log.SugarLog.Infow("failed to fetch URL",
		)
		c.HTML(404,"views/404.html",nil)
	})

	r.GET("/",controller.Index) // url + handler
	r.GET("/category",controller.CategoryHandler)

	r.GET("/article/detail",controller.DetailHandler)

	r.GET("/article/new/",controller.Authentication(),controller.NewArticle)

	r.POST("/article/submit/",controller.AddArticle)

	// 登录
	r.Any("/login",controller.LoginHandler)

	// 评论
	r.POST("/comment/submit/",controller.SubmitComment)

	// 新增分类
	r.POST("/category/add",controller.AddCategory)

	r.Run("0.0.0.0:8080")
}
