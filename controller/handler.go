package controller

import (
	"blogger/db"
	"blogger/log"
	"blogger/logic"
	"blogger/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

type comment struct {
	Comment string `form:"comment" binding:"required"`
	Username string `form:"author" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func Index(c *gin.Context) {
	pageStr := c.Query("page")
	var pageNum int64
	if pageStr == ""{
		pageStr = "0"
	}
	pageNum,err := strconv.ParseInt(pageStr,10,64)
	if err != nil {
		c.JSON(500,gin.H{"err":"page must be int"})
		return
	}
	
	categoryList, err := db.GetAllCategoryList()

	if err != nil {
		c.HTML(500, "views/500.html", gin.H{})
		return
	}
	article_list, err := logic.GetArticleRecord(int(pageNum), 15)

	if err != nil {
		c.HTML(500, "views/500.html", gin.H{})
		return
	}
	fmt.Println(article_list)
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"category_list": categoryList,
		"article_list":  article_list,
	})

}

func CategoryHandler(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.Redirect(301, "/")
		return
	}

	pageStr := c.Query("page")
	var pageNum int64
	if pageStr == ""{
		pageStr = "0"
	}
	pageNum,err = strconv.ParseInt(pageStr,10,64)
	if err != nil {
		c.JSON(500,gin.H{"err":"page must be int"})
		return
	}
	
	
	// 根据id查询分类对应的文章
	categoryList, err := db.GetAllCategoryList()
	if err != nil {
		c.HTML(200, "views/500.html", nil)
		return
	}
	article_list, err := logic.GetArticleRecordByCategoryId(categoryId, int(pageNum), 20)
	if err != nil {
		c.HTML(200, "views/500.html", nil)
		return
	}

	c.HTML(200, "views/index.html", gin.H{
		"category_list": categoryList,
		"article_list":  article_list,
	})

}

func DetailHandler(c *gin.Context) {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(500, "views/500.html", nil)
		return
	}

	log.SugarLog.Infow("article detail ", "articleID", articleId)
	// 需要查询三篇文章的信息，-1 +1
	// 当前文章展示细节
	detail, err := db.GetArticleDetail(articleId)
	if err != nil {
		log.SugarLog.Info("article is not existed")
		c.HTML(500, "views/500.html", nil)
		return
	}
	pre, next := logic.GetPrevNextArticleArticleInfo(articleId)

	//获取评论
	commentList,err:=db.GetCommentList(articleId)
	if err != nil {
		log.SugarLog.Infow("get commentlist failed","err",err)
		c.HTML(500, "views/500.html", nil)
		return
	}
	data := gin.H{
		"detail": detail,
		"prev":   pre,
		"next":   next,
		"comment_list":commentList,
		"content":template.HTML(detail.Content),
	}
	// 为增加一次阅读次数
	go db.AddViewTime(articleId,1)
	c.HTML(200, "views/detail.html", data)
}

func NewArticle(c *gin.Context) {
	categorylist, err := db.GetAllCategoryList()
	if err != nil {
		log.SugarLog.Info("get category failed")
		c.HTML(500, "views/500.html", nil)
		return
	}
	c.HTML(200, "views/post_article.html", categorylist)
}

func AddArticle(c *gin.Context) {
	// 获取参数 body
	// 字段 author title category_id content
	userName := c.PostForm("author")
	title := c.PostForm("title")
	categoryIdStr := c.PostForm("category_id")
	content := c.PostForm("content")

	// 接收md文件，保存到load目录下
	file, err := c.FormFile("file")
	// 没有上传文件的话
	if err == nil {
		//log.SugarLog.Infow("upload file file,","filename",file.Filename)
		fileName := fmt.Sprintf("./loads/%s",file.Filename)

		err:=c.SaveUploadedFile(file,fileName)
		if err != nil {
			log.SugarLog.Info("file save failed ","err",err)
			c.String(501, "file save failed")
			return
		}
		//修改content
		content = file.Filename
	}


	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		log.SugarLog.Info("categoryId must be int")
		c.String(501, "categoryId must be int")
		return
	}

	articleId, err := db.InsertArticle(categoryId, userName, title, content)
	if err != nil {
		log.SugarLog.Infow("insert article failed", "err", err)
		c.HTML(500, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/article/detail?article_id=%d", articleId)
	c.Redirect(http.StatusFound, url)
	return
}

type LoginForm struct {
	User     string `form:"username" binding:"required"`
	Password string `form:"pwd" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	targetUrl := c.DefaultQuery("target","/")
	fmt.Println(targetUrl)
	token, _ := c.Cookie("token")

	if c.Request.Method != "POST" {
		c.HTML(200, "views/login.html", nil)
		return
	} else {
		var form LoginForm
		err := c.ShouldBind(&form)
		if err != nil {
			log.SugarLog.Infow("get login postfrom failed", "err", err)
			c.Redirect(302, "/")
			return
		}
		if form.User == "masongjie" && form.Password == "123456" {
			fmt.Println("login success")
			// 登录成功 设置session
			ok := utils.SetIsLogin(token)
			fmt.Println(ok)

			if !ok {
				log.SugarLog.Infow("redis store failed", "err", err)
				c.HTML(500, "views/500.html", nil)
				return
			}
			c.SetCookie("token", token, 3600, "/", "/", false, false)
			c.Redirect(302, targetUrl)
			return
		} else {
			c.JSON(401, gin.H{"status": "unautherized"})
			return
		}
	}

}


// 添加一条评论
func SubmitComment(c *gin.Context){
	var com1 comment
	article_id := c.Query("article_id")
	url := fmt.Sprintf("/article/detail?article_id=%s",article_id)
	articleId,err:=strconv.ParseInt(article_id,10,64)
	if err != nil {
		c.JSON(501,gin.H{
			"err":"article_id must be int",
		})
		return
	}


	err =c.ShouldBind(&com1)
	if err != nil {
		fmt.Println(err)
		c.HTML(500,"views/500.html",nil)
		return
	}



	err =db.InsertComment(com1.Comment,com1.Username,com1.Email,articleId)
	if err != nil {
		fmt.Println(err)
		c.HTML(500,"views/500.html",nil)
		return
	}
	go db.AddCommentCount(articleId,1)
	c.Redirect(302,url)

}

// 增加一个分类
func AddCategory(c *gin.Context){
	categoryName :=c.PostForm("category_name")
	if categoryName ==""{
		c.JSON(200,gin.H{"err":"category name must be existed"})
		return
	}
	go db.InsertCategory(categoryName)
	c.Redirect(301,"/article/new/")
}
