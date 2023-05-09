package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func postProduct(ctx *gin.Context, store products.ProductStoreIface) {
	prd := &products.ProductRequest{}
	err := ctx.BindJSON(prd)
	if err != nil {
		fmt.Printf("invalid json: %v\n", err)

		return
	}

	id := ctx.Param("id")

	storedPrd, err := store.StoreOrUpdate(id, *prd)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]string{"error": "could not store product"},
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

func getManyPrds(ctx *gin.Context, store products.ProductStoreIface) {
	limit := ctx.Query("limit")
	var limitPtr *int
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				map[string]string{"error": "limit must be a number"},
			)

			return
		}

		limitPtr = &limitInt
	}

	prds, err := store.FindMany(limitPtr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, prds)
}

func deletePrd(ctx *gin.Context, store products.ProductStoreIface) {
	id := ctx.Param("id")
	err := store.Delete(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": "internal service error, could not delete object"})

		return
	}

	ctx.String(http.StatusOK, "ok")
}

func main() {
	store := products.NewMemProductStore()

	router := gin.Default()

	// closure in the ad store to conform to interface
	putPrdHandler := func(c *gin.Context) {
		putProduct(c, store)
	}
	router.PUT("/products", putPrdHandler)
	router.POST("/products", putPrdHandler)

	postPrdHandler := func(c *gin.Context) {
		postProduct(c, store)
	}
	router.POST("/products/:id", postPrdHandler)

	getPrdHandler := func(c *gin.Context) {
		getPrd(c, store)
	}
	router.GET("/products/:id", getPrdHandler)

	getManyPrdsHandler := func(c *gin.Context) {
		getManyPrds(c, store)
	}
	router.GET("/products", getManyPrdsHandler)

	deletePrdHandler := func(c *gin.Context) {
		deletePrd(c, store)
	}
	router.DELETE("products/:id", deletePrdHandler)

	router.Run("localhost:8080")
}
