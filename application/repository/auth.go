package repository

import (
	"context"
	"course-api/domain/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	collection *mongo.Collection
}

func NewAuthRepository(collection *mongo.Collection) *AuthRepository {
	return &AuthRepository{
		collection: collection,
	}
}

func (ar *AuthRepository) Login(creds model.LoginRequest) (*model.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	var user *model.UserModel = new(model.UserModel)
	err := ar.collection.FindOne(ctx, bson.M{
		"username": creds.User},
	).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create - insert user register
//
// Params:
//   - payload model.UserModel
//
// Return:
//   - error
func (ar *AuthRepository) Create(payload model.UserModel) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	_, err := ar.collection.InsertOne(ctx, payload)
	return err
}
