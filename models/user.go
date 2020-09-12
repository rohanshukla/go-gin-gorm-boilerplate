package models

import (
	"fmt"
	"webapp/utils"
)

type User struct {
	Model
	Name  string `json:"name,omitempty" binding:"required,lte=40" gorm:"type:varchar(150)"`
	Email string `json:"email,omitempty" binding:"required,email" gorm:"type:varchar(256);unique;not null"`
}

func (user *User) Create() map[string]interface{} {
	result := GetDBInstance().Create(user)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"error": result.Error,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"message": "Created Successful",
	})
}

func GetUser(id int) map[string]interface{} {
	user := &User{}
	result := GetDBInstance().First(user, id)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"user": nil,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"user": user,
	})
}

func GetAllUser(fields []string) map[string]interface{} {
	users := make([]*User, 0)
	fmt.Printf("%v, %T\n", fields, fields)
	result := GetDBInstance().Select(fields).Find(&users)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"error": result.Error,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"users": users,
	})
}

func (user *User) Update() map[string]interface{} {
	result := GetDBInstance().Updates(user)
	if result.Error != nil {
		return utils.SendResponse(false, map[string]interface{}{
			"error": result.Error,
		})
	}
	return utils.SendResponse(true, map[string]interface{}{
		"message": "Updated Successful",
	})
}
