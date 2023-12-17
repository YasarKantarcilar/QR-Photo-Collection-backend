package TC

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleGet(c *fiber.Ctx) error {
	imgPath := fmt.Sprintf("./images/%s.png", c.Params("imgName"))

	return c.SendFile(imgPath)
}

func HandleGetAll(c *fiber.Ctx) error {
	content, err := os.ReadFile("db.txt")
	if err != nil {
		return err
	}

	return c.Status(200).SendString(string(content))
}

func HandlePost(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	uuid, _ := uuid.NewUUID()
	imgName := fmt.Sprintf("./images/%s%s.png", uuid.String(), file.Size)
	saveErr := c.SaveFile(file, imgName)
	if saveErr != nil {
		return saveErr
	}
	payload := struct {
		ImgPath string `json:"imgPath"`
	}{}

	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return err
	}

	payload.ImgPath = imgName
	dbImageName := fmt.Sprintf("%s%s", uuid.String(), file.Size)
	writeToDb(dbImageName)

	return c.Status(200).JSON(payload)
}

func writeToDb(name string) {
	db, err := os.OpenFile("db.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	writeName := fmt.Sprintf("%s\n", name)
	content, readErr := os.ReadFile("db.txt")
	if readErr != nil {
		return
	}
	readContent := []byte(content)
	var writeIndex int64 = int64(len(readContent))

	db.WriteAt([]byte(writeName), writeIndex)
	db.Close()
}
