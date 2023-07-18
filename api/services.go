package api

import "net/http"

func (api *Api) InitServices() {
	api.BaseRoutes.Services.Handle("", api.APIHandler(getAllServices)).Methods("GET")
	api.BaseRoutes.Services.Handle("", api.APISessionRequired(createService)).Methods("POST")
}

func createService(c *Context, w http.ResponseWriter, r *http.Request) {

}

func getAllServices(c *Context, w http.ResponseWriter, r *http.Request) {

}
