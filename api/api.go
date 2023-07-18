package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inarithefox/status/app"
	"github.com/inarithefox/status/web"
)

type Api struct {
	BaseRoutes *Routes
	server     *app.Server
}

type Routes struct {
	Root    *mux.Router
	ApiRoot *mux.Router

	Services *mux.Router
	Service  *mux.Router
	Groups   *mux.Router
	Group    *mux.Router
}

func Init(srv *app.Server) (*Api, error) {
	log.Println("initializing api...")

	api := &Api{
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = srv.Router
	api.BaseRoutes.ApiRoot = srv.Router.PathPrefix("/api").Subrouter()
	api.BaseRoutes.Services = api.BaseRoutes.ApiRoot.PathPrefix("/services").Subrouter()
	api.BaseRoutes.Service = api.BaseRoutes.ApiRoot.PathPrefix("/services/{service_id:[a-zA-Z0-9]+}").Subrouter()
	api.BaseRoutes.Groups = api.BaseRoutes.ApiRoot.PathPrefix("/groups").Subrouter()
	api.BaseRoutes.Group = api.BaseRoutes.ApiRoot.PathPrefix("/groups/{service_id:[a-zA-Z0-9]+}").Subrouter()

	api.InitServices()
	api.InitGroups()

	// TODO: Handle 404
	srv.Router.Handle("/api/{anything:.*}", http.HandlerFunc(api.Handle404))

	return api, nil
}

func (api *Api) Handle404(w http.ResponseWriter, r *http.Request) {
	web.Handle404(w, r)
}
