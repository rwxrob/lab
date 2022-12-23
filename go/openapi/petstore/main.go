package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// router.Use(auth.Middleware)
	router.GET("/api/v1", Index)
	return router
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}

func main() {
	log.Printf("Server started")
	router := NewRouter()
	log.Fatal(router.Run(":8080"))
}
