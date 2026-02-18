package routers

import (
	"portal/internal/handlers"

	"github.com/gin-gonic/gin"
)

func productRoutes(r *gin.Engine) {
	r.GET("portal/products", handlers.ListProducts)
	r.GET("portal/product", handlers.GetProductById)
	r.POST("portal/product", handlers.CreateProduct)
	r.PATCH("portal/product", handlers.UpdateProduct)
	r.DELETE("portal/product", handlers.DeleteProduct)
}
