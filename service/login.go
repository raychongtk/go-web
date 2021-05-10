package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Service) Login(ctx *gin.Context) {
	var req loginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid params")
		return
	}

	authenticated := h.userRepo.Authenticate(req.Username, req.Password)
	if !authenticated {
		ctx.JSON(http.StatusUnauthorized, &loginResponse{Result: false, ErrorCode: "LOGIN_FAILED"})
		return
	}

	ctx.JSON(http.StatusOK, &loginResponse{Result: authenticated})
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Result    bool   `json:"result" binding:"required"`
	ErrorCode string `json:"error_code,omitempty"`
}
