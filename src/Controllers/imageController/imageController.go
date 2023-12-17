package TC

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

// SetClient sets the MongoDB client for the package
func SetClient(c *mongo.Client) {
	client = c
}

func HandleGet(c *fiber.Ctx) error {
	imgPath := fmt.Sprintf("./images/%s.png", c.Params("imgName"))

	return c.SendFile(imgPath)
}

func HandleGetAll(c *fiber.Ctx) error {
	// Retrieve all documents from the "qrimages" collection
	cursor, err := client.Database("qrimages").Collection("qrimages").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer cursor.Close(context.Background())

	// Decode the cursor into a slice to send as a JSON response
	var images []bson.M
	if err := cursor.All(context.Background(), &images); err != nil {
		log.Fatal(err)
		return err
	}

	return c.Status(200).JSON(images)
}

func HandlePost(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	uuid, _ := uuid.NewUUID()
	imgName := fmt.Sprintf("./images/%s%d.png", uuid.String(), file.Size)
	saveErr := c.SaveFile(file, imgName)
	if saveErr != nil {
		return saveErr
	}

	payload := struct {
		ImgPath    string `json:"imgPath"`
		SchoolName string `json:"schoolName"`
	}{}

	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return err
	}

	dbImageName := fmt.Sprintf("%s%d", uuid.String(), file.Size)
	payload.SchoolName = "test"
	payload.ImgPath = dbImageName
	writeToDb(&payload)

	return c.Status(200).JSON(payload)
}

func writeToDb(payload *struct {
	ImgPath    string `json:"imgPath"`
	SchoolName string `json:"schoolName"`
},
) {
	qrImagesCol := client.Database("qrimages").Collection("qrimages")

	// Correct way to insert a document
	_, insertErr := qrImagesCol.InsertOne(context.TODO(), map[string]interface{}{"imgPath": payload.ImgPath, "schoolName": payload.SchoolName})
	if insertErr != nil {
		log.Fatal(insertErr)
	}

}
