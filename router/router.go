package router

import (
	"github.com/caesar003/day-2-golang-praisindo-advanced-gin-crud/handler"
	"github.com/caesar003/day-2-golang-praisindo-advanced-gin-crud/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	userPublic := r.Group("/api/user")
	userPublic.GET("/", handler.GetUsers)
	userPublic.GET("/:id", handler.GetUser)

	userPrivate := r.Group("/api/user")
	userPrivate.Use(middleware.AuthMiddleWare())
	userPrivate.POST("/", handler.AddUser)
	userPrivate.PUT("/:id", handler.UpdateUser)
	userPrivate.DELETE("/:id", handler.DeleteUser)
}
