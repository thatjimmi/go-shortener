package server

import (
	"lenk/model"
	"lenk/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	LenkUrl := c.Params("redirect")
	lenk, err := model.FindByShorten(LenkUrl)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	lenk.Clicked++

	err = model.UpdateLenk(lenk)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect(lenk.Redirect)
}

func getAllLenks(c *fiber.Ctx) error {
	lenks, err := model.GetAllLenks()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(lenks)
}

func createLenk(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var lenk model.Lenk
	err := c.BodyParser(&lenk)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if lenk.Random {
		lenk.Shorten = utils.RandomUrlString()
	}

	err = model.CreateLenk(&lenk)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(lenk)
}

func SetupRoutes() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/api/lenks", getAllLenks)
	router.Post("/api/lenks", createLenk)
	router.Get("/:redirect", redirect)

	router.Listen(":3000")
}
