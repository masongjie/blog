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

func TestGetCategory(t *testing.T) {
	category,err:=GetCategory(int64(1))
	if err != nil {
		fmt.Println(1111,err)
	}
	fmt.Println(category)
}

func TestGetCategoryList(t *testing.T) {
	categorylist,err :=GetAllCategoryList()
	if err != nil {
		t.Fatal(err)
		return
	}
	for _,cate := range categorylist{
		fmt.Println(cate.CategoryName)
	}
}
