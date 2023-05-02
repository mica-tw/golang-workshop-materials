package main

import (
	"fmt"
	"net/http"

	"gin-exercise/pkg/products"

	"github.com/gin-gonic/gin"
)

func putProduct(ctx *gin.Context, store products.ProductStoreIface) {
	prd := &products.ProductRequest{}
	err := ctx.BindJSON(prd)
	if err != nil {
		fmt.Printf("invalid json: %v\n", err)

		return
	}

	storedPrd, err := store.Store(*prd)
	if err != nil {
		ctx.String(
			http.StatusInternalServerError,
			fmt.Sprintf("could not store product: %v", err),
		)

		return
	}

	ctx.JSON(http.StatusOK, &storedPrd)
}

func getPrd(ctx *gin.Context, store products.ProductStoreIface) {
	id := ctx.Param("id")
	prd, err := store.Find(id)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": fmt.Sprintf("could not find product '%s': %v", id, err)},
		)

		return
	}

	if prd == nil {
		ctx.JSON(
			http.StatusNotFound,
			map[string]string{"error": fmt.Sprintf("could not find product '%s'", id)},
		)

		return
	}

	ctx.JSON(http.StatusOK, prd)
}

func main() {
	store := products.NewMemProductStore()

	router := gin.Default()

	// closure in the ad store to conform to interface
	putPrdHandler := func(c *gin.Context) {
		putProduct(c, store)
	}
	router.PUT("/products", putPrdHandler)

	getPrdHandler := func(c *gin.Context) {
		getPrd(c, store)
	}
	router.GET("/products/:id", getPrdHandler)

	router.Run("localhost:8080")
}
