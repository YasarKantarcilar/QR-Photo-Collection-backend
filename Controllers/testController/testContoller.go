package TC

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleGet(c *fiber.Ctx) error {
	imgPath := fmt.Sprintf("./images/%s.png", c.Params("imgName"))
	return c.SendFile(imgPath)
}

func HandlePost(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	fmt.Println(file.Filename)

	uuid, _ := uuid.NewUUID()
	imgName := fmt.Sprintf("./images/%s%s.png", uuid.String(), file.Filename)
	saveErr := c.SaveFile(file, imgName)
	if saveErr != nil {
		return saveErr
	}
	payload := struct {
		Name           string `json:"name"`
		Price          string `json:"price"`
		ItemType       string `json:"itemType"`
		ItemLength     string `json:"itemLength"`
		AvailableSizes string `json:"availableSizes"`
		Stock          string `json:"stock"`
		ImgPath        string `json:"imgPath"`
	}{}

	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return err
	}

	payload.ImgPath = imgName

	/* return c.SendFile("./images/test.png") */
	fmt.Println(payload)
	return c.JSON(payload)
}
