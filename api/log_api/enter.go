package log_api

import (
	"blogx/common/res"
	"blogx/global"
	"blogx/model"
	"blogx/model/enum"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LogApi struct{}

type LogListRequest struct {
	Limit       int               `form:"limit"`
	Page        int               `form:"page"`
	Key         string            `form:"key"`
	LogType     enum.LogType      `form:"logType"` // 日志类型 1 2 3
	Level       enum.LogLevelType `form:"level"`
	UserID      uint              `form:"userID"`
	IP          string            `form:"ip"`
	LoginStatus bool              `form:"loginStatus"`
	ServiceName string            `form:"serviceName"`
}

type LogListResponse struct {
	model.Log
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	// 分页 查询（精确查询， 模糊匹配）
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr) 
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var list []model.Log
	if cr.Page> 20 {
		cr.Page = 1
	}
	if cr.Page <= 0 {
		cr.Page = 1
	}
	if cr.Limit == 0 || cr.Limit > 100 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	model := model.Log{
		LogType:     cr.LogType,
		Level:       cr.Level,
		UserID:      cr.UserID,
		IP:          cr.IP,
		LoginStatus: cr.LoginStatus,
		ServiceName: cr.ServiceName,
	}
	like := global.DB.Where("title like ?", fmt.Sprintf("%%%s%%", cr.Key))

	global.DB.Preload("User").Where(like).Where(model).Offset(offset).Limit(cr.Limit).Find(&list)

	var count int64
	global.DB.Where(like).Where(model).Model(model).Count(&count)

	var _list = make([]LogListResponse, 0)
	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			Log: logModel,
			UserNickname: logModel.UserName,
			UserAvatar: logModel.User.Avatar,
		})
	}

	 res.OkWithList(_list, int(count), c)
}