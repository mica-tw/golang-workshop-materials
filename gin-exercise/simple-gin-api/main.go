package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World!")
}

func greetHandler(ctx *gin.Context) {
	greeting := ctx.DefaultQuery("greeting", "Hello")

	name := ctx.Param("name")

	ctx.String(http.StatusOK, fmt.Sprintf("%s %s!", greeting, name))
}

func panicHandler(ctx *gin.Context) {
	panic("oh no")

	ctx.String(http.StatusInternalServerError, "panicked")
}

type postReq struct {
	FieldOne string
	FieldTwo string `json:"FieldTwo" binding:"required"`
}

func postHandler(ctx *gin.Context) {
	req := &postReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "malformed json"})
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("Okay %v", req))
}

func main() {
	router := gin.Default()

	router.GET("/hello", helloHandler)
	router.GET("/hello/:name", greetHandler)

	router.GET("/panic", panicHandler)

	router.POST("/post", postHandler)

	router.Run("localhost:8080")
}
