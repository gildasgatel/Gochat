package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Home(c *gin.Context)
}

type Handle struct {
}

func New() Handler {
	return &Handle{}
}

func (h *Handle) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
