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

func (ck *sCheck) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"full-path": c.FullPath(),
	})
}
