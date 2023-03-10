package global

import (
	"github.com/eroneko/gin-service/pkg/settings"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *settings.ServerSettingS
	AppSetting      *settings.AppSettingS
	DatabaseSetting *settings.DatabaseSettingS
	DBEngine        *gorm.DB
)
