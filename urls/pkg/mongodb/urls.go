package mongodb

import (
	"context"
	"errors"

	"github.com/lewis-catley/webscrape/urls/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// URLModel represent mongo session with URL data
type URLModel struct {
	C *mongo.Collection
}

// All retreive all URL objects
func (m *URLModel) All() ([]models.URL, error) {
	ctx := context.TODO()
	urls := []models.URL{}

	// Find all URLS
	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &urls)
	if err != nil {
		return nil, err
	}

	return urls, err
}

func (m *URLModel) GetByID(id string) (*models.URL, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find user by id
	url := &models.URL{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(url)
	if err != nil {
		// Check if the URL is not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments") // TODO: This wants changing to be a NotFoundErr
		}

		return nil, err
	}

	return url, nil
}

// Insert insert a new URL object
func (m *URLModel) Insert(u models.URLPost) (*mongo.InsertOneResult, error) {
	url := models.URL{
		URL: u.URL,
	}
	return m.C.InsertOne(context.TODO(), url)
}

func (m *URLModel) DeleteAll() error {
	return m.C.Drop(context.TODO())
}

// TODO: Think about adding a delete method
