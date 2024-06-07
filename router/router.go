package router

import (
	"github.com/caesar003/golang-praisindo-advanced/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)
	r.POST("/", handler.PostHandler)

	r.GET("/api/user", handler.RootHandler)
	r.POST("/api/user", handler.RootHandler)
	r.PUT("/api/user", handler.RootHandler)
	r.DELETE("/api/user", handler.RootHandler)

}
