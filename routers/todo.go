package routers

import (
	"strconv"
	"webapp/models"
	"webapp/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTodo(c *gin.Context) {
	var requestBody utils.Todo
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     "Invalid request body",
		})
		return
	}

	requestBody.ID = uuid.New().String()
	utils.DBStore.Todos = append(utils.DBStore.Todos, requestBody)
	c.JSON(201, gin.H{
		"success": true,
	})
}

func GetUserTodo(c *gin.Context) {
	uid := c.Param("uid")
	todos := []utils.Todo{}
	for _, item := range utils.DBStore.Todos {
		if item.UID == uid {
			todos = append(todos, item)
		}
	}

	c.JSON(200, gin.H{
		"success": true,
		"todos":   todos,
	})
}

/* Database API Calls */
func CreateTodoDB(c *gin.Context) {
	todo := &models.Todo{}
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     err.Error(),
		})
		return
	}
	result := todo.Create()
	if result["success"] == true {
		c.JSON(201, result)
	} else {
		c.JSON(200, result)
	}
}

func GetAllTodoByUser(c *gin.Context) {
	// Read request body using below struct
	type Request struct {
		Fields []string `json:"fields" binding:"required"`
		Limit  int      `json:"limit" binding:"required"`
		Offset int      `json:"offset" binding:"required"`
		UserID uint     `json:"user_id" binding:"required"`
	}
	request := &Request{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     err.Error(),
		})
		return
	}

	result := models.GetAllTodoByUser(request.UserID, request.Fields, request.Limit, request.Offset)
	if result["success"] == true {
		c.JSON(200, result)
	} else {
		c.JSON(404, result)
	}
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	intVal, _ := strconv.Atoi(id)
	result := models.DeleteTodo(intVal)
	if result["success"] == true {
		c.JSON(200, result)
	} else {
		c.JSON(404, result)
	}
}

func UpdateTodo(c *gin.Context) {
	id := c.Query("id")
	intVal, _ := strconv.Atoi(id)

	todo := &models.Todo{}
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"msg":     err.Error(),
		})
		return
	}
	todo.ID = uint(intVal)
	result := todo.Update()
	if result["success"] == true {
		c.JSON(200, result)
	} else {
		c.JSON(404, result)
	}
}
