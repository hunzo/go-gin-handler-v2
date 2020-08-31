package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hunzo/go-gin-handler-v2/controllers"
)

func Routers() *gin.Engine {
	h := controllers.RequestHandlers{}
	h.Initial()

	r := gin.Default()
	r.GET("/", h.GetAll)
	r.POST("/", h.PostALL)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	return r

}
