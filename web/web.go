package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inarithefox/status/app"
)

type Web struct {
	MainRouter *mux.Router
	server     *app.Server
}

func New(srv *app.Server) *Web {
	log.Println("initializing web server...")
	web := &Web{
		server:     srv,
		MainRouter: srv.Router,
	}

	web.InitStatic()

	return web
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	log.Printf("request received: %v, status: 404", r.URL.Path)
	http.NotFound(w, r)
}
