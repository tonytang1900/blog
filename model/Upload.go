package model

import (
	"blog/utils"
	"blog/utils/errmsg"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"mime/multipart"
)

func UploadFile(file multipart.File, fsize int64) (url string, code int) {

	putPolicy := storage.PutPolicy{
		Scope: utils.Bucket,
	}
	mac := qbox.NewMac(utils.AccessKey, utils.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone: &storage.ZoneHuadong,
		UseHTTPS: false,
		UseCdnDomains: false,
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fsize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "",errmsg.ERROR
	}

	fmt.Println("=================", ret.Key)
	imgUrl := utils.QiniuServer + ret.Key
	return imgUrl, errmsg.Success
}
