package api

import (
	"net/http"
)

func (api *Api) InitGroups() {
	api.BaseRoutes.Groups.Handle("", api.APIHandler(getAllGroups)).Methods("GET")
	api.BaseRoutes.Groups.Handle("", api.APISessionRequired(createGroup)).Methods("POST")
}

func createGroup(c *Context, w http.ResponseWriter, r *http.Request) {

}

func getAllGroups(c *Context, w http.ResponseWriter, r *http.Request) {

}
