package route

import (
	"github.com/gin-gonic/gin"
	"github.com/raychongtk/go-web/service"
)

func ProvideRoutes(service *service.Service) *gin.Engine {
	r := gin.New()

	r.PUT("/login", service.Login)

	return r
}
