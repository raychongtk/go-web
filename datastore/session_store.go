package datastore

import (
	"github.com/gin-contrib/sessions/redis"
	"github.com/raychongtk/go-web/util"
)

func ProvideSessionStore() redis.Store {
	var config *SessionConfig
	util.Load("dev", ".", &config)
	store, _ := redis.NewStore(10, "tcp", config.SessionAddress, config.SessionPassword, []byte(config.SessionKey))
	return store
}

type SessionConfig struct {
	SessionAddress  string `mapstructure:"SESSION_ADDRESS"`
	SessionPassword string `mapstructure:"SESSION_PASSWORD"`
	SessionKey      string `mapstructure:"SESSION_KEY"`
}
