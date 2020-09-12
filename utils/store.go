package utils

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required,lte=40"`
	Email string `json:"email" binding:"required,email"`
}

type Todo struct {
	ID   string `json:"id"`
	Todo string `json:"todo" binding:"required"`
	Link string `json:"link" binding:"required,url"`
	UID  string `json:"uid" binding:"required"`
}

type Store struct {
	Users []User `json:"users"`
	Todos []Todo `json:"todos"`
}

var DBStore Store = Store{
	Users: []User{},
	Todos: []Todo{},
}
