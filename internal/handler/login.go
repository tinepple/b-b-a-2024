package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.iStorage.GetUserByID(c, req.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.authService.GenerateJWT(user.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, LoginResponse{
		Token: token,
	})
}
