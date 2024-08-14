package handler

import (
	"backend-bootcamp-assignment-2024/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FlatUpdate(c *gin.Context) {
	var req FlatUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	flat, err := h.iStorage.UpdateFlat(c, storage.Flat{
		ID:         req.ID,
		Price:      req.Price,
		RoomsCount: req.Rooms,
		Status:     req.Status,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, FlatUpdateResponse{
		ID:      flat.ID,
		HouseID: flat.HouseID,
		Price:   flat.Price,
		Rooms:   flat.RoomsCount,
		Status:  flat.Status,
	})
}
