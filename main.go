package main

import (
	"github.com/caesar003/day-2-golang-praisindo-advanced-gin-crud/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router.SetupRouter(r)

	r.Run(":9876")
}
