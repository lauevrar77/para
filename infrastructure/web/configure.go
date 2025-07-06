package web

import "github.com/gin-gonic/gin"

func Configure(router *gin.Engine) {
	router.GET("/ping", Ping)
	router.POST("/projects/publish_event", PublishEvent)
}
