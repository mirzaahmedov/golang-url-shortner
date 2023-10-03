package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirzaahmedov/golang-url-shortner/internal/storage"
)

type HTTPRouter struct {
	storage storage.Storage
	router  *gin.Engine
}

func NewRouter(storage storage.Storage) *HTTPRouter {
	return &HTTPRouter{
		storage: storage,
		router:  gin.Default(),
	}
}

func (r *HTTPRouter) Run(addr string) error {
	return r.router.Run(addr)
}

func (r *HTTPRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
