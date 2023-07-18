package web

import (
	"net/http"
	"os"
	"path/filepath"
)

func (w *Web) InitStatic() {
	w.MainRouter.Handle("/", w.NewStaticHandler(root)).Methods("GET")
	w.MainRouter.Handle("/{anything:.*}", http.HandlerFunc(Handle404))
}

func root(c *Context, w http.ResponseWriter, r *http.Request) {
	contents, err := os.ReadFile(filepath.Join("./public/", "root.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(contents)
}
