package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, route string) {
	// Options
	handler.Use(gin.Logger(), gin.Recovery())

	v1 := handler.Group(route)

	v1.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
}
