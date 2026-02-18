package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouters() {
	r := gin.Default()
	portalRouters(r)
	r.Run()
}
