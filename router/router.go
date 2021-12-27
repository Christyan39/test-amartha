package router

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params) Response

type RouterHelper struct {
	router *httprouter.Router
}

type RouterIO interface {
	Serve(port string)
	GET(path string, handle Handle)
	POST(path string, handle Handle)
}

type HttpRouterIO interface {
	GET(path string, handle httprouter.Handle)
	POST(path string, handle httprouter.Handle)
}

func getHttpRouterIO(router *httprouter.Router) HttpRouterIO {
	return router
}

func NewRouter() RouterIO {
	return &RouterHelper{
		router: httprouter.New(),
	}
}

func (r *RouterHelper) Serve(port string) {
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r.router))
}

func (r *RouterHelper) GET(path string, handle Handle) {
	log.Println(path)
	getHttpRouterIO(r.router).GET(path, RespHandler(handle))
}

func (r *RouterHelper) POST(path string, handle Handle) {
	log.Println(path)
	getHttpRouterIO(r.router).POST(path, RespHandler(handle))
}
