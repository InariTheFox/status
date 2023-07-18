package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/inarithefox/status/utils"
)

type Server struct {
	RootRouter    *mux.Router
	Router        *mux.Router
	Server        *http.Server
	ListenAddress *net.TCPAddr

	app             *App
	didFinishListen chan struct{}
}

func NewServer() (*Server, error) {
	rootRouter := mux.NewRouter()

	app := New()

	s := &Server{
		RootRouter: rootRouter,
		app:        app,
	}

	s.Router = s.RootRouter.PathPrefix("/").Subrouter()

	log.Printf("current version %f", 1.0)

	return s, nil
}

func (s *Server) Shutdown() {
	s.StopHTTPServer()
	log.Println("server stopped")
}

func (s *Server) Start() error {
	log.Println("starting server...")

	var handler http.Handler = s.RootRouter

	s.Server = &http.Server{
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 30,
	}

	addr := utils.Params.GetString("SERVER_PORT")
	if !strings.HasPrefix(addr, ":") {
		if addr == "" {
			addr = ":http"
		} else {
			addr = ":" + addr
		}
	}

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	s.ListenAddress = listener.Addr().(*net.TCPAddr)

	log.Printf("server is listening on %v", listener.Addr().String())

	s.didFinishListen = make(chan struct{})
	go func() {
		s.Server.Serve(listener)
		close(s.didFinishListen)
	}()

	return nil
}

func (s *Server) StopHTTPServer() {
	if s.Server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		didShutdown := false
		for s.didFinishListen != nil && !didShutdown {
			if err := s.Server.Shutdown(ctx); err != nil {
				log.Printf("unable to shutdown server %v", err)
			}

			timer := time.NewTimer(time.Millisecond * 50)
			select {
			case <-s.didFinishListen:
				didShutdown = true
			case <-timer.C:
			}

			timer.Stop()
		}

		s.Server.Close()
		s.Server = nil
	}
}
