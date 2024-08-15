package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router      *gin.Engine
	iStorage    iStorage
	authService authService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	h.router.GET("/dummyLogin", h.DummyLogin)
	h.router.POST("/login", h.Login)
	h.router.POST("/register", h.Register)

	authGroup := h.router.Group("/", h.Auth)
	authGroup.GET("/house/{id}", h.HouseGet)
	authGroup.POST("/house/{id}/subscribe", h.HouseGet)
	authGroup.POST("/flat/create", h.FlatCreate)

	moderatorGroup := h.router.Group("/", h.ModerationAuth)
	moderatorGroup.POST("/house/create", h.HouseCreate)
	moderatorGroup.POST("/flat/update", h.FlatUpdate)
}

func New(iStorage iStorage, authService authService) *Handler {
	h := &Handler{
		router:      gin.New(),
		iStorage:    iStorage,
		authService: authService,
	}

	h.initRoutes()

	return h
}
