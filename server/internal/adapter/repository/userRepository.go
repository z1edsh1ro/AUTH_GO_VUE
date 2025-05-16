package repository

import (
	"context"
	"main/internal/domain/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	Collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserCollection {
	collection := client.Database("auth_db").Collection("user")

	return &UserCollection{Collection: collection}
}

func (repo *UserCollection) GetAll() ([]model.User, error) {
	var users []model.User

	cursor, err := repo.Collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo model.User
		cursor.Decode(&todo)
		users = append(users, todo)
	}

	return users, nil

}

func (repo *UserCollection) GetByEmail(email string) (model.User, error) {
	var user model.User

	err := repo.Collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserCollection) Create(entry model.User) error {
	_, err := repo.Collection.InsertOne(context.TODO(), model.User{
		Name:      entry.Name,
		Email:     entry.Email,
		Password:  entry.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return err
}
