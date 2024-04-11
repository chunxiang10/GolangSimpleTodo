package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitSession() gin.HandlerFunc {
	store := cookie.NewStore(MyKey)
	return sessions.Sessions("session", store)
}

func DefaultSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
