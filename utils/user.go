package utils

import (
	"blogger/db"
	"blogger/log"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"github.com/satori/go.uuid"
)

// 生成用户登录后的uuid

func SetCookieUUid(ttl int) (uuidStr string, err error) {
	u1 := uuid.Must(uuid.NewV4(), nil)
	uuidStr = u1.String()
	//err = db.RedisCli.Set(uuidStr, 1, time.Duration(ttl)*time.Second).Err()
	data := map[string]interface{}{
		"isLogin": true,
	}
	p := db.RedisCli.Pipeline()
	p.HMSet(uuidStr, data)
	p.Expire(uuidStr, time.Duration(ttl)*time.Second)
	p.Exec()
	err = p.Close()
	if err != nil {
		log.SugarLog.Infow("redis store uuid failed")
		return
	}
	return

}

// 生成一个cookie
func MakeCookie(c *gin.Context,ttl int)(token string,err error){
	u1 := uuid.Must(uuid.NewV4(),nil)
	uuidStr:=u1.String()
	data := map[string]interface{}{
		"isLogin":0,
		"view_times":0,
	}

	p:=db.RedisCli.Pipeline()
	p.HMSet(uuidStr,data)
	p.Expire(uuidStr,time.Duration(ttl)*time.Second)
	p.Exec()
	err =p.Close()
	if err != nil {
		return
	}

	c.SetCookie("token",uuidStr,ttl,"/","/",false,false)
	return
}

// 判断token是否存在
func IsExistToken(token string)(isExist bool){
	num,err:=db.RedisCli.Exists(token).Result()
	if err != nil {
		return false
	}
	if num == 0{
		return false
	}
	typeStr,_:=db.RedisCli.Type(token).Result()
	if typeStr == "hash"{
		return true
	}
	return
}

// view_times +1
func AddViewTimes(token string){
	db.RedisCli.HIncrBy(token,"view_times",1)
}

// 判断 view_time 和 isLogin // true 表示不限制
func IsViewLimit(token string)(ok bool){
	data:=db.RedisCli.HGetAll(token).Val()
	if data["isLogin"] == "1"{
		// 已登录 不限制
		return true
	}
	// view_times
	viewTimesStr:=data["view_times"]
	viewTimes,err:=strconv.ParseInt(viewTimesStr,10,64)
	if err != nil {
		return false
	}
	if viewTimes <=100{
		return true
	}
	return 

}

func IsLogin(token string)bool{
	isLogin := db.RedisCli.HGet(token,"isLogin").Val()
	if isLogin == "1"{
		return true
	}
	return false
}

func SetIsLogin(token string)bool{
	ok:=db.RedisCli.HSet(token,"isLogin",1).Val()
	return !ok
}


