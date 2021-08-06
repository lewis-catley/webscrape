package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// URL The complete document that's stored in mongo
type URL struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL        string             `bson:"url" json:"url"`
	Results    []URLSFound        `bson:"results,omitempty" json:"results"`
	Count      int                `bson:"count" json:"count"`
	IsComplete bool               `bson:"isComplete" json:"isComplete"`
}

type URLPost struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type URLSFound struct {
	URL       string   `bson:"url" json:"url"`
	URLSFound []string `bson:"urlsFound" json:"urlsFound"`
}
