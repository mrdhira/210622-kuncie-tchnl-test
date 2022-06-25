package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dto"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/services"
	"github.com/mrdhira/210622-kuncie-tchnl-test/pkg/helpers"
)

type CartControllers struct {
	cartServices services.ICartServices
}

// InitCartControllers func
func InitCartControllers(cartServices services.ICartServices) *CartControllers {
	return &CartControllers{
		cartServices: cartServices,
	}
}

// GetCarts func
func (c *CartControllers) GetCarts(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Request.Header[http.CanonicalHeaderKey("X-User-ID")][0])
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("failed to parse userID to int", err.Error(), nil))
		return
	}

	carts, err := c.cartServices.GetCarts(ctx, uint(userID))
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("something wrong", err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseHTTP("success", nil, carts))
}

// CreateCarts func
func (c *CartControllers) CreateCarts(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Request.Header[http.CanonicalHeaderKey("X-User-ID")][0])
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("failed to parse userID to int", err.Error(), nil))
		return
	}

	var data dto.CreateCartProducts
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, helpers.ResponseHTTP("payload error", err.Error(), nil))
		return
	}

	carts, err := c.cartServices.CreateCartProducts(ctx, uint(userID), data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("something wrong", err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseHTTP("success", nil, carts))
}

// UpdateCarts func
func (c *CartControllers) UpdateCarts(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Request.Header[http.CanonicalHeaderKey("X-User-ID")][0])
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("failed to parse userID to int", err.Error(), nil))
		return
	}

	var data dto.UpdateCartProducts
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, helpers.ResponseHTTP("payload error", err.Error(), nil))
		return
	}

	carts, err := c.cartServices.UpdateCartProducts(ctx, uint(userID), data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("something wrong", err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseHTTP("success", nil, carts))
}

// DeleteCarts func
func (c *CartControllers) DeleteCarts(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Request.Header[http.CanonicalHeaderKey("X-User-ID")][0])
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("failed to parse userID to int", err.Error(), nil))
		return
	}

	var data dto.DeleteCartProducts
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, helpers.ResponseHTTP("payload error", err.Error(), nil))
		return
	}

	err = c.cartServices.DeleteCartProducts(ctx, uint(userID), data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("something wrong", err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseHTTP("success", nil, nil))
}
