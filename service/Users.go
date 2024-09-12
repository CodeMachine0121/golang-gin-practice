package service

import (
	"first-gin/models"
	"github.com/gin-gonic/gin"
)

var useList = []models.User{
	{ID: 1, Name: "John", Email: "john@mail"},
	{ID: 2, Name: "Doe", Email: "doe@mail"},
}

func FetchAll(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": useList,
	})
}

func Insert(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	useList = append(useList, user)
	c.JSON(200, gin.H{
		"message": useList,
	})
}
