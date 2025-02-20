package image_api

import (
	"blogx/common"
	"blogx/common/res"
	"blogx/global"
	"blogx/model"
	"blogx/service/log_service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ImageApi struct{}

type ImageListResponse struct {
	model.Image
	WebPath string `json:"webPath"`
}

func (ImageApi) ImageListView(c *gin.Context) {
	var cr common.PageInfo
	c.ShouldBindQuery(&cr)

	_list, count, _ := common.ListQuery(model.Image{}, common.Options{
		PageInfo: cr,
		Likes:    []string{"filename"},
	})
	var list = make([]ImageListResponse, 0)
	for _, model := range _list {
		list = append(list, ImageListResponse{
			Image:   model,
			WebPath: model.WebPath(),
		})
	}
	res.OkWithList(list, count, c)
}

func (ImageApi) ImageRemoveView(c *gin.Context) {
	var cr model.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var list []model.Image
	global.DB.Find(&list, "id in ?", cr.IDList)

	var successCount, errCount int64
	if len(list) > 0 {
		successCount = global.DB.Delete(&list).RowsAffected
	}
	errCount = int64(len(list)) - successCount

	msg := fmt.Sprintf("操作成功，成功%d 失败%d", successCount, errCount)
	res.OkWithMsg(msg, c)
}
