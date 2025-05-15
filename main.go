package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type HostnameResponse struct {
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpPort := os.Getenv("HTTP_PORT")
	version := os.Getenv("VERSION")

	fmt.Println("HTTP_PORT: ", httpPort)
	fmt.Println("VERSION: ", version)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		hostname, err := os.Hostname()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get hostname"})
		}
		return c.JSON(http.StatusOK, HostnameResponse{Hostname: hostname, Version: version})
	})

	e.Logger.Fatal(e.Start(":" + httpPort))
}
