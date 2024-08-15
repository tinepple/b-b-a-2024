package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"backend-bootcamp-assignment-2024/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	userID, err := h.iStorage.CreateUser(c, storage.User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.UserType,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, RegisterResponse{
		UserID: userID,
	})
}
