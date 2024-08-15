package handler

import (
	"backend-bootcamp-assignment-2024/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FlatCreate(c *gin.Context) {
	var req FlatCreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	flat, err := h.iStorage.CreateFlat(c, storage.Flat{
		HouseID:    req.HouseID,
		Price:      req.Price,
		RoomsCount: req.Rooms,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, FlatCreateResponse{
		ID:      flat.ID,
		HouseID: flat.HouseID,
		Price:   flat.Price,
		Rooms:   flat.RoomsCount,
		Status:  flat.Status,
	})
}
