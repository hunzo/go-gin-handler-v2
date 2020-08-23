package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hunzo/go-gin-handler-v2/controllers"
)

func Routers() *gin.Engine {
	h := controllers.RequestHandlers{}

	r := gin.Default()
	r.GET("/", h.GetAll)
	r.POST("/", h.PostALL)

	return r

}
