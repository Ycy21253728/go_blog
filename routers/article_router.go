package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAdmin(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("article_id_title", app.ArticleIDTitleListView)
	router.GET("categorys", app.ArticleCategoryListView)
	router.GET("articles/detail", app.ArticleDetailByTitleView)
	router.GET("articles/calendar", app.ArticleCalendarView)
	router.GET("articles/tags", app.ArticleTagListView)
	router.PUT("articles", middleware.JwtAdmin(), app.ArticleUpdateView)
	router.DELETE("articles", middleware.JwtAdmin(), app.ArticleRemoveView)
	router.POST("articles/collects", middleware.JwtAuth(), app.ArticleCollCreateView)
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollListView)
	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	router.GET("articles/text", app.FullTextSearchView)
	router.POST("articles/digg", app.ArticleDiggView)
	router.GET("articles/content/:id", app.ArticleContentByIDView)
	router.GET("articles/:id", app.ArticleDetailView)
}
