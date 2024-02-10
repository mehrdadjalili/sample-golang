package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	VerifyCode struct {
		Id         primitive.ObjectID `json:"id" bson:"_id"`
		UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
		TypeCode   string             `json:"type_code" bson:"type_code"` //register | forgot | login
		CheckCount int                `json:"check_count" bson:"check_count"`
		Sender     string             `json:"sender" bson:"sender"` //email / sms / voice
		Code       string             `json:"code" bson:"code"`
		Receiver   string             `json:"receiver"  bson:"receiver"` //email or phone
		Timestamp  int64              `json:"timestamp" bson:"timestamp"`
	}
)
