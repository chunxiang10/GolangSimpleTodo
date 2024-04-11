package handler

import (
	"net/http"

	"simple-todo/server"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", "")
}

func CheckLogin(c *gin.Context) {
	session := DefaultSession(c)
	var body struct {
		Uid string `json:"uid"`
		Pwd string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	uid := body.Uid
	pwd := body.Pwd
	check := server.CheckLogin(uid, pwd)
	token, e := SendToken(uid)
	if check != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "faild",
		})
		return
	}
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "faild",
		})
		return
	}
	session.Set("token", token)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})

}
