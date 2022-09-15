package handlers

import (
	"backend/databases"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	var todo []models.Todo

	databases.Database.Find(&todo)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	result := databases.Database.Find(&todo, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&todo)
}

func AddTask(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	databases.Database.Create(&todo)
	return c.Status(201).JSON(todo)
}

func UpdateTask(c *fiber.Ctx) error {
	todo := new(models.Todo)
	id := c.Params("id")

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	databases.Database.Where("id = ?", id).Updates(&todo)
	return c.Status(200).JSON(todo)
}

func RemoveTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	result := databases.Database.Delete(&todo, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
