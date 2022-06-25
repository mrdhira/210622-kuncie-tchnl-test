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

type CheckoutControllers struct {
	checkoutServices services.ICheckoutServices
}

// InitCheckoutControllers func
func InitCheckoutControllers(checkoutServices services.ICheckoutServices) *CheckoutControllers {
	return &CheckoutControllers{
		checkoutServices: checkoutServices,
	}
}

// Checkouts func
func (c *CheckoutControllers) Checkouts(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Request.Header[http.CanonicalHeaderKey("X-User-ID")][0])
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("failed to parse userID to int", err.Error(), nil))
		return
	}

	var data dto.Checkouts
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, helpers.ResponseHTTP("payload error", err.Error(), nil))
		return
	}
	data.UserID = uint(userID)

	orders, err := c.checkoutServices.Checkouts(ctx, data)
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, helpers.ResponseHTTP("something wrong", err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseHTTP("success", nil, orders))
}
