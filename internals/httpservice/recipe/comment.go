package httpservice

import (
	"recipeApps/utils/request"

	"github.com/gofiber/fiber/v2"
)

func (r *recipeHandler) AddComment(c *fiber.Ctx) error {
	req := request.CommentRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if req.RecipeID == "" || req.Comment == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Recipe ID and Comment cannot be empty",
		})
	}

	err = r.recipeService.AddComment(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}
