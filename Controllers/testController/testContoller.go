package TC

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
	saveErr := c.SaveFile(file, "./images/test.png")
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
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	return c.SendFile("./images/test.png")
	/* return c.JSON(payload) */
}
