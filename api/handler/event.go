package handler

import (
	"fmt"
	"go-sse/internal"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
}

func NewEventHandler() *Event {
	return &Event{}
}

func (e *Event) GetEvents(c *gin.Context) {
	clientID := c.Query("id")
	if clientID == "" {
		clientID = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	client := internal.NewClient(clientID, make(chan string, 10))
	internal.AddClient(client)
	defer internal.RemoveClient(clientID)
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	for {
		select {
		case message := <-client.Chan:
			fmt.Fprintf(c.Writer, "data: %s\n\n", message)
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

func (e *Event) PostEvents(c *gin.Context) {
	var body struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	internal.Broadcast(body.Message)
	c.JSON(http.StatusOK, gin.H{"status": "Message broadcasted"})
}
