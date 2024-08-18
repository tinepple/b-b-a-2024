package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HouseSubscribe(c *gin.Context) {
	houseID, err := strconv.Atoi(c.Param("id"))
	if err != nil || houseID < 1 {
		h.logger.Errorf("handler.HouseGet,invalid id: %s", c.Param("id"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req HouseSubscribeRequest
	if err = c.BindJSON(&req); err != nil {
		h.logger.Errorf("handler.HouseSubscribe,BindJSON error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	user, err := h.storage.GetUserByEmail(c, req.Email)
	if err != nil {
		h.logger.Errorf("handler.HouseSubscribe,storage.GetUserByEmail error: %v", err)
		h.handleError(c, errors.New("error getting user"))
		return
	}

	err = h.storage.CreateHouseUserSubscription(c, int64(houseID), user.ID)
	if err != nil {
		h.logger.Errorf("handler.HouseSubscribe,storage.CreateHouseUserSubscription error: %v", err)
		h.handleError(c, errors.New("error creating house user subscription"))
		return
	}

	c.Status(http.StatusOK)
}
