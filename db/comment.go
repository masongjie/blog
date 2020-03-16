package db

import "blogger/models"

// 插入单条评论
func InsertComment(content,username,email string,articleId int64)(err error){
	sqlStr := "insert into comment (content,username,status,article_id,email) values (?,?,1,?,?);"
	
	_,err =DB.Exec(sqlStr,content,username,articleId,email)
	if err != nil {
		return err
	}
	return 
}

// 展示评论
func GetCommentList(article_id int64)(commentList []*models.Comment,err error){
	sqlStr := `select content,create_time,username from comment where article_id =? order by create_time asc`
	err =DB.Select(&commentList,sqlStr,article_id)
	return
}