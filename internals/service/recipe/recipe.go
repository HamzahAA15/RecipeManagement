package recipe

import (
	"context"
	"fmt"
	"recipeApps/internals/models"
	"recipeApps/internals/repository"
	"recipeApps/utils/notification"
	"recipeApps/utils/request"
	"recipeApps/utils/response"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type recipeService struct {
	recipeRepository    repository.IRecipeRepository
	userRepository      repository.IUserRepository
	notificationService notification.INotificationService
}

type RecipeServiceConfig struct {
	RecipeRepository    repository.IRecipeRepository
	UserRepository      repository.IUserRepository
	NotificationService notification.INotificationService
}

func NewRecipeService(cfg RecipeServiceConfig) IRecipeService {
	return &recipeService{
		recipeRepository:    cfg.RecipeRepository,
		userRepository:      cfg.UserRepository,
		notificationService: cfg.NotificationService,
	}
}

func (s *recipeService) CreateRecipe(ctx context.Context, request request.RecipeRequest) (response.RecipeResponseData, error) {
	user := ctx.Value("username").(string)
	upperTitle := strings.ToUpper(request.Title)
	recipe := models.Recipe{
		Title:       upperTitle,
		Category:    strings.ToUpper(request.Category),
		Author:      user,
		Ingredients: request.Ingredients,
		Procedures:  request.Procedures,
		ImageURL:    request.ImageURL,
		Comments:    []string{},
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	//check if title exist
	recipeExist, err := s.recipeRepository.GetRecipeByTitle(upperTitle)
	if err != nil {
		return response.RecipeResponseData{}, err
	}

	if recipeExist.Title == upperTitle {
		return response.RecipeResponseData{}, fmt.Errorf("recipe already exist, create a unique title")
	}

	err = s.recipeRepository.CreateRecipe(recipe)
	if err != nil {
		return response.RecipeResponseData{}, err
	}

	return response.RecipeResponseData{
		Title:  recipe.Title,
		Author: recipe.Author,
	}, nil
}

func (s *recipeService) UpdateRecipeByID(ctx context.Context, request request.RecipeRequest) (response.RecipeResponseData, error) {
	user := ctx.Value("username").(string)
	recipe := models.Recipe{
		ID:          request.ID,
		Title:       strings.ToUpper(request.Title),
		Category:    strings.ToUpper(request.Category),
		Ingredients: request.Ingredients,
		Procedures:  request.Procedures,
		ImageURL:    request.ImageURL,
		UpdatedAt:   time.Now().UTC(),
	}

	if user != "" {
		recipe.Author = user
	}

	err := s.recipeRepository.UpdateRecipeByID(recipe)

	if err != nil {
		return response.RecipeResponseData{}, err
	}

	return response.RecipeResponseData{
		Title:  request.Title,
		Author: user,
	}, nil
}

func (s *recipeService) GetRecipeByID(id string) (response.RecipeResponseData, error) {
	recipe, err := s.recipeRepository.GetRecipeByID(id)
	if err != nil {
		return response.RecipeResponseData{}, err
	}

	return response.RecipeResponseData{
		ID:          recipe.ID,
		Title:       recipe.Title,
		Ingredients: recipe.Ingredients,
		Procedures:  recipe.Procedures,
		Description: recipe.Description,
		Category:    recipe.Category,
		Author:      recipe.Author,
		ImageURL:    recipe.ImageURL,
		Comments:    recipe.Comments,
		Rating:      float32(recipe.Rating),
	}, nil
}

func (s *recipeService) GetRecipesByFilter(author, title, category string) ([]response.RecipeResponseData, error) {
	queryFilter := bson.M{}
	if author != "" {
		queryFilter["author"] = author
	}

	if title != "" {
		queryFilter["title"] = strings.ToUpper(title)
	}

	if category != "" {
		queryFilter["category"] = strings.ToUpper(category)
	}

	recipes, err := s.recipeRepository.GetRecipesByFilter(queryFilter)
	if err != nil {
		return []response.RecipeResponseData{}, err
	}

	if recipes == nil {
		return []response.RecipeResponseData{}, fmt.Errorf("recipe not found")
	}

	var responses []response.RecipeResponseData
	for _, recipe := range recipes {
		responses = append(responses, response.RecipeResponseData{
			ID:          recipe.ID,
			Title:       recipe.Title,
			Ingredients: recipe.Ingredients,
			Procedures:  recipe.Procedures,
			Description: recipe.Description,
			Category:    recipe.Category,
			Author:      recipe.Author,
			ImageURL:    recipe.ImageURL,
			Comments:    recipe.Comments,
			Rating:      recipe.Rating,
		})
	}

	return responses, nil
}

func (s *recipeService) GetAllRecipe() ([]response.RecipeResponseData, error) {
	var responses []response.RecipeResponseData
	recipes, err := s.recipeRepository.GetAllRecipe()
	if err != nil {
		return []response.RecipeResponseData{}, err
	}

	for _, recipe := range recipes {
		responses = append(responses, response.RecipeResponseData{
			ID:          recipe.ID,
			Title:       recipe.Title,
			Ingredients: recipe.Ingredients,
			Procedures:  recipe.Procedures,
			Description: recipe.Description,
			Category:    recipe.Category,
			Author:      recipe.Author,
			ImageURL:    recipe.ImageURL,
			Comments:    recipe.Comments,
			Rating:      recipe.Rating,
		})

	}
	return responses, nil
}

func (s *recipeService) DeleteRecipeByID(ctx context.Context, id string) error {
	recipe, err := s.recipeRepository.GetRecipeByID(id)
	if err != nil {
		return err
	}

	user := ctx.Value("username").(string)
	if user != recipe.Author {
		return fmt.Errorf("you are not allowed to delete this recipe")
	}

	err = s.recipeRepository.DeleteRecipeByID(id)
	if err != nil {
		return err
	}

	return nil
}
