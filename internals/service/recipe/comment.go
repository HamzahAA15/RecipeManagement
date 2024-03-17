package recipe

import (
	"fmt"
	"recipeApps/internals/models"
	"recipeApps/utils"
	"recipeApps/utils/notification"
	"recipeApps/utils/request"
)

func (s *recipeService) AddComment(request request.CommentRequest) error {
	recipe, err := s.recipeRepository.GetRecipeByID(request.RecipeID)
	if err != nil {
		return err
	}

	user, err := s.userRepository.GetUser(recipe.Author)
	if err != nil {
		return err
	}

	err = s.recipeRepository.AddCommentByRecipeID(request.Comment, request.RecipeID)
	if err != nil {
		return err
	}

	if utils.CONFIG_EMAIL_SERVICE {
		body := fmt.Sprintf(models.CommentBodyEmailTemplate, user.Name, request.Comment, recipe.Title)

		emailPayload := notification.EmailPayload{
			To:      user.Email,
			Subject: "Comment Notification",
			Body:    body,
		}

		go s.notificationService.SendNotification(emailPayload)
	}

	return nil
}
