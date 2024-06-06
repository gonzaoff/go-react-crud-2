package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gonzaoff/go-react-crud-2/models"
)

func main() {
	fmt.Println("Despegando 🚀🚀🚀")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	

	// --------------- Inicio base de datos -------------------

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/go-mongo"))

	if err != nil{
		panic(err)
	}



	// ---------------- Fin base de datos --------------------
	

	// ------------------- Inicio Peticiones -----------------

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error{
		return c.JSON(&fiber.Map{
			"data": "usuarios desde el back",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {

		var user models.User

		c.BodyParser(&user)

		coll := client.Database("go-mongo").Collection("users")

		coll.InsertOne(
			context.TODO(),
			bson.D{{
				Key: "name", 
				Value: user.Name,
			}})

		return c.JSON(&fiber.Map{
			"data": "Guardando usuario",
		})
	})

	// --------------------- Fin Peticiones ------------------
 

	app.Listen(":" + port)
	fmt.Println("Server on port 3000")
}