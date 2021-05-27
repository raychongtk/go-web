// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raychongtk/go-web/datastore"
	"github.com/raychongtk/go-web/repository"
	"github.com/raychongtk/go-web/route"
	"github.com/raychongtk/go-web/service"
)

// Injectors from inject_routes.go:

func injectRoutes(ctx context.Context) (*gin.Engine, error) {
	db := datastore.ProvideDBConnection()
	userRepository := repository.NewRepository(db)
	serviceService, err := service.ProvideService(userRepository)
	if err != nil {
		return nil, err
	}
	store := datastore.ProvideSessionStore()
	engine := route.ProvideRoutes(serviceService, store)
	return engine, nil
}
