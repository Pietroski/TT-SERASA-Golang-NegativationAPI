package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Check iCheck = &sCheck{}
)

type iCheck interface {
	Ping(c *gin.Context)
}

type sCheck struct {}

// Ping godoc
// @Summary Health check
// @Description checks server health
// @ID ping
// @Produce  json
// @Success 200 {object} SuccessfulPing
// @Failure default {object} ErrorStruct
// @host localhost:8009
// @BasePath /v1-legacy
// @Router /ping [get]
// @host localhost:8008
// @BasePath /v2
// @Router /ping [get]
func (ck *sCheck) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"full-path": c.FullPath(),
	})
}

type ErrorStruct struct {
	Error string `json:"error"`
}

type SuccessfulPing struct {
	Message  string      `json:"message"`
	FullPath string `json:"full-path"`
}