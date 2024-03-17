package request

type RecipeRequest struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Procedures  []string `json:"procedures"`
	Category    string   `json:"category"`
	ImageURL    string   `json:"image_url"`
}

type CommentRequest struct {
	RecipeID string `json:"recipe_id"`
	Comment  string `json:"comment"`
}

type RatingRequest struct {
	RecipeID string  `json:"recipe_id"`
	Rating   float32 `json:"rating"`
}
