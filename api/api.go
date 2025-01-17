package api

import (
	"go-sse/api/router"

	"github.com/gin-gonic/gin"
)

func InitAPI(r *gin.Engine) {
	g := r.Group("/api")
	router.InitEventRouter(g)
}
