package middlers

import (
	"net/http"
	"simple-todo/handler"

	"github.com/gin-gonic/gin"
)

// func SessionMiddleware(c *gin.Context) {
func SessionMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := handler.DefaultSession(c)
		var strToken string
		token := session.Get("token")
		if token == nil {
			// Abort the request with the appropriate error code
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}
		strToken = token.(string)
		_, err := handler.ParseToken(strToken)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}
		// Continue down the chain to handler etc
		c.Next()
	}
}
