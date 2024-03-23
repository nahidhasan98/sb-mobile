package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/sb-mobile/app/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.LoadHTMLGlob("view/*")
	r.Static("/assets", "./assets")

	router.CreateRoute(r)

	log.Println("Server running on port 6002...")
	r.Run(":6002")
}
