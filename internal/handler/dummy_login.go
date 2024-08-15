package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DummyLogin(c *gin.Context) {
	userType := utils.GetQueryString(c, "user_type")

	token, err := h.authService.GenerateJWT(userType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, DummyLoginResponse{
		Token: token,
	})
}
