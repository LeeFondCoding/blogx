package log_api

import (
	"blogx/common"
	"blogx/common/res"
	"blogx/global"
	"blogx/model"
	"blogx/model/enum"
	"blogx/service/log_service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LogApi struct{}

type LogListRequest struct {
	common.PageInfo ``
	LogType         enum.LogType      `form:"logType"` // 日志类型 1 2 3
	Level           enum.LogLevelType `form:"level"`
	UserID          uint              `form:"userID"`
	IP              string            `form:"ip"`
	LoginStatus     bool              `form:"loginStatus"`
	ServiceName     string            `form:"serviceName"`
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

	list, count, err := common.ListQuery(model.Log{
		LogType:     cr.LogType,
		Level:       cr.Level,
		UserID:      cr.UserID,
		IP:          cr.IP,
		LoginStatus: cr.LoginStatus,
		ServiceName: cr.ServiceName,
	}, common.Options{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"title"},
		Preloads:     []string{"User"},
		DefaultOrder: "created_at desc",
	})

	var _list = make([]LogListResponse, 0)
	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			Log:          logModel,
			UserNickname: logModel.UserName,
			UserAvatar:   logModel.User.Avatar,
		})
	}

	res.OkWithList(_list, int(count), c)
}

func (LogApi) LogReadView(c *gin.Context) {
	var cr model.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var log model.Log
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的日志", c)
		return
	}
	if !log.IsRead {
		global.DB.Model(&log).Update("is_read", true)
	}
	res.OkWithMsg("日志读取成功", c)
}

func (LogApi) LogRemoveView(c *gin.Context) {
	var cr model.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var logList []model.Log
	global.DB.Find(&logList, "id in ?", cr.IDList)

	if len(logList) > 0 {
		global.DB.Delete(&logList)
	}

	msg := fmt.Sprintf("日志删除成功，共删除%d条", len(logList))

	res.OkWithMsg(msg, c)
}
