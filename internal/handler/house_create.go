package handler

import (
	"backend-bootcamp-assignment-2024/internal/storage"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HouseCreate(c *gin.Context) {
	var req HouseCreateRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Errorf("handler.HouseCreate,BindJSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if req.Address == "" || req.Year < 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	house, err := h.storage.CreateHouse(c, storage.House{
		Address: req.Address,
		Year:    req.Year,
		Developer: sql.NullString{
			String: req.Developer,
			Valid:  req.Developer != "",
		},
	})
	if err != nil {
		h.logger.Errorf("handler.HouseCreate,storage.CreateHouse error: %v", err)
		h.handleError(c, errors.New("error creating house"))
		return
	}

	c.JSON(http.StatusOK, HouseCreateResponse{
		ID:        house.ID,
		Address:   house.Address,
		Year:      house.Year,
		Developer: house.Developer.String,
		CreatedAt: house.CreatedAt,
		UpdatedAt: house.UpdatedAt,
	})
}
