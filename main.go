package main

import (
	"first-gin/service"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/users", service.FetchAll)
	router.POST("", service.Insert)

	router.Run(":8080")
}
