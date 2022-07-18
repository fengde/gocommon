package main

import (
	"github.com/fengde/gocommon/ginx"
	"github.com/fengde/gocommon/ginx/middlewarex"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewarex.AccessLog())
	router.GET("/health", ginx.Handler(func(c *ginx.Context) {
		c.OutSuccess(map[string]interface{}{
			"health": "ok",
		})
	}))
	router.Run(":8089")
}
