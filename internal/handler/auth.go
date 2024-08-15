package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Auth(c *gin.Context) {
	token := utils.GetTokenFromRequest(c)
	err := h.authService.ValidateClientRoleJWT(token)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (h *Handler) ModerationAuth(c *gin.Context) {
	token := utils.GetTokenFromRequest(c)
	err := h.authService.ValidateModeratorRoleJWT(token)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
