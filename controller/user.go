package controller

import (
	"net/http"
	"strconv"
	"../structs"
	"github.com/gin-gonic/gin"
)

// create new user data
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user structs.User
		result gin.H
	)
	username := c.PostForm("UserName")
	parent := c.PostForm("Parent")
	parentInt, _ := strconv.Atoi(parent)
	user.UserName = username
	user.Parent = parentInt

	idb.DB.Create(&user)
	result = gin.H {
		"result": user,
	}
	c.JSON(http.StatusOK, result)
}