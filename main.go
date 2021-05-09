package main

import (
	"github.com/raychongtk/go-web/repository"
	"github.com/raychongtk/go-web/route"
	"github.com/raychongtk/go-web/service"
	"net/http"
)

func main() {
	mockRepository := repository.NewRepository()
	userService, err := service.ProvideService(mockRepository)
	if err != nil {
		return
	}
	http.ListenAndServe(":8080", route.ProvideRoutes(userService))
}
