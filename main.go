package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/sb-mobile/app/router"
)

//go:embed view/*
var viewFS embed.FS

//go:embed assets/**
var assetsFS embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Load templates from embedded filesystem
	tmpl := template.Must(template.ParseFS(viewFS, "view/*"))
	r.SetHTMLTemplate(tmpl)

	// Serve static files from embedded filesystem
	r.StaticFS("/assets", http.FS(assetsFS))

	router.CreateRoute(r)

	log.Println("Server running on port 6002...")
	r.Run(":6002")
}
