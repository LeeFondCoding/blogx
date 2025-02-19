package image_api

import (
	"blogx/common/res"
	"blogx/global"
	"blogx/model"
	"blogx/util"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	// 文件大小判断
	s := global.Conf.Upload.Size
	if fileHeader.Size > s*1024*1024 {
		res.FailWithMsg(fmt.Sprintf("文件大小大于%dMB", s), c)
		return
	}
	// 后缀判断
	filename := fileHeader.Filename
	suffix, err := imageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	// 文件hash
	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	byteData, _ := io.ReadAll(file)
	hash := util.MD5(byteData)
	// 判断这个hash有没有
	var image model.Image
	err = global.DB.Take(&image, "hash = ?", hash).Error
	if err == nil {
		// 找到了
		logrus.Infof("上传图片重复 %s <==> %s  %s", filename, image.FileName, hash)
		res.Ok(image.WebPath(), "上传成功", c)
		return
	}

	// 文件名称一样，但是文件内容不一样

	filePath := fmt.Sprintf("uploads/%s/%s.%s", global.Conf.Upload.UploadDir, hash, suffix)
	// 入库
	image = model.Image{
		FileName: filename,
		Path:     filePath,
		Size:     fileHeader.Size,
		Hash:     hash,
	}
	err = global.DB.Create(&image).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	c.SaveUploadedFile(fileHeader, filePath)
	res.Ok(image.WebPath(), "图片上传成功", c)

}

func imageSuffixJudge(filename string) (suffix string, err error) {
	_list := strings.Split(filename, ".")
	if len(_list) == 1 {
		err = errors.New("错误的文件名")
		return
	}
	// xxx.jpg   xxx  xxx.jpg.exe
	suffix = _list[len(_list)-1]
	if !util.InList(suffix, global.Conf.Upload.WhiteList) {
		err = errors.New("文件非法")
		return
	}
	return
}