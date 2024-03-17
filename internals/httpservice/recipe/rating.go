package httpservice

import (
	"recipeApps/utils/request"

	"github.com/gofiber/fiber/v2"
)

func (r *recipeHandler) AddRating(c *fiber.Ctx) error {
	req := request.RatingRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if req.Rating <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Rating should be greater than 0",
		})
	}

	if req.RecipeID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Recipe ID is required",
		})
	}

	resp, err := r.recipeService.AddRating(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Success",
		"new_rating": resp,
	})
}
