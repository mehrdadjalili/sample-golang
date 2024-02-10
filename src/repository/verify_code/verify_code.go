package verify_code

import (
	"context"
	"github.com/mehrdadjalili/facegram_auth_service/resources/messages"
	"github.com/mehrdadjalili/facegram_auth_service/src/entity/models"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) Create(data *models.VerifyCode) error {
	if _, err := r.db.InsertOne(context.TODO(), data); err != nil {
		utils.SubmitSentryLog(logSection, "Create", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) ById(id string) (*models.VerifyCode, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, derrors.InternalError()
	}
	var model models.VerifyCode
	err = r.db.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, derrors.New(derrors.StatusNotFound, messages.NotFound)
		}
		utils.SubmitSentryLog(logSection, "ById", err)
		return nil, derrors.InternalError()
	}
	return &model, nil
}

func (r *repository) Edit(data *models.VerifyCode) error {
	if _, err := r.ById(data.Id.Hex()); err != nil {
		return err
	}
	_, err := r.db.UpdateOne(context.TODO(), bson.D{{"_id", data.Id}}, bson.D{{Key: "$set", Value: data}})
	if err != nil {
		utils.SubmitSentryLog(logSection, "Edit", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) DeleteById(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return derrors.InternalError()
	}
	_, err = r.db.DeleteOne(context.TODO(), bson.D{{"_id", _id}})
	if err != nil {
		utils.SubmitSentryLog(logSection, "DeleteById", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) DeleteToDate(date int64) error {
	_, err := r.db.DeleteMany(context.TODO(), bson.D{{"timestamp", bson.M{"$lte": date}}})
	if err != nil {
		utils.SubmitSentryLog(logSection, "DeleteToDate", err)
		return derrors.InternalError()
	}
	return nil
}
