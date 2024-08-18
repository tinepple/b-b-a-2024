package handler

import (
	"backend-bootcamp-assignment-2024/internal/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FlatCreate(c *gin.Context) {
	var req FlatCreateRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Errorf("handler.FlatCreate,BindJSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if req.HouseID < 1 || req.Price < 0 || req.Rooms < 1 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	flat, err := h.storage.CreateFlat(c, storage.Flat{
		HouseID:    req.HouseID,
		Price:      req.Price,
		RoomsCount: req.Rooms,
	})
	if err != nil {
		h.logger.Errorf("handler.FlatCreate,storage.CreateFlat error: %v", err)
		h.handleError(c, errors.New("error creating flat"))
		return
	}

	err = h.kafkaService.Produce(req.HouseID)
	if err != nil {
		h.logger.Errorf("handler.FlatCreate,kafkaService.Produce error: %v", err)
		h.handleError(c, errors.New("error producing message"))
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
