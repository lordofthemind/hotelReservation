package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lordofthemind/hotelReservation/api"
	"github.com/lordofthemind/hotelReservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoUri = "mongodb://localhost:27017/"
const dbName = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	ctx := context.Background()
	coll := client.Database(dbName).Collection(userColl)
	log.Println(client)
	user := types.User{
		FirstName: "James",
		LastName:  "At the water cooler",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	var james types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
		log.Fatal(err)
	}
	fmt.Println(james)
	listenAddr := flag.String("listenAddr", ":9090", "The listen port of the API server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	userHandler :=
		apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
