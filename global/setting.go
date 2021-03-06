package global

import (
	"github.com/nathan-tw/blog/pkg/logger"
	"github.com/nathan-tw/blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
