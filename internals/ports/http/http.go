package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrdhira/210622-kuncie-tchnl-test/configs"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/ports/http/controllers"
)

// InitHttp func
func InitHttp(
	httpConfigs *configs.HttpConfigs,
	cartControllers *controllers.CartControllers,
	checkoutControllers *controllers.CheckoutControllers,
) *gin.Engine {
	router := gin.Default()
	router.Use(commonHeaderValidation)

	cartRouter := router.Group("/carts")
	{
		cartRouter.GET("/", cartControllers.GetCarts)
		cartRouter.POST("/", cartControllers.CreateCarts)
		cartRouter.PUT("/", cartControllers.UpdateCarts)
		cartRouter.DELETE("/", cartControllers.DeleteCarts)
	}

	checkoutRouter := router.Group("/checkouts")
	{
		checkoutRouter.POST("/", checkoutControllers.Checkouts)
	}

	return router
}
