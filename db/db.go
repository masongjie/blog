package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis"


)

// 初始化一个mysql的连接
var DB *sqlx.DB
var RedisCli *redis.Client



func InitDB(dsn string)(err error){
	// 初始化db
	// dsn="username@password@tcp(127.0.0.1:3306)/test?parseTime=true"
	DB,err = sqlx.Connect("mysql",dsn)
	//DB.SetMaxOpenConns(30) // 最大连接数
	//DB.SetMaxIdleConns(3)  // 最大闲置连接数
	return
}

func InitRedisClient(addr string,db int)(err error){
	RedisCli = redis.NewClient(&redis.Options{
		Addr:addr,
		DB:db,
		Password:"",
	})
	_,err =RedisCli.Ping().Result()
	return
}




