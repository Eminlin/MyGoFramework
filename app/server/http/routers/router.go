package routers

import (
	"MyGoFramework/common/log"

	"github.com/gorilla/mux"
)

type Router struct {
	Log *log.Log
}

func (r *Router) SetRouter() *mux.Router {

	//总路由
	router := mux.NewRouter()

	r.openAPI(router)

	return router
}
