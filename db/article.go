package db

import (
	"blogger/models"
	"fmt"
)

func GetArticleList(pageNum, pageSize int) (articleList []*models.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
					order by create_time desc
					limit ?, ?`
	err = DB.Select(&articleList, sqlstr, pageNum*pageSize, pageSize)
	return
}

func GetArticleListByCategoryId(categoryId int64, pageNum, pageSize int) (articleList []*models.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1 and category_id =?
					order by create_time desc
					limit ?, ?`
	err = DB.Select(&articleList, sqlstr, categoryId, pageNum*pageSize, pageSize)
	return
}

func GetArticleDetail(id int64) (detail models.ArticleDetail, err error) {
	sqlStr := `select article.id "article.id",content,category_id "article.category_id",summary "article.summary",
title "article.title",view_count "article.view_count",article.create_time "article.create_time",comment_count "article.comment_count",username "article.username",
category.id "category.id" ,category_name "category.category_name",category_no "category.category_no"
from article,category 
where article.id=? and article.category_id = category.id `

	err = DB.Get(&detail, sqlStr, id)
	if err != nil {
		return
	}
	return

}

func GetArticleinfoById(id int64) (articleInfo models.ArticleInfo, err error) {
	sqlStr := `select id,title from article where id=?`
	err =DB.Get(&articleInfo,sqlStr,id)
	return
}

func InsertArticle(categoryId int64,username,title,content string)(articleid int64,err error){
	sqlStr := `insert into article (category_id,title,content,username,view_count,comment_count,summary) values (?,?,?,?,0,0,?);`
	result,err :=DB.Exec(sqlStr,categoryId,title,content,username,title)
	if err != nil {
		fmt.Println(err)
		return 0,err
	}

	articleid,err = result.LastInsertId()
	return 
}

// 增加查看次数
func AddViewTime(article_id int64,view_times int){
	sqlStr := `update article set view_count = view_count+? where id = ?`
	DB.Exec(sqlStr,view_times,article_id)
}

// 增加评论次数
func AddCommentCount(article_id int64,count int){
	sqlStr := `update article set comment_count = comment_count+? where id = ?`
	DB.Exec(sqlStr,count,article_id)
}
