package repoModel

import "time"
import "app/dao/daoModel"

type UserInfo struct {
	Username   string
	Email      string
	Phone      string
	Password   string
	CreateTime time.Time
}

func NewUserInfo(username, email, phone, password string) *UserInfo {
	return &UserInfo{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
}

func ToUserModel(userInfo *UserInfo) *daoModel.UserModel {
	return daoModel.NewUserModel(userInfo.Username, userInfo.Email, userInfo.Password, userInfo.Phone)
}

func FromUserModel(userModel *daoModel.UserModel) *UserInfo {
	return NewUserInfo(userModel.Username, userModel.Email, userModel.Phone, userModel.Password)
}
