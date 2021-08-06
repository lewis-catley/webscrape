package mongodb

import (
	"context"

	"github.com/lewis-catley/webscrape/scraper/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// URLModel represent mongo session with URL data
type URLModel struct {
	C *mongo.Collection
}

// Update the object that has the passed ID
func (m *URLModel) Update(id string, u models.URL) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": p}
	return m.C.ReplaceOne(context.TODO(), filter, u)
}

// UpdateURLSFound will add the found urls to the object
func (m *URLModel) UpdateURLSFound(id string, found models.URLSFound) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": p}
	update := bson.M{
		"$push": bson.M{
			"results": found,
		},
	}
	return m.C.UpdateOne(context.TODO(), filter, update)
}
