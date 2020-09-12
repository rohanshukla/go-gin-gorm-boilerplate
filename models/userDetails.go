package models

type UserDetails struct {
	Model
	Mobile   string `json:"mobile" binding:"required,lte=10,gte=10" gorm:"type:varchar(10)"`
	Age      int    `json:"age"`
	Verified bool   `json:"verified" binding:"required"`
	User     User   `json:"uid" binding:"required" gorm:"foreignKey:UserID"`
	UserID   uint64 `json:"-"`
}
