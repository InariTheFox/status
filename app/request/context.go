package request

import (
	"context"

	"github.com/inarithefox/status/models"
)

type Context struct {
	session        models.Session
	requestId      string
	ipAddress      string
	path           string
	userAgent      string
	acceptLanguage string

	context context.Context
}

func NewContext(ctx context.Context, requestId, ipAddress, path, userAgent, acceptLanguage string, session models.Session) *Context {
	return &Context{
		session:        session,
		requestId:      requestId,
		ipAddress:      ipAddress,
		path:           path,
		userAgent:      userAgent,
		acceptLanguage: acceptLanguage,
		context:        ctx,
	}
}

func (c *Context) Session() *models.Session {
	return &c.session
}

func (c *Context) SetAcceptLanguage(s string) {
	c.acceptLanguage = s
}

func (c *Context) SetIpAddress(s string) {
	c.ipAddress = s
}

func (c *Context) SetPath(s string) {
	c.path = s
}

func (c *Context) SetRequestId(s string) {
	c.requestId = s
}

func (c *Context) SetUserAgent(s string) {
	c.userAgent = s
}
