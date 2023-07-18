package api

import (
	"net/http"

	"github.com/inarithefox/status/web"
)

type Context = web.Context

type handlerFunc func(*Context, http.ResponseWriter, *http.Request)

func (api *Api) APIHandler(h handlerFunc) http.Handler {
	handler := &web.Handler{
		Server:         api.server,
		HandlerFunc:    h,
		HandlerName:    web.GetHandlerName(h),
		RequireSession: false,
		RequireMfa:     false,
		IsStatic:       false,
	}

	return handler
}

func (api *Api) APISessionRequired(h handlerFunc) http.Handler {
	handler := &web.Handler{
		Server:         api.server,
		HandlerFunc:    h,
		HandlerName:    web.GetHandlerName(h),
		RequireSession: true,
		RequireMfa:     true,
		IsStatic:       false,
	}

	return handler
}
