package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"backend-bootcamp-assignment-2024/internal/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Errorf("handler.Register,BindJSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, ok := ValidUserTypes[req.UserType]
	if req.Email == "" || req.Password == "" || !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		h.logger.Errorf("handler.Register,utils.HashPassword error: %v", err)
		h.handleError(c, errors.New("error hashing password"))
		return
	}

	userID, err := h.storage.CreateUser(c, storage.User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.UserType,
	})
	if err != nil {
		h.logger.Errorf("handler.Register,storage.CreateUser error: %v", err)
		h.handleError(c, errors.New("error creating user"))
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		UserID: userID,
	})
}
