package db

import (
	"fmt"
	"testing"
)

func init(){
	// 初始化连接
	err :=InitDB("root:123456@tcp(127.0.0.1:3306)/blog?parseTime=true")
	if err != nil {
		fmt.Println(err)
		panic("connect to mysql database failed")
	}
}

func TestGetArticleList(t *testing.T){
	list,err:=GetArticleList(0,2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list)
}

func TestGetArticleDetail(t *testing.T){
	detail,err:=GetArticleDetail(1)
	if err != nil {
		t.Fatal(err)

	}
	fmt.Printf("%#v\n",detail)
}
