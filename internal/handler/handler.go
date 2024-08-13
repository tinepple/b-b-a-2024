package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router   *gin.Engine
	iStorage iStorage
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	h.router.GET("/dummyLogin", h.DummyLogin)
	h.router.POST("/login", h.Login)
	h.router.POST("/register", h.Register)

	authGroup := h.router.Group("/", h.Auth)

	h.router.POST("/house/create", h.HouseCreate) //moderator
	authGroup.GET("/house/{id}", h.HouseGet)
	authGroup.POST("/house/{id}/subscribe", h.HouseGet)
	authGroup.POST("/flat/create", h.FlatCreate)
	h.router.POST("/flat/update", h.FlatUpdate) //moderator
}

func New(iStorage iStorage) *Handler {
	h := &Handler{
		router:   gin.New(),
		iStorage: iStorage,
	}

	h.initRoutes()

	return h
}
