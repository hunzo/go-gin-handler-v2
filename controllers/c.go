package controllers

import (
	"fmt"
	"net/http"

	"github.com/hunzo/go-gin-handler-v2/models"

	"github.com/gin-gonic/gin"
)

type RequestHandlers struct{}

func (h *RequestHandlers) initial() {
	fmt.Println("Initial Functions")
}

func (h *RequestHandlers) GetAll(c *gin.Context) {
	h.initial()
	c.JSON(http.StatusOK, gin.H{
		"info": "GetALL",
	})
}

func (h *RequestHandlers) PostALL(c *gin.Context) {
	u := models.LoginUser{}

	if e := c.ShouldBindJSON(&u); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "erros",
			"info":   e.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
