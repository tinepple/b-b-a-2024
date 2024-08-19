package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"backend-bootcamp-assignment-2024/internal/storage"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FlatUpdate(c *gin.Context) {
	var req FlatUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Errorf("handler.FlatUpdate,BindJSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, ok := ValidFlatStatuses[req.Status]
	if req.ID < 1 || req.Price < 0 || req.Rooms < 1 || !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID, err := h.authService.GetUserID(utils.GetTokenFromRequest(c))
	if err != nil {
		h.logger.Errorf("handler.GetUserID error: %v", err)
		h.handleError(c, errors.New("error getting userID from token"))
		return
	}

	if userID == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	isFlatUpdateAvailable, err := h.isFlatUpdateAvailable(c, userID, req.ID, req.Status)
	if err != nil {
		h.handleError(c, errors.New("access checking error"))
		return
	}

	if !isFlatUpdateAvailable {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	flat, err := h.storage.UpdateFlat(c, storage.Flat{
		ID:         req.ID,
		Price:      req.Price,
		RoomsCount: req.Rooms,
		Status:     req.Status,
		ModeratorID: sql.NullString{
			String: userID,
			Valid:  userID != "",
		},
	})
	if err != nil {
		h.logger.Errorf("handler.FlatUpdate,storage.UpdateFlat error: %v", err)
		h.handleError(c, errors.New("error updating flat"))
		return
	}

	c.JSON(http.StatusOK, FlatUpdateResponse{
		ID:      flat.ID,
		HouseID: flat.HouseID,
		Price:   flat.Price,
		Rooms:   flat.RoomsCount,
		Status:  flat.Status,
	})
}

func (h *Handler) isFlatUpdateAvailable(c *gin.Context, userID string, flatID int64, status string) (bool, error) {
	flat, err := h.storage.GetFlatByID(c, flatID)
	if err != nil {
		h.logger.Errorf("handler.FlatUpdate,storage.GetFlatByID error: %v", err)
		return false, err
	}

	if flat.Status == ApprovedStatus || flat.Status == DeclinedStatus || flat.Status == OnModerationStatus && (flat.ModeratorID.Valid && flat.ModeratorID.String != userID) {
		return false, nil
	}

	return true, nil
}
