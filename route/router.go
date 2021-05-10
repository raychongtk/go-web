package route

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/raychongtk/go-web/service"
)

var (
	WireSet = wire.NewSet(ProvideRoutes)
)

func ProvideRoutes(service *service.Service) *gin.Engine {
	r := gin.New()

	r.PUT("/user/login", service.Login)
	r.POST("/user", service.Register)

	return r
}
