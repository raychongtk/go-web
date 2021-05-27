package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("SessionId", uuid.New().String())
	session.Save()
}

func GetSessionId(ctx *gin.Context) interface{} {
	session := sessions.Default(ctx)
	sessionId := session.Get("SessionId")
	return sessionId
}
