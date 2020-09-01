package routers

import (
	"blog/api/v1"
	"blog/middleware"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouters()  {
	//设置gin框架为开发调试模式，或是发布模式
	gin.SetMode(utils.AppMode)

	//gin.New()
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())
	engine.Use(middleware.NewCors())

	auth := engine.Group("/api/v1")
	auth.Use(middleware.VerifyJWT())
	{
		//user模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		//article模块的路由接口
		auth.POST("article/add", v1.AddArticle)

		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		//category模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		//上传文件
		auth.POST("upload", v1.Upload)
	}

	r := engine.Group("/api/v1")
	{
		r.GET("users", v1.ListAllUsers)
		r.POST("user/add", v1.AddUser)

		r.GET("articles", v1.ListArticles)
		r.GET("article/:id", v1.ListOneArticle)
		r.GET("articles/category/:cid", v1.ListArticlesUnderCategory)

		r.GET("categories", v1.ListCategories)


		r.POST("login", v1.Login)
	}


	err := engine.Run(utils.HttpPort)
	if err != nil {
		panic("运行作物")
	}

}
