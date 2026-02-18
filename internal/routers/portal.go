package routers

import (
	"portal/internal/handlers"

	"github.com/gin-gonic/gin"
)

func portalRouters(r *gin.Engine) {
	r.GET("portal", handlers.CheckPortal)
	productRoutes(r)
}
