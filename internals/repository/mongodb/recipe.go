package repository

import (
	"context"
	"fmt"
	"recipeApps/internals/models"
	"recipeApps/internals/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepository struct {
	collection *mongo.Collection
}

func NewMongoDBRepository(collection *mongo.Collection) repository.IRecipeRepository {
	return &recipeRepository{
		collection: collection,
	}
}

func (r *recipeRepository) CreateRecipe(request models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeRepository) GetRecipeByTitle(title string) (models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var recipe models.Recipe
	filter := bson.M{"title": title, "deleted_at": nil}

	err := r.collection.FindOne(ctx, filter).Decode(&recipe)
	if err != nil && err.Error() != mongo.ErrNoDocuments.Error() {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func (r *recipeRepository) GetRecipeByID(id string) (models.Recipe, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Recipe{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID, "deleted_at": nil}

	var recipe models.Recipe
	err = r.collection.FindOne(ctx, filter).Decode(&recipe)
	if err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func (r *recipeRepository) GetRecipesByFilter(filter bson.M) ([]models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter["deleted_at"] = nil

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return []models.Recipe{}, fmt.Errorf("recipe not found")
		}

		return []models.Recipe{}, err
	}

	var recipes []models.Recipe
	if err = cur.All(ctx, &recipes); err != nil {
		return []models.Recipe{}, err
	}

	return recipes, nil
}

func (r *recipeRepository) GetAllRecipe() ([]models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := r.collection.Find(ctx, bson.M{"deleted_at": nil})
	if err != nil {
		return []models.Recipe{}, err
	}

	var recipes []models.Recipe
	if err = cur.All(ctx, &recipes); err != nil {
		return []models.Recipe{}, err
	}

	return recipes, nil
}

func (r *recipeRepository) UpdateRecipeByID(request models.Recipe) error {
	objectID, err := primitive.ObjectIDFromHex(request.ID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID, "deleted_at": nil}
	request.ID = ""
	update := bson.M{"$set": request}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeRepository) DeleteRecipeByID(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().UTC(),
		},
	}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
