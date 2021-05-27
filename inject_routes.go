//+build wireinject

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raychongtk/go-web/datastore"
	"github.com/raychongtk/go-web/repository"
	"github.com/raychongtk/go-web/route"
	"github.com/raychongtk/go-web/service"
)
import "github.com/google/wire"

func injectRoutes(ctx context.Context) (*gin.Engine, error) {
	panic(wire.Build(
		datastore.WireSet,
		repository.WireSet,
		service.WireSet,
		route.WireSet,
	))
}
