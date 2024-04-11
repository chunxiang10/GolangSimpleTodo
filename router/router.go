package router

import (
	"net/http"

	"simple-todo/handler"
	"simple-todo/middlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(handler.InitSession())
	r.StaticFS("/static", http.Dir("./assets"))
	r.StaticFile("/favicon.ico", "./assets/img/favicon.ico")
	r.LoadHTMLGlob("template/*")
	r.Use(handler.Core)
	r.GET("/login", handler.Login)
	r.POST("/login", handler.CheckLogin)
	newr := r.Group("/").Use(middlers.SessionMiddleware())
	{
		newr.GET("/list", handler.AllList)
		newr.GET("/", handler.List)
		newr.POST("/", handler.Add)
		newr.POST("/edit", handler.Edit)
		newr.POST("/del", handler.Delete)
	}
	return r
}
