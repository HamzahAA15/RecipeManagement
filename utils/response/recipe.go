package response

type RecipeResponseData struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `bson:"title" json:"title"`
	Ingredients []string `bson:"ingredients,omitempty" json:"ingredients,omitempty"`
	Procedures  []string `bson:"procedures,omitempty" json:"procedures,omitempty"`
	Description string   `bson:"description,omitempty" json:"description,omitempty"`
	Category    string   `bson:"category,omitempty" json:"category,omitempty"`
	Author      string   `bson:"author" json:"author"`
	ImageURL    string   `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Comments    []string `bson:"comments,omitempty" json:"comments,omitempty"`
	Rating      float32  `bson:"rating,omitempty" json:"rating,omitempty"`
}

type UserResponseData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
