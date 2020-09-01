package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Secretkey string

	//qiniu storage parameters
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件路径错误", err)
		return
	}

	LoadServer(file)
	LoadDatabse(file)
	LoadJWT(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadDatabse(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("blog")
}

func LoadJWT(file *ini.File) {
	Secretkey = file.Section("jwt").Key("Secretkey").MustString("hello,world")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("fSCaBUu2JYm1eU3JIf0NG8VDsCyEUihIWouqutKd")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("yB00OMVxzPTLVCQhkPIm-hiU9tReYg5nTPZtIVZc")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("bloggerimg")
	QiniuServer = file.Section("qiniu").Key("qiniuServer").MustString("http://qfr31upvd.hd-bkt.clouddn.com/")
}
