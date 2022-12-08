package main

import (
	"MyGoFramework/bootstrap"
	"MyGoFramework/common/conf"
	"MyGoFramework/common/log"
	"net/http"

	"MyGoFramework/app/server/http/routers"
)

func main() {
	bootstrap.Init()
	r := routers.Router{Log: log.NewLog()}
	log.NewLog().Infoln("starting server at:", conf.AppConf.App.HostPort)
	if err := http.ListenAndServe(conf.AppConf.App.HostPort, r.SetRouter()); err != nil {
		log.NewLog().Errorln(err)
	}
}
