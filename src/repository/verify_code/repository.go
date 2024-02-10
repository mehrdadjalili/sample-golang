package verify_code

import (
	"context"
	repo "github.com/mehrdadjalili/facegram_auth_service/src/repository"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type repository struct {
	db *mongo.Collection
}

var logSection = "repository->verify_code"

func New(url, database string) (repo.VerifyCode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		utils.SubmitSentryLog(logSection, "New", err)
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		utils.SubmitSentryLog(logSection, "New", err)
		return nil, err
	}
	db := client.Database(database).Collection("verify_codes")
	return &repository{db: db}, nil
}
