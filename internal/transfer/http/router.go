package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirzaahmedov/golang-url-shortner/internal/storage"
)

type HTTPRouter struct {
	storage storage.Storage
	router  *gin.Engine
	logger  *slog.Logger
}

func NewRouter(storage storage.Storage, logger *slog.Logger) *HTTPRouter {
	return &HTTPRouter{
		storage: storage,
		router:  gin.Default(),
		logger:  logger,
	}
}

func (r *HTTPRouter) Run(addr string) error {
	return r.router.Run(addr)
}

func (r *HTTPRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
