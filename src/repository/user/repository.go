package user

import (
	"context"
	"github.com/mehrdadjalili/facegram_auth_service/src/config"
	repo "github.com/mehrdadjalili/facegram_auth_service/src/repository"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type repository struct {
	db *mongo.Collection
}

var logSection = "repository->user"

func New(cfg config.MongoDb) (repo.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Url))
	if err != nil {
		utils.SubmitSentryLog(logSection, "New", err)
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		utils.SubmitSentryLog(logSection, "New", err)
		return nil, err
	}
	db := client.Database(cfg.Database).Collection("users")
	idx := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "phone", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "is_agent", Value: 1}},
			Options: options.Index(),
		},
	}
	for _, item := range idx {
		_, err = db.Indexes().CreateOne(context.Background(), item)
		if err != nil {
			utils.SubmitSentryLog(logSection, "New", err)
			return nil, err
		}
	}
	return &repository{db: db}, nil
}
