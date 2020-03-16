# blog
### 介绍
这是简单的blog，是基于gin框架开发的blog
<br>
数据使用：mysql和redis。

### 配置文件
debug 模式：conf/config.ini

release 模式：conf/config_release.ini

## 启动
### debug模式
    将对应的mysql，redis的信息填入config.ini中
mysql 生成表文件（sql文件）：./blog/sql/blog.sql
```shell script
mysql -uroot -p****** <./blog/sql/blog.sql
```
编译成对应的系统的执行文件(go verison 1.11以上需要支持go mod模式)
```shell script
go build
```
### 运行
```shell script
./blogger
```

## release 模式
### 启动（使用docker-compose 启动）
需要在对应的系统新建/data/blog_data/和 /data/blog/log/文件夹
，为持久化数据。
mysql 容器第一次启动后，生成数据表
```shell script
docker-compose up
```
```shell script
# blog目录下
docker cp ./sql/blog.sql containerid:/root
docker exec -it containerid /bin/sh
containerid $:mysql -uroot -p123456</root/blog.sql
containerid $:exit
```

## 访问
```shell script
curl 127.0.0.1:8080
```

