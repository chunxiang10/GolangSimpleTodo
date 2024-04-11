package middlers

import (
	"net/http"
	"simple-todo/handler"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}
		authString := strings.Split(authHeader, " ")
		if !(len(authString) == 2 && authString[0] == "Bearer") {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}
		uid, err := handler.ParseToken(authString[1])
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}
