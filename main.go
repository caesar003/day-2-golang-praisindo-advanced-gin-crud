package main

import (
	"github.com/caesar003/golang-praisindo-advanced/middleware"
	"github.com/caesar003/golang-praisindo-advanced/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(middleware.AuthMiddleWare())

	router.SetupRouter(r)

	r.Run(":9876")
}
