package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!\n")
}

func panicHandler(c *gin.Context) {
	panic("oh no")

	c.String(http.StatusOK, "everything is fine 🔥")
}

func main() {
	// gin's default router will recover from all panics automatically
	// router := gin.Default()

	// or we can write middleware ourselves
	router := gin.New()
	router.Use(gin.Logger()) // turn on default logger
	router.Use(func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from: %v\n", r)
				debug.PrintStack()

				c.String(http.StatusInternalServerError, "internal server error")
				c.Abort()
			}
		}()

		c.Next()
	})

	router.GET("/hello", helloHandler)

	router.GET("/panic", panicHandler)

	router.Run("localhost:8080")
}
