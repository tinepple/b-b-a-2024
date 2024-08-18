package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) handleError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, InternalErrorResponse{
		Code:    500,
		Message: err.Error(),
	})
}
