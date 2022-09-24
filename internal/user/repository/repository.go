package repository

import (
	"context"
	"errors"
	"github.com/igilgyrg/todo-echo/internal/domain"
	"github.com/igilgyrg/todo-echo/internal/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) user.Repository {
	return &userMongoRepository{db: db, collection: db.Collection("users")}
}

func (a userMongoRepository) Store(ctx context.Context, user *domain.User) (domain.ID, error) {
	res, err := a.collection.InsertOne(ctx, user)
	if err != nil {
		return "", errors.New("error of inserting user to db")
	}

	if oid, ok := res.InsertedID.(string); ok {
		return domain.ID(oid), nil
	}

	return "", errors.New("error of parsing object id")
}

func (a userMongoRepository) Get(ctx context.Context, id domain.ID) (*domain.User, error) {
	var u *domain.User
	objectID := string(id)

	filter := bson.M{"_id": objectID}
	res := a.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	err := res.Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a userMongoRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	filter := bson.M{"email": bson.M{"$regex": primitive.Regex{Pattern: email, Options: "i"}}}
	var u *domain.User
	res := a.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a userMongoRepository) Update(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (a userMongoRepository) Delete(ctx context.Context, id domain.ID) error {
	//TODO implement me
	panic("implement me")
}
