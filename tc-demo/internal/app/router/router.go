package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tc-demo/internal/app/handler"
)

func NewRouter(productHandler handler.ProductHandler) http.Handler {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// Product API endpoints
	v1 := r.Group("/v1")
	{
		v1.GET("/products", productHandler.GetProducts)
		v1.GET("/products/:id", productHandler.GetProduct)
		v1.POST("/products", productHandler.CreateProduct)
		v1.PUT("/products/:id", productHandler.UpdateProduct)
		v1.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	return r
}
