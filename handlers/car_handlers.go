package handlers

import (
	"cars/pranay/github.com/models"
	"fmt"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var mu sync.Mutex

func CreateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JSON input",
			"details": err.Error(),
		})
	}
	car.Insert()
	fmt.Printf("Car saved with ID %v", car.Id)
	c.Status(fiber.StatusCreated).JSON(car)
	return nil
}

func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}
	car.Id = int64(id)
	if err := car.Delete(); err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Car not found",
		})
	}
	fmt.Printf("Car deleted successfully with ID %d", car.Id)
	c.Status(fiber.StatusNoContent)
	return nil
}

func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}
	car.Id = int64(id)
	if err := car.Get(); err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Car not found",
			"id":    car.Id,
		})
	}
	c.Status(fiber.StatusOK).JSON(car)
	return nil
}

func UpdateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JSON input",
			"details": err.Error(),
		})
	}
	car.Update()
	c.Status(fiber.StatusOK).JSON(car)
	return nil
}
