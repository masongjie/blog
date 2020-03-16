package controller

import (
	"blogger/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cookie 中间件

// 限制访问次数 一次过期时间内限制50次,登录不限次
func VisitTimesLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 是否首次进入 或者 没有token
		uuidStr, err := c.Cookie("token")
		if err != nil {
			// 没有cookie ，生成uuid 保存cookie
			uuidStr, _ = utils.MakeCookie(c, 60*60)
		}
		// uuidStr 判断是否存在redis中，存在 view_times+1,
		// 不存在 生成一个新的token 在
		
		ok:= utils.IsExistToken(uuidStr)
		if ok{
			// view_time +1
			utils.AddViewTimes(uuidStr)
		}else{
			//生成新token
			uuidStr,_ = utils.MakeCookie(c,60*60)
			utils.AddViewTimes(uuidStr)
		}
		// 如果没有登录并且views_times 超过20次的话
		isLimit := utils.IsViewLimit(uuidStr)
		if isLimit{
			c.Next()
		}else{
			c.AbortWithStatusJSON(200,gin.H{
				"err":"view to many",
			})
			
			return
		}

		
		
	}
}

// 根据路由缓存数据
func CacheUrl(ttl int)gin.HandlerFunc{
	return func(c *gin.Context) {
		if c.Request.Method == "GET"{
			// 获取url
			url := c.Request.RequestURI
			fmt.Println("request url:",url)
			// 读取redis中缓存的内容，在保存在url中发送给客户端
			
		}
		// 不是get请求不读缓存
		c.Next()
		// 读取
		
		
	}
}


//认证装饰器
func Authentication() gin.HandlerFunc{
	return func(c *gin.Context) {
		targetUrl := c.Request.RequestURI
		
		url:= fmt.Sprintf("/login?target=%s",targetUrl)
		fmt.Println("target url middle ware",url)
		uuidStr,err:=c.Cookie("token")
		if err != nil {
			// 未登录 跳转到登录
			c.Redirect(http.StatusFound,url)
			return
		}
		// 查看当前用户是否登录
		isLogin:=utils.IsLogin(uuidStr)
		if isLogin{
			c.Next()
		}else{
			c.Redirect(http.StatusFound,url)
			return
		}
	}
}
