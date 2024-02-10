package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Session struct {
		Id         primitive.ObjectID `json:"id" bson:"_id"`
		DeviceName string             `json:"device_name" bson:"device_name"` //name and model of device
		DeviceId   string             `json:"device_id" bson:"device_id"`
		MacAddress string             `json:"mac_address" bson:"mac_address"`
		UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
		CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
		TimeStamp  int64              `json:"timestamp" bson:"timestamp"`
		IP         string             `json:"ip" bson:"ip"`
	}
)
