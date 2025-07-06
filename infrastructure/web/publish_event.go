package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"para.evrard.online/bcs/projects/services"
	"para.evrard.online/infrastructure/commandbus"
)

type ProjectEvent struct {
	ProjectName string
	EventType   string
	EventData   string
}

func PublishEvent(c *gin.Context) {
	var event ProjectEvent

	err := c.ShouldBind(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := commandbus.NewContext(context.Background())
	_, err = commandbus.Dispatch(ctx, &services.ProjectPublishEventAction{SearchString: event.ProjectName, EventType: event.EventType, Data: event.EventData})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Event published"})
		return
	}
}
