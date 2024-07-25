package dao

import (
	"app/dao/daoModel"
)

func InsertUser(user *daoModel.UserModel) {
	// 创建
	dbPool.Create(&user)
}

func GetUserByEmail(user *daoModel.UserModel) *[]daoModel.UserModel {
    var users []daoModel.UserModel
    result := dbPool.Raw("SELECT * FROM User WHERE email = ?", user.Email).Scan(&users)
    if result.Error != nil {
    }
	return &users
}

func UpdateUser(user *daoModel.UserModel) {
	id := user.Id
	updateValue := make(map[string]interface{})
	if (user.Username != "") {
		updateValue["username"] = user.Username
	}
	if (user.Password != "") {
		updateValue["username"] = user.Password
	}
	if (user.Email != "") {
		updateValue["username"] = user.Email
	}
	if (user.Phone != "") {
		updateValue["username"] = user.Phone
	}
	// 更新
	dbPool.Model(&daoModel.UserModel{}).Where("id = ?", id).Updates(updateValue)

}

func DeleteUser(user *daoModel.UserModel) {
	// 删除
	dbPool.Delete(&daoModel.UserModel{}, user.Id)
}
