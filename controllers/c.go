package controllers

import (
	"fmt"
	"net/http"

	"github.com/hunzo/go-gin-handler-v2/models"

	"github.com/gin-gonic/gin"
)

type RequestHandlers struct{}

func (h *RequestHandlers) Initial() {
	fmt.Println("Initial Functions")
}

func (h *RequestHandlers) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"info": "GetALL",
	})
}

func (h *RequestHandlers) PostALL(c *gin.Context) {
	u := models.LoginUser{}
	h.Initial()

	if e := c.ShouldBindJSON(&u); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"u status": u,
			"status":   "errors",
			"info":     e.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"info":   u,
	})
}
