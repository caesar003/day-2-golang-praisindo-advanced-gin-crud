package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/caesar003/day-2-golang-praisindo-advanced-gin-crud/entity"
	"github.com/gin-gonic/gin"
)

var (
	users  []entity.User
	nextID int
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is running",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var user *entity.User

	for _, u := range users {
		if u.ID == id {
			user = &u
			break
		}
	}

	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)

}

func AddUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = nextID
	nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser entity.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user *entity.User
	for i, u := range users {
		if u.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Password = updatedUser.Password
			users[i].Email = updatedUser.Email
			users[i].UpdatedAt = time.Now()
			user = &users[i]
			break
		}
	}

	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var index int
	var user *entity.User
	for i, u := range users {
		if u.ID == id {
			user = &u
			index = i
			break
		}
	}

	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	users = append(users[:index], users[index+1:]...)

	c.JSON(200, gin.H{"message": "User deleted"})
}
