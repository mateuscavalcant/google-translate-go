package main

import (
	"google-translate-go/pkg/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/translate", func(c *gin.Context) {
		c.HTML(http.StatusOK, "translate.html", gin.H{})
	})

	routes.InitRoutes(r.Group("/"))

	r.Run(":8080")

}
