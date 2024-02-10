package user

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

func (r *repository) Create(data *models.User) error {
	if _, err := r.db.InsertOne(context.TODO(), data); err != nil {
		utils.SubmitSentryLog(logSection, "Create", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) ById(id string) (*models.User, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, derrors.InternalError()
	}
	var model *models.User
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

func (r *repository) Edit(data *models.User) error {
	if _, err := r.ById(data.ID.Hex()); err != nil {
		return err
	}
	_, err := r.db.UpdateOne(context.TODO(), bson.M{"_id": data.ID}, bson.D{{Key: "$set", Value: data}})
	if err != nil {
		utils.SubmitSentryLog(logSection, "Edit", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) ExistsEmail(hashedEmail string) (bool, error) {
	var model *models.User
	err := r.db.FindOne(context.TODO(), bson.M{"email": hashedEmail}).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		utils.SubmitSentryLog(logSection, "ExistsEmail", err)
		return false, derrors.InternalError()
	}
	return true, nil
}

func (r *repository) ExistsPhone(hashedPhone string) (bool, error) {
	var model *models.User
	err := r.db.FindOne(context.TODO(), bson.M{"phone": hashedPhone}).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		utils.SubmitSentryLog(logSection, "ExistsPhone", err)
		return false, derrors.InternalError()
	}
	return true, nil
}

func (r *repository) ByEmailOrPhone(user string) (*models.User, error) {
	var model *models.User
	err := r.db.FindOne(context.TODO(), bson.M{"$or": bson.A{
		bson.M{"phone": user},
		bson.M{"email": user},
	}}).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, derrors.New(derrors.StatusNotFound, messages.NotFound)
		}
		utils.SubmitSentryLog(logSection, "ByEmailOrPhone", err)
		return nil, derrors.InternalError()
	}
	return model, nil
}

func (r *repository) CountByRole(isAgent bool) (int64, error) {
	count, err := r.db.CountDocuments(context.TODO(), bson.M{"is_agent": isAgent})
	if err != nil {
		utils.SubmitSentryLog(logSection, "CountByRole", err)
		return 0, derrors.InternalError()
	}
	return count, nil
}

func (r *repository) Count() (int64, error) {
	count, err := r.db.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		utils.SubmitSentryLog(logSection, "Count", err)
		return 0, derrors.InternalError()
	}
	return count, nil
}

func (r *repository) List(search, sort string, page, perPage int) ([]models.User, int64, error) {
	var result []models.User
	offset := (page - 1) * perPage
	var s int
	if s = -1; sort == "asc" {
		s = 1
	}
	filter := bson.D{{"deleted_at", bson.D{{"$exists", false}}}}
	if search != "" {
		hash := utils.NewSHA256([]byte(search))
		filter = append(filter, bson.E{Key: "phone", Value: bson.D{
			{"$regex", primitive.Regex{Pattern: hash, Options: "i"}},
		}})
		filter = append(filter, bson.E{Key: "email", Value: bson.D{
			{"$regex", primitive.Regex{Pattern: hash, Options: "i"}},
		}})
	}
	if perPage == -1 {
		perPage = 0
	}
	pipeline := []bson.M{
		{"$match": bson.M{"deleted_at": bson.D{{"$exists", false}}}},
		{"$sort": bson.M{"_id": s}},
		{"$skip": offset},
		{"$limit": perPage},
	}
	c, err := r.db.Aggregate(context.TODO(), pipeline)
	if err != nil {
		utils.SubmitSentryLog(logSection, "List", err)
		return nil, 0, derrors.InternalError()
	}
	err = c.All(context.TODO(), &result)
	if err != nil {
		utils.SubmitSentryLog(logSection, "List", err)
		return nil, 0, derrors.InternalError()
	}
	count, err := r.db.CountDocuments(context.TODO(), filter)
	if err != nil {
		utils.SubmitSentryLog(logSection, "List", err)
		return nil, 0, derrors.InternalError()
	}
	return result, count, nil
}
