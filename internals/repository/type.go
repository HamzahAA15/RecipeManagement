package repository

import (
	"context"
	"fmt"
	"log"
	"recipeApps/internals/models"
	"recipeApps/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IRecipeRepository interface {
	//recipe
	CreateRecipe(request models.Recipe) error
	GetRecipeByID(id string) (models.Recipe, error)
	GetRecipeByTitle(title string) (models.Recipe, error)
	GetAllRecipe() ([]models.Recipe, error)
	GetRecipesByFilter(filter bson.M) ([]models.Recipe, error)
	UpdateRecipeByID(request models.Recipe) error
	DeleteRecipeByID(id string) error
	//comment
	AddCommentByRecipeID(comment string, recipeID string) error
	//rating
	AddRating(rating float32, recipeID string) error
}

type IUserRepository interface {
	GetUser(name string) (*models.User, error)
	CreateUser(user *models.User) error
}

func ConnectMongoDB(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(utils.MongoDBURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB!")

	return client.Database("kaskus"), nil
}
