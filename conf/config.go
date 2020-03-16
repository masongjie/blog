package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Conf struct {
	MysqlConf `ini:"mysql"`
	RedisConf `ini:"redis"`
}

type MysqlConf struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

type RedisConf struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
	Db   int    `ini:"db"`
}

var conf Conf



func InitConfig() (err error) {
	if os.Getenv("GIN_MODE")=="release"{
		err = ini.MapTo(&conf, "./conf/config_release.ini")
		return
	}
	err = ini.MapTo(&conf, "./conf/config.ini")
	return

}

func GetMysqlDSN() (dsn string) {
	dsnFormat := `%s:%s@tcp(%s:%d)/blog?parseTime=True`
	dsn = fmt.Sprintf(dsnFormat, conf.User, conf.Password, conf.MysqlConf.Host, conf.MysqlConf.Port)
	return
}

func GetRedisConf() (string, int) {
	url := fmt.Sprintf("%s:%d", conf.RedisConf.Host, conf.RedisConf.Port)
	return url, conf.RedisConf.Db
}
