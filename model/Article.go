package model

import (
	"blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:cid"`
	gorm.Model
	Title   string `gorm:"type varchar(80);not null"   json:"title"`
	Cid     int    `gorm:"type int;not null"           json:"cid"`
	Desc    string `gorm:"type varchar(200); not null" json:"desc"`
	Content string `gorm:"type longtext;not null"      json:"content"`
	Img     string `gorm:"type varchar(100)"           json:"img"`
}

//在数据库增加一条文章数据
func CreateArticle(article *Article) (code int) {
	err2 := db.Create(article).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}

//查询单个文章
func GetOneArticle(id int) (article Article, code int) {
	err := db.Preload("Category").Where("id=?", id).Find(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICE_NOT_EXIST
	}
	return article, errmsg.Success
}

//查询分类下的所有文章
func GetArticlesUnderCategory(cid, offset, pagesize int) (articles []Article, code int) {
	if offset < 0 {
		offset = -1
	}
	if pagesize < 0 {
		pagesize = -1
	}
	err := db.Preload("Category").Limit(pagesize).
		Offset((offset - 1) * pagesize).
		Where("cid=?", cid).
		Find(&articles).Error
	if err != nil {
		return articles, errmsg.ERROR
	}
	return articles, errmsg.Success
}

//分页查询文章
func GetArticles(offset, pagesize int) []Article {
	var articles []Article

	if offset < 0 {
		offset = -1
	}
	if pagesize < -1 {
		pagesize = -1
	}
	db.Preload("Category").Limit(pagesize).Offset((offset - 1) * pagesize).Find(&articles)

	return articles
}

//编辑文章
func UpdateArticle(id int, article Article) (code int) {

	m := make(map[string]interface{}, 1)
	m["title"] = article.Title
	m["Cid"] = article.Cid
	m["Desc"] = article.Desc
	m["Content"] = article.Content
	m["Img"] = article.Img

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	err := db.Model(&article).Where("id=?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}

//删除文章
func DeleteArticle(id int) (code int) {

	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.Success
}
