package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/sb-mobile/app/controller"
)

func CreateRoute(router *gin.Engine) {
	router.GET("/", controller.Index)

	apiRG := router.Group("/api")
	apiRG.GET("/getStations/:id", controller.GetStationsByCounter)
	apiRG.POST("/getSchedule", controller.GetSchedule)
}
