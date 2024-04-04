package router

import (
	"net/http"

	"simple-todo/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./assets"))
	r.LoadHTMLGlob("template/*")
	r.Use(handler.Core)
	r.GET("/", handler.List)
	r.GET("/list", handler.AllList)
	r.POST("/", handler.Add)
	r.POST("/edit", handler.Edit)
	r.POST("/del", handler.Delete)
	return r
}
