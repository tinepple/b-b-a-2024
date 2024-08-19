package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"backend-bootcamp-assignment-2024/internal/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Errorf("handler.Login,BindJSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if !utils.IsValidUUID(req.ID) || req.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := h.storage.GetUserByID(c, req.ID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		h.logger.Errorf("handler.Login,storage.GetUserByID error: %v", err)
		h.handleError(c, errors.New("error getting user"))
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.authService.GenerateJWT(user.Role, user.ID)
	if err != nil {
		h.logger.Errorf("handler.Login,authService.GenerateJWT error: %v", err)
		h.handleError(c, errors.New("error generating JWT"))
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
	})
}
