package httpservice

import (
	"fmt"
	"recipeApps/internals/service/recipe"
	"recipeApps/utils/request"

	"github.com/gofiber/fiber/v2"
)

type recipeHandler struct {
	recipeService recipe.IRecipeService
}

func NewRecipeHandler(service recipe.IRecipeService) *recipeHandler {
	return &recipeHandler{
		recipeService: service,
	}
}

func (r *recipeHandler) CreateRecipe(c *fiber.Ctx) error {
	req := request.RecipeRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = createValidator(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resp, err := r.recipeService.CreateRecipe(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    resp,
	})
}

func (r *recipeHandler) GetRecipeByID(c *fiber.Ctx) error {
	queryParam := c.Query("id")

	if queryParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Recipe ID cannot be empty",
		})
	}

	resp, err := r.recipeService.GetRecipeByID(queryParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    resp,
	})
}

func (r *recipeHandler) GetRecipesByFilter(c *fiber.Ctx) error {
	queryAuthor := c.Query("author")
	queryTitle := c.Query("title")
	queryCategory := c.Query("category")

	if queryAuthor == "" && queryTitle == "" && queryCategory == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "one of the Author, Title, or Category must be provided",
		})
	}

	resp, err := r.recipeService.GetRecipesByFilter(queryAuthor, queryTitle, queryCategory)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    resp,
	})
}

func (r *recipeHandler) GetAllRecipes(c *fiber.Ctx) error {
	resp, err := r.recipeService.GetAllRecipe()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    resp,
	})
}

func (r *recipeHandler) UpdateRecipe(c *fiber.Ctx) error {
	req := request.RecipeRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if req.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Recipe ID cannot be empty",
		})
	}

	resp, err := r.recipeService.UpdateRecipeByID(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    resp,
	})
}

func (r *recipeHandler) DeleteRecipe(c *fiber.Ctx) error {
	queryParam := c.Query("id")

	if queryParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Recipe ID cannot be empty",
		})
	}

	err := r.recipeService.DeleteRecipeByID(c.Context(), queryParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func createValidator(request request.RecipeRequest) error {
	if request.Title == "" {
		return fmt.Errorf("recipe title cannot be empty")
	} else if request.Category == "" {
		return fmt.Errorf("recipe category cannot be empty")
	} else if request.Ingredients == nil {
		return fmt.Errorf("recipe ingredients cannot be empty")
	} else if request.Procedures == nil {
		return fmt.Errorf("recipe procedures cannot be empty")
	}

	return nil
}
