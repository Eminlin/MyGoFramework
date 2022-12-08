package routers

import (
	"MyGoFramework/app/server/service"

	"github.com/gorilla/mux"
)

func (r *Router) openAPI(router *mux.Router) {
	open := router.PathPrefix("/api/open").Subrouter()
	open.HandleFunc("/hi", service.Hi)
}
