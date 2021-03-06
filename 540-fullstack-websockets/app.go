package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"log"
)

// definiujemy struckturę
type (
	Geolocation struct {
		Altitude  float64 `json:"alt"`
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	}
)

// i nasze przykładowe dane
var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
)

// handler dla WebSocketów
func wsHandler(c *echo.Context) (err error) {
	var response Geolocation
	ws := c.Socket()

	for {
		// wyplujemy wszystkie lokacje co 100ms
		for _, location := range locations {
			if err = websocket.JSON.Send(ws, location); err != nil {
				log.Println(err)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}

		if err = websocket.JSON.Receive(ws, &response); err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("%+v\n", response)
	}
	return
}

func main() {
	e := echo.New()

	e.Static("/", ".")
	e.WebSocket("/ws", wsHandler)

	e.Run(":8080")
}
