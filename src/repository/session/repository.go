package session

import (
	"context"
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

var logSection = "repository->session"

func New(url, database string) (repo.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		utils.SubmitSentryLog(logSection, "Delete", err)
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		utils.SubmitSentryLog(logSection, "Delete", err)
		return nil, err
	}
	db := client.Database(database).Collection("sessions")
	idx := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_id", Value: 1}},
		Options: options.Index(),
	}
	_, err = db.Indexes().CreateOne(context.Background(), idx)
	if err != nil {
		utils.SubmitSentryLog(logSection, "Delete", err)
		return nil, err
	}
	return &repository{db: db}, nil
}
