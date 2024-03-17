package httpservice

import (
	"context"
	httpserviceRecipe "recipeApps/internals/httpservice/recipe"
	httpserviceUser "recipeApps/internals/httpservice/user"
	"recipeApps/internals/repository"
	inMemoryRepo "recipeApps/internals/repository/inMemory"
	mongoRepo "recipeApps/internals/repository/mongodb"
	"recipeApps/internals/service/recipe"
	userService "recipeApps/internals/service/user"
	"recipeApps/utils"
	"recipeApps/utils/middleware/authentication"
	"recipeApps/utils/notification"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

func InitRecipeRoute(app *fiber.App) {
	mongoDBCollection, err := repository.ConnectMongoDB(context.Background())
	if err != nil {
		panic(err)
	}

	userRepo := inMemoryRepo.NewUserInMemoryRepository()
	auth := authentication.AuthMiddleware(userRepo)
	gmailSMTP := gomail.NewDialer(utils.CONFIG_SMTP_HOST, utils.CONFIG_SMTP_PORT, utils.CONFIG_AUTH_EMAIL, utils.CONFIG_AUTH_PASSWORD)
	notificationService := notification.NewGmailNotification(gmailSMTP)

	recipeRepo := mongoRepo.NewMongoDBRepository(mongoDBCollection.Collection("recipe"))
	recipeService := recipe.NewRecipeService(recipe.RecipeServiceConfig{
		RecipeRepository:    recipeRepo,
		UserRepository:      userRepo,
		NotificationService: notificationService,
	})

	recipeHandler := httpserviceRecipe.NewRecipeHandler(recipeService)
	recipeGroup := app.Group("/api/recipes")
	commentGroup := app.Group("/api/recipes/comments")
	ratingGroup := app.Group("/api/recipes/ratings")

	//recipe domain route
	// need authentication
	recipeGroup.Post("/", auth, recipeHandler.CreateRecipe)
	recipeGroup.Patch("/", auth, recipeHandler.UpdateRecipe)
	recipeGroup.Delete("/", auth, recipeHandler.DeleteRecipe)

	recipeGroup.Get("/", recipeHandler.GetRecipeByID)
	recipeGroup.Get("/all", recipeHandler.GetAllRecipes)
	recipeGroup.Get("/filter", recipeHandler.GetRecipesByFilter)

	//comment domain route
	commentGroup.Patch("/", recipeHandler.AddComment)

	//rating domain route
	ratingGroup.Patch("/", recipeHandler.AddRating)
}

func InitUserRoute(app *fiber.App) {
	userRepo := inMemoryRepo.NewUserInMemoryRepository()
	userService := userService.NewUserService(userRepo)
	userHandler := httpserviceUser.NewUserHandler(userService)
	userGroup := app.Group("/api/user")

	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Get("/", userHandler.GetUserByName)
}
