package routers

import (
	"strconv"
	"webapp/models"
	"webapp/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"users":   utils.DBStore.Users,
	})
}

func CreateUser(c *gin.Context) {
	var requestBody utils.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     err.Error(),
		})
		return
	}

	requestBody.ID = uuid.New().String()
	utils.DBStore.Users = append(utils.DBStore.Users, requestBody)
	c.JSON(201, gin.H{
		"success": true,
	})
}

/* Database API Calls */
func CreateUserDB(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     err.Error(),
		})
		return
	}
	result := user.Create()
	if result["success"] == true {
		c.JSON(201, result)
	} else {
		c.JSON(200, result)
	}
}

func GetUserDB(c *gin.Context) {
	id := c.Param("id")
	// intVal, _ := strconv.ParseUint(id, 10, 64)
	intVal, _ := strconv.Atoi(id)
	result := models.GetUser(intVal)
	if result["success"] == true {
		c.JSON(200, result)
	} else {
		c.JSON(404, result)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	intVal, _ := strconv.Atoi(id)

	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     err.Error(),
		})
		return
	}
	user.ID = uint(intVal)
	result := user.Update()
	if result["success"] == true {
		c.JSON(200, result)
	} else {
		c.JSON(404, result)
	}
}

func GetAllUsersDB(c *gin.Context) {
	// Read request body using below struct
	type Request struct {
		Fields []string `json:"fields" binding:"required"`
	}
	request := &Request{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     err.Error(),
		})
		return
	}
	result := models.GetAllUser(request.Fields)
	if result["success"] == true {
		c.JSON(200, result)
	} else {
		c.JSON(404, result)
	}
}
