package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DummyLogin(c *gin.Context) {
	userType := utils.GetQueryString(c, "user_type")
	_, ok := ValidUserTypes[userType]
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.authService.GenerateJWT(userType, 0)
	if err != nil {
		h.logger.Errorf("handler.DummyLogin,authService.GenerateJWT error: %v", err)
		h.handleError(c, errors.New("jwt generation error"))
		return
	}

	c.JSON(http.StatusOK, DummyLoginResponse{
		Token: token,
	})
}
