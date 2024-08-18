package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"backend-bootcamp-assignment-2024/internal/storage"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HouseGet(c *gin.Context) {
	houseID, err := strconv.Atoi(c.Param("id"))
	if err != nil || houseID < 1 {
		h.logger.Errorf("handler.HouseGet,invalid id: %s", c.Param("id"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	flats, err := h.getFlatsByRole(c, int64(houseID))
	if err != nil {
		h.logger.Errorf("handler.HouseGet,storage.GetFlatsByHouseID error: %v", err)
		h.handleError(c, errors.New("error getting flats"))
		return
	}

	c.JSON(http.StatusOK, HouseGetResponse{
		Flats: flats,
	})
}

func (h *Handler) getFlatsByRole(c *gin.Context, houseID int64) ([]Flat, error) {
	var (
		flats []storage.Flat
		err   error
	)

	token := utils.GetTokenFromRequest(c)
	if h.authService.ValidateModeratorRoleJWT(token) == nil {
		flats, err = h.storage.GetFlatsByHouseID(c, houseID, "")
		if err != nil {
			return nil, err
		}
	} else {
		flats, err = h.storage.GetFlatsByHouseID(c, houseID, ApprovedStatus)
		if err != nil {
			return nil, err
		}
	}

	result := make([]Flat, 0, len(flats))

	for _, flat := range flats {
		result = append(result, Flat{
			ID:      flat.ID,
			HouseID: flat.HouseID,
			Price:   flat.Price,
			Rooms:   flat.RoomsCount,
			Status:  flat.Status,
		})
	}

	return result, nil
}
