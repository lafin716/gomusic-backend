package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MyCustomMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("v", 123)
		c.Next()

		//status := c.Writer.Status()
	}
}

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("******************************")
		c.Next()
		fmt.Println("******************************")
	}
}
