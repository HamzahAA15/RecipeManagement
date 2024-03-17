package recipe

import (
	"fmt"
	"recipeApps/internals/models"
	"recipeApps/utils"
	"recipeApps/utils/notification"
	"recipeApps/utils/request"
)

func (s *recipeService) AddRating(request request.RatingRequest) (float32, error) {
	recipe, err := s.GetRecipeByID(request.RecipeID)
	if err != nil {
		return 0, err
	}

	user, err := s.userRepository.GetUser(recipe.Author)
	if err != nil {
		return 0, err
	}

	newRatings := (recipe.Rating + request.Rating) / 2

	if recipe.Rating == 0 {
		newRatings = request.Rating
	}

	err = s.recipeRepository.AddRating(newRatings, request.RecipeID)
	if err != nil {
		return 0, err
	}

	if utils.CONFIG_EMAIL_SERVICE {
		body := fmt.Sprintf(models.RatingBodyEmailTemplate, user.Name, request.Rating, recipe.Title)

		emailPayload := notification.EmailPayload{
			To:      user.Email,
			Subject: "Rating Notification",
			Body:    body,
		}

		go s.notificationService.SendNotification(emailPayload)
	}

	return newRatings, nil
}
