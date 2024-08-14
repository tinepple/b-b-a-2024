package handler

import (
	"backend-bootcamp-assignment-2024/internal/storage"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HouseCreate(c *gin.Context) {
	var req HouseCreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	house, err := h.iStorage.CreateHouse(c, storage.House{
		Address: req.Address,
		Year:    req.Year,
		Developer: sql.NullString{
			String: req.Developer,
			Valid:  req.Developer != "",
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, HouseCreateResponse{
		ID:        house.ID,
		Address:   house.Address,
		Year:      house.Year,
		Developer: house.Developer.String,
		CreatedAt: house.CreatedAt,
		UpdatedAt: house.UpdatedAt,
	})
}
