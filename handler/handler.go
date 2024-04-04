package handler

import (
	"net/http"
	"simple-todo/server"

	"github.com/gin-gonic/gin"
)

type PostBody struct {
	Id      uint   `json:"id"`
	Content string `form:"content"`
}

func Core(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET,POST")

	c.Next()
}
func List(c *gin.Context) {
	res := server.GetList()
	c.HTML(http.StatusOK, "index.html", res)
}
func AllList(c *gin.Context) {
	res := server.GetAllList()
	c.HTML(http.StatusOK, "list.html", res)
}
func Add(c *gin.Context) {
	content := c.PostForm("content")
	server.AddTodo(content)
	c.Redirect(http.StatusMovedPermanently, "/")

}
func Edit(c *gin.Context) {
	var body PostBody
	c.ShouldBindJSON(&body)
	res := server.EditTodo(body.Id)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": res,
	})

}
func Delete(c *gin.Context) {
	var body PostBody
	c.ShouldBindJSON(&body)
	res := server.DeleteTodo(body.Id)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": res,
	})

}
