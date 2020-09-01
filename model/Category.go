package model

import (
	"blog/utils/errmsg"
)

type Category struct {
	Id uint       `gorm:"type int;not null;auto_increment" json:"id"`
	Name string   `gorm:"type varchar(20);not null"        json:"name"`
}

//查询用户是否存在
func DoesCategoryExist(name string) (code int) {
	var categ Category
	db.Select("id").Where("name=?", name).First(&categ)
	if categ.Id > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED
	}
	return errmsg.ERROR_CATEGORY_NOT_EXIST
}

//在数据库增加一条用户数据
func CreateCategory(categ *Category) (code int) {
	err2 := db.Create(categ).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}

//分页查询用户
func GetCategories(offset, pagesize int) []Category {
	var categs []Category

	if offset < 0 {
		offset = -1
	}
	if pagesize < -1 {
		pagesize = -1
	}
	db.Limit(pagesize).Offset((offset - 1) * pagesize).Find(&categs)

	return categs
}

//编辑用户
func UpdateCategory(id int, categ Category) (code int) {

	m := make(map[string]interface{}, 1)
	m["name"] = categ.Name

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	err := db.Model(&categ).Where("id=?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}

//删除用户
func DeleteCategory(id int) (code int) {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}