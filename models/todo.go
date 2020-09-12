package models

import "webapp/utils"

type Todo struct {
	Model
	Title  string `json:"title,omitempty" binding:"required" gorm:"type:varchar(255)"`
	Todo   string `json:"todo,omitempty" binding:"required"`
	Link   string `json:"link,omitempty"`
	User   User   `json:"user,omitempty" binding:"required" gorm:"foreignKey:UserID"`
	UserID uint64 `json:"-"`
}

func (todo *Todo) Create() map[string]interface{} {
	result := GetDBInstance().Create(todo)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"error": result.Error,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"message": "Created Successfully",
	})
}

func GetAllTodoByUser(userId uint, fields []string, limit int, offset int) map[string]interface{} {
	todos := make([]*Todo, 0)
	result := GetDBInstance().Joins("User").Select(fields).Where("user_id = ?", userId).Limit(limit).Offset(offset).Order("todos.created_at").Find(&todos)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"error": result.Error,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"todos": todos,
	})
}

func DeleteTodo(id int) map[string]interface{} {
	todo := &Todo{}
	result := GetDBInstance().Delete(todo, id)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"mesage": result.Error,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"message": "Deleted Successfully",
	})
}

func (todo *Todo) Update() map[string]interface{} {
	result := GetDBInstance().Save(todo)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"error": result.Error,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"message": "Updated Successful",
	})
}
