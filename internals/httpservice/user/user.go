package httpservice

import (
	"fmt"
	"recipeApps/internals/service/user"
	"recipeApps/utils/request"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService user.IUserService
}

func NewUserHandler(service user.IUserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	req := request.CreateUser{}
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

	err = h.userService.CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *userHandler) GetUserByName(c *fiber.Ctx) error {
	queryParam := c.Query("name")
	if queryParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User name cannot be empty",
		})
	}

	resp, err := h.userService.GetUserByName(queryParam)
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

func createValidator(req request.CreateUser) error {
	if req.Name == "" {
		return fmt.Errorf("user name cannot be empty")
	} else if req.Email == "" {
		return fmt.Errorf("user email cannot be empty")
	} else if req.Password == "" {
		return fmt.Errorf("user password cannot be empty")
	}
	return nil
}
