package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *recipeRepository) AddCommentByRecipeID(comment string, recipeID string) error {
	// Convert recipeID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filter to find the recipe by ID
	filter := bson.M{"_id": objectID}

	// Update to push the new comment to the comments array
	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}

	// Perform the update operation
	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
