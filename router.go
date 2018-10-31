package prometheusRouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router struct {
	promHandler http.Handler
}

func NewRouter() *Router {
	return &Router{
		promHandler: promhttp.Handler(),
	}
}

func (r *Router) Register(router gin.IRouter) {
	router.GET("/metrics", func(c *gin.Context) {
		r.promHandler.ServeHTTP(c.Writer, c.Request)
	})
}
