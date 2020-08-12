package simple

import (
	"net"
	"net/http"
)

type Router interface {
	http.Handler

	Use(middleware ...HandlerFunc) Router
	Group(pattern string, middleware ...HandlerFunc) Router

	Handle(pattern string, h Handler)
	HandleFunc(pattern string, hf HandlerFunc)

	Get(pattern string, h Getter)
	GetFunc(pattern string, hf HandlerFunc)

	Post(pattern string, h Poster)
	PostFunc(pattern string, hf HandlerFunc)

	Put(pattern string, h Putter)
	PutFunc(pattern string, hf HandlerFunc)

	Delete(pattern string, h Deleter)
	DeleteFunc(pattern string, hf HandlerFunc)

	Basic(pattern string, h BasicHandler)
	Rest(pattern string, h RestHandler)

	StaticFile(pattern string, file string)
	StaticDir(pattern string, dir string)
	StaticFS(pattern string, fs http.FileSystem)

	Run(addr string) error
	RunTLS(addr, certFile, keyFile string) error

	Serve(lis net.Listener) error
	ServeTLS(lis net.Listener, certFile, keyFile string) error
}
