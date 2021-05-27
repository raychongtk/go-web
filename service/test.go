package service

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) Test(ctx *gin.Context) {
	ctx.JSON(200, &testResponse{Result: "verified"})
}

type testResponse struct {
	Result string `json:"result" binding:"required"`
}
