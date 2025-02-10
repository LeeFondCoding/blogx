package flag

import (
	"blogx/global"
	"blogx/model"

	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&model.User{},
		&model.UserConf{},
		&model.Article{},
		&model.Category{},
		&model.ArticleDigg{},
		&model.Collect{},
		&model.UserArticleCollectModel{},
		&model.UserArticleLookHistory{},
		&model.Comment{},
		&model.Banner{},
		&model.Log{},
	)
	if err != nil {
		logrus.Errorf("数据迁移失败 %s", err)
		return
	}
	logrus.Infof("数据库迁移成功！")
}
