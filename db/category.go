package db

import (
	"blogger/models"
	"github.com/jmoiron/sqlx"
)

func GetCategory(id int64)(category models.Category,err error){
	sqlStr := "select id,category_name,category_no from category where id =?"
	err = DB.Get(&category,sqlStr,id)
	return
}


func GetAllCategoryList()(categotyList []*models.Category,err error){
	sqlStr := "select id ,category_name from category order by category_no asc"
	err = DB.Select(&categotyList,sqlStr)
	return
}

func GetCategoryList(categoryids []int64)(categoryList []*models.Category,err error){
	sqlStr,args,err:=sqlx.In("select id,category_name,category_no from category where id in (?)",categoryids)
	if err != nil {
		return
	}
	DB.Select(&categoryList,sqlStr,args...)
	return

}

func InsertCategory(categoryName string){
	sqlStr := `insert into category (category_name,category_no) value (?,0);`
	DB.Exec(sqlStr,categoryName)
}


