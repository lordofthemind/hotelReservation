package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lordofthemind/hotelReservation/db"
	"github.com/lordofthemind/hotelReservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At water cooler",
	}
	return c.JSON(u)
}
