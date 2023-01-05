package main

import (
	"github.com/davidgodeness/tododemo/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	routes.RegisterWebRoutes(r)
	routes.RegisterApiRoutes(r)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
