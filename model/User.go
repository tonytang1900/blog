package model

import (
	"blog/utils/errmsg"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=6,max=20" label:"用户名"`
	Password string `gorm:"type:varchar(50);not null" json:"password" validate:"required,min=6,max=30" label:"密码"`
	Role     int    `gorm:"type:int;default:2"        json:"role"     validate:"required,gte=2"        label:"角色"`
}

//查询用户是否存在
func DoesUserExist(name string) (code int) {
	var u User
	db.Select("id").Where("username=?", name).First(&u)
	if u.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.ERROR_USER_NOT_EXIST
}

//在数据库增加一条用户数据
func CreateUser(user *User) (code int) {
	//加密明文密码
	user.Password = cryptPWD(user.Password)

	err2 := db.Create(user).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}

//分页查询用户
func GetUsers(offset, pagesize int) []User {
	var users []User

	if offset < 0 {
		offset = -1
	}
	if pagesize < -1 {
		pagesize = -1
	}
	db.Limit(pagesize).Offset((offset - 1) * pagesize).Find(&users)

	return users
}

//对密码做hash处理，然后存base64格式
func cryptPWD(password string) string {
	salt := []byte{10, 20, 30, 49, 55, 68}

	key, _ := scrypt.Key([]byte(password), salt, 32768, 8, 1, 8)
	base64PWD := base64.StdEncoding.EncodeToString(key)
	return base64PWD
}

//删除用户
func DeleteUser(id int) (code int) {

	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}

//编辑用户
func UpdateUser(id int, user User) (code int) {
	fmt.Println("=========")

	m := make(map[string]interface{}, 1)
	m["username"] = user.Username
	m["role"] = user.Role

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	err := db.Model(&user).Where("id=?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}

//登录验证用户名和密码
func ValidateLogin(username, password string) (code int) {
	var user User
	err := db.Where("username=?", username).Find(&user).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if cryptPWD(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 2 {
		return errmsg.ERROR_USER_NOT_AUTHORIZED
	}
	return errmsg.Success
}
