package bootstrap

import (
	"MyGoFramework/app/cron"
	"MyGoFramework/app/server/http/code"
	"MyGoFramework/common/cache"
	"MyGoFramework/common/conf"
	"MyGoFramework/common/database"
	"MyGoFramework/common/log"
	"MyGoFramework/common/tools"
)

func Init() {
	log.InitLog()
	conf.Init()
	database.Init()
	cache.Init()
	code.I18n()
	tools.Go(func() { cron.NewCrontab().InitTask().Start() })
}
