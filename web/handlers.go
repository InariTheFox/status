package web

import (
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/inarithefox/status/app"
	"github.com/inarithefox/status/app/request"
	"github.com/inarithefox/status/models"
	"github.com/inarithefox/status/utils"
)

type Handler struct {
	Server         *app.Server
	HandlerFunc    func(*Context, http.ResponseWriter, *http.Request)
	HandlerName    string
	RequireSession bool
	RequireMfa     bool
	IsStatic       bool
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode        int
	statusCodeWritten bool
	flusher           http.Flusher
}

func GetHandlerName(h func(*Context, http.ResponseWriter, *http.Request)) string {
	handlerName := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
	pos := strings.LastIndex(handlerName, ".")
	if pos != -1 && len(handlerName) > pos {
		handlerName = handlerName[pos+1:]
	}

	return handlerName
}

func (w *Web) NewStaticHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		Server:         w.server,
		HandlerFunc:    h,
		HandlerName:    GetHandlerName(h),
		RequireSession: false,
		RequireMfa:     false,
		IsStatic:       true,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = newWrappedWriter(w)
	//now := time.Now()

	app := app.New()
	requestId := models.NewId()
	var statusCode string
	defer func() {
		log.Printf("request received [%v]: %v, method: %v, status code: %v", requestId, r.URL.Path, r.Method, statusCode)
	}()

	c := &Context{
		AppContext: &request.Context{},
		App:        app,
	}
	c.AppContext.SetRequestId(requestId)
	c.AppContext.SetIpAddress(utils.GetIPAddress(r))
	c.AppContext.SetUserAgent(r.UserAgent())
	c.AppContext.SetAcceptLanguage(r.Header.Get("Accept-Language"))
	c.AppContext.SetPath(r.URL.Path)

	if c.Err == nil && h.RequireSession {
		c.SessionRequired()
	}

	if c.Err == nil && h.RequireMfa {
		c.MfaRequired()
	}

	if c.Err == nil {
		h.HandlerFunc(c, w, r)
	}

	statusCode = strconv.Itoa(w.(*responseWriterWrapper).StatusCode())
}

func newWrappedWriter(o http.ResponseWriter) *responseWriterWrapper {
	flusher, _ := o.(http.Flusher)

	return &responseWriterWrapper{
		ResponseWriter:    o,
		statusCodeWritten: false,
		flusher:           flusher,
	}
}

func (rw *responseWriterWrapper) StatusCode() int {
	return rw.statusCode
}

func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.statusCodeWritten = true
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriterWrapper) Write(data []byte) (int, error) {
	if !rw.statusCodeWritten {
		rw.statusCode = http.StatusOK
	}
	return rw.ResponseWriter.Write(data)
}

func (rw *responseWriterWrapper) Flush() {
	if rw.flusher != nil {
		rw.flusher.Flush()
	}
}
