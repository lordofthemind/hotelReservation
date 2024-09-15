package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lordofthemind/hotelReservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At water cooler",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("james")
}
