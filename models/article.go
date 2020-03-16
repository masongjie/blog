package models

import "time"

type ArticleInfo struct {
	Id         int64 `db:"id"`
	CategoryId int64 `db:"category_id"`
	// 文章摘要
	Summary   string `db:"summary"`
	Title     string `db:"title"`
	ViewCount uint32 `db:"view_count"`
	// 时间
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// 用于文章详情页的实体
// 为了提升效率
type ArticleDetail struct {
	ArticleInfo `db:"article,prefix=article."`
	// 文章内容
	Content string `db:"content"`
	Category `db:"category,prefix=category."`
}

// 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}

