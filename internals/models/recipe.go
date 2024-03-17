package models

import "time"

type Recipe struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string    `bson:"title" json:"title"`
	Ingredients []string  `bson:"ingredients" json:"ingredients"`
	Procedures  []string  `bson:"procedures" json:"procedures"`
	Description string    `bson:"description" json:"description"`
	Category    string    `bson:"category" json:"category"`
	Author      string    `bson:"author" json:"author"`
	ImageURL    string    `bson:"image_url" json:"image_url"`
	Comments    []string  `bson:"comments" json:"comments"`
	Rating      float32   `bson:"rating" json:"rating"`
	CreatedAt   time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt   time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

const CommentBodyEmailTemplate = `<p>Dear %s,</p>
<p>Someone Commented %s on your recipe %s</p>
`

const RatingBodyEmailTemplate = `<p>Dear %s,</p>
<p>Someone Rated %f on your recipe %s</p>
`
