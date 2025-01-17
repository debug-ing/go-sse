package router

import (
	"go-sse/api/handler"

	"github.com/gin-gonic/gin"
)

func InitEventRouter(r *gin.RouterGroup) {
	eventHandler := handler.NewEventHandler()
	r.GET("/events", eventHandler.GetEvents)
}
