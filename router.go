package main

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "test",
		})
	})
	//r.GET("/look", lookFile)
	//r.GET("/status/get/prefop", viewFile)
}
