package web

import (
	"errors"

	"github.com/inarithefox/status/app"
	"github.com/inarithefox/status/app/request"
)

type Context struct {
	AppContext *request.Context
	App        *app.App
	Err        error
}

func (c *Context) MfaRequired() {
	if c.AppContext.Session().IsOAuth {
		return
	}
}

func (c *Context) SessionRequired() {
	if c.AppContext.Session().UserId == "" {
		c.Err = errors.New("user required")
		return
	}
}
