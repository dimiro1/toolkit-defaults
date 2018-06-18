package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	mux *mux.Router
}

func New() *Router {
	return &Router{
		mux: mux.NewRouter(),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("GET")
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("POST")
}

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("PUT")
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("DELETE")
}

func (r *Router) Head(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("HEAD")
}

func (r *Router) Connect(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("CONNECT")
}

func (r *Router) Options(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("OPTIONS")
}

func (r *Router) Patch(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("PATCH")
}

func (r *Router) Trace(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods("TRACE")
}

func (r *Router) Handle(method string, path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods(method)
}

func (r *Router) NotFound(handler http.HandlerFunc) {
	r.mux.NotFoundHandler = handler
}
