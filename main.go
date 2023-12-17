package main

import (
	"context"
	"log"

	TC "QRSweet.com/src/Controllers/imageController"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	client, connectErr := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if connectErr != nil {
		return
	}

	TC.SetClient(client)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173", // Specify the origin of your client application
	}))

	app.Get("/image/:imgName", TC.HandleGet)
	app.Post("/image", TC.HandlePost)
	app.Get("/images", TC.HandleGetAll)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Fatal(app.Listen(":3000"))
}
