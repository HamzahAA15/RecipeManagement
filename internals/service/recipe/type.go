package recipe

import (
	"context"
	"recipeApps/utils/request"
	"recipeApps/utils/response"
)

type IRecipeService interface {
	//recipe
	CreateRecipe(ctx context.Context, request request.RecipeRequest) (response.RecipeResponseData, error)
	UpdateRecipeByID(ctx context.Context, request request.RecipeRequest) (response.RecipeResponseData, error)
	GetRecipeByID(id string) (response.RecipeResponseData, error)
	GetAllRecipe() ([]response.RecipeResponseData, error)
	GetRecipesByFilter(author, title, category string) ([]response.RecipeResponseData, error)
	DeleteRecipeByID(ctx context.Context, id string) error
	//comment
	AddComment(request request.CommentRequest) error
	//rating
	AddRating(request request.RatingRequest) (float32, error)
}
