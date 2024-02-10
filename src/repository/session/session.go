package session

import (
	"context"
	"github.com/mehrdadjalili/facegram_auth_service/resources/messages"
	"github.com/mehrdadjalili/facegram_auth_service/src/entity/models"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) Create(data *models.Session) error {
	if _, err := r.db.InsertOne(context.TODO(),
		bson.M{
			"_id":         data.Id,
			"ip":          data.IP,
			"device_id":   data.DeviceId,
			"user_id":     data.UserID,
			"device_name": data.DeviceName,
			"mac_address": data.MacAddress,
			"created_at":  data.CreatedAt,
			"timestamp":   data.TimeStamp,
		}); err != nil {
		utils.SubmitSentryLog(logSection, "Create", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) ById(id string) (*models.Session, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, derrors.InternalError()
	}
	var model *models.Session
	err = r.db.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, derrors.New(derrors.StatusNotFound, messages.NotFound)
		}
		utils.SubmitSentryLog(logSection, "ById", err)
		return nil, derrors.InternalError()
	}
	return model, nil
}

func (r *repository) Edit(data *models.Session) error {
	if _, err := r.ById(data.Id.Hex()); err != nil {
		return err
	}
	_, err := r.db.UpdateOne(context.TODO(), bson.M{"_id": data.Id}, bson.D{{Key: "$set", Value: bson.M{
		"ip":          data.IP,
		"device_id":   data.DeviceId,
		"user_id":     data.UserID,
		"device_name": data.DeviceName,
		"mac_address": data.MacAddress,
		"created_at":  data.CreatedAt,
		"timestamp":   data.TimeStamp,
	}}})
	if err != nil {
		utils.SubmitSentryLog(logSection, "Edit", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) DeleteById(id, userId string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	_userId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return derrors.InternalError()
	}
	_, err = r.db.DeleteOne(context.TODO(), bson.M{"_id": _id, "user_id": _userId})
	if err != nil {
		utils.SubmitSentryLog(logSection, "DeleteById", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) DeleteByIds(ids []string) error {
	var objects []primitive.ObjectID
	for _, item := range ids {
		_id, err := primitive.ObjectIDFromHex(item)
		if err != nil {
			return derrors.InternalError()
		}
		objects = append(objects, _id)
	}
	_, err := r.db.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": objects}})
	if err != nil {
		utils.SubmitSentryLog(logSection, "DeleteByIds", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) DeleteToDate(date int64) error {
	_, err := r.db.DeleteMany(context.TODO(), bson.M{"timestamp": bson.M{"$lte": date}})
	if err != nil {
		utils.SubmitSentryLog(logSection, "DeleteToDate", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) UserSessions(userId string, page, limit int) ([]models.Session, error) {
	_id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, derrors.InternalError()
	}
	var result []models.Session
	l := int64(limit)
	skip := int64(page*limit - limit)
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}
	c, err := r.db.Find(context.TODO(), bson.M{"user_id": _id}, &fOpt)
	if err != nil {
		utils.SubmitSentryLog(logSection, "UserSessions", err)
		return nil, derrors.InternalError()
	}
	err = c.All(context.TODO(), &result)
	if err != nil {
		utils.SubmitSentryLog(logSection, "UserSessions", err)
		return nil, derrors.InternalError()
	}
	return result, nil
}
