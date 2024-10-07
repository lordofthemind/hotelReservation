package db

import "github.com/lordofthemind/hotelReservation/types"

type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStore struct{}
