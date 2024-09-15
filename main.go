package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/lordofthemind/hotelReservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9090", "The listen port of the API server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
