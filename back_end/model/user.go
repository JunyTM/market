package model

type User struct {
	Id int	`json:"id" gorm:"primaryKey"`
	UserName	string `json:"userName" gorm:"user_name"`
	PassWord	string	`json:"passWord" gorm:"pass_word"`
	Role		string	`json:"role" gorm:"role"`

}

type UserResponse struct {
	Id int	`json:"id" gorm:"primaryKey"`
	UserName	string `json:"userName" gorm:"user_name"`
	Role		string	`json:"role" gorm:"role"`
}

type UserRepository interface {
	GetById(id int) (*UserResponse, error)
	GetAll() ([]UserResponse, error)
	Create(newUser User) (*User, error)
	Delete(id int) (*User, error)
}