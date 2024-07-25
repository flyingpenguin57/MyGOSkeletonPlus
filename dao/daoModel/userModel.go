package daoModel

import (
	"time"
)

type UserModel struct {
	Id uint64 `gorm:"column:id"`
    Username string `gorm:"column:username"`
    Email  string `gorm:"column:email"`
    Password string `gorm:"column:password"`
	Phone string `gorm:"column:phone"`
	Created_at time.Time `gorm:"column:created_at"`
}


//在 Go 语言中，函数签名中的 (User) 是方法接收者（method receiver）。
//它指定了该方法是与 User 结构体类型关联的方法。换句话说，这个方法是 User 类型的一个方法，而不是普通的函数。
func (UserModel) TableName() string {
    return "User"
}

func NewUserModel(name, email, password, phone string) *UserModel {
    return &UserModel{
        Username:      name,
        Email:         email,
        Password:      password,
        Phone:         phone,
    }
}