package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/sb-mobile/app/api"
)

type schedulePayload struct {
	CounterId   string `json:"counterId"`
	StationId   string `json:"stationId"`
	JourneyDate string `json:"journeyDate"`
}

func GetStationsByCounter(c *gin.Context) {
	counterId := c.Param("id")
	_, err := strconv.Atoi(counterId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid counter id",
		})
		return
	}

	stations, err := api.GetStationsByCounter(counterId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "something went wrong, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   stations,
	})
}

func GetSchedule(c *gin.Context) {
	var postData schedulePayload
	err := c.BindJSON(&postData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid JSON object",
		})
		return
	}
	userId := "10000603100"

	schedule, err := api.GetSchedule(postData.CounterId, postData.StationId, postData.JourneyDate, userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "something went wrong, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   schedule,
	})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Mobile | SB Super Deluxe",
	})
}
