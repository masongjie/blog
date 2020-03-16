package logic

import (
	"blogger/db"
	"blogger/models"
)

// 获取文章列表

func GetArticleRecord(pageNum int,pageSize int)(articleReocrdList []*models.ArticleRecord,err error){
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 返回页面，做聚合
	// 遍历所有文章
	for _, article := range articleInfoList {
		// 根据当前文章，生成结构体
		articleRecord := &models.ArticleRecord{
			ArticleInfo: *article,
		}
		// 文章取出分类id
		categoryId := article.CategoryId
		// 遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleReocrdList = append(articleReocrdList, articleRecord)
	}
	return
}


func getCategoryIds(articleInfoList []*models.ArticleInfo)(ids []int64){
	LABEL:
		for _,articleInfo := range articleInfoList{
			categoryId := articleInfo.CategoryId
			for _,id := range ids{
				if id == categoryId{
					continue LABEL
				}
			}
			ids = append(ids, int64(categoryId))
		}
		return
}

//根据categoryid查询文章
func GetArticleRecordByCategoryId(categoryId int64,pageNum int,pageSize int)(articleRecordlist []*models.ArticleRecord,err error){
	articleList,err:=db.GetArticleListByCategoryId(categoryId,pageNum,pageSize)
	if err != nil {
		return 
	}
	// 查询单个category
	category,err:=db.GetCategory(categoryId)
	if err != nil {
		return 
	}
	for _,article := range articleList{
		articleRcord := &models.ArticleRecord{
			ArticleInfo: *article,
			Category:    category,
		}
		articleRecordlist =append(articleRecordlist, articleRcord)
		
	}
	return 
	
}


func GetPrevNextArticleArticleInfo(id int64)(pre,next models.ArticleInfo){
	pre,err := db.GetArticleinfoById(id-1)
	if err != nil {
		pre = models.ArticleInfo{
			Id: int64(-1),
		}
	}
	next,err =db.GetArticleinfoById(id+1)
	if err != nil {
		next = models.ArticleInfo{
			Id: int64(-1),
		}
	}
	return 
}