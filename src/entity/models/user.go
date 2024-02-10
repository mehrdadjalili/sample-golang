package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	ConnectionStatus struct {
		IsOnline  bool      `json:"is_online" bson:"is_online"`
		LastVisit time.Time `json:"last_visit" bson:"last_visit"`
	}
	User struct {
		ID               primitive.ObjectID `json:"id" bson:"_id"`
		Username         string             `json:"username" bson:"username"`
		PhoneStatus      bool               `json:"phone_status" bson:"phone_status"`
		EmailStatus      bool               `json:"email_status" bson:"email_status"`
		Password         string             `json:"password" bson:"password"`
		Phone            string             `json:"phone" bson:"phone"`
		Email            string             `json:"email" bson:"email"`
		FirstName        string             `json:"first_name" bson:"first_name"`
		LastName         string             `json:"last_name" bson:"last_name"`
		IsAgent          bool               `json:"is_agent" bson:"is_agent"`
		Gender           string             `json:"gender" bson:"gender"`
		Age              int                `json:"age" bson:"age"`
		Status           bool               `json:"status" bson:"status"`
		Avatar           *string            `json:"avatar" bson:"avatar"`
		ConnectionStatus ConnectionStatus   `json:"connection_status" bson:"connection_status"`
		Statistics       UserStatistics     `json:"statistics" bson:"statistics"`
		CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
	}
	UserStatistics struct {
		TotalLives           int64 `json:"total_lives" bson:"total_lives"`
		LiveDuration         int64 `json:"live_duration" bson:"live_duration"`
		TotalReceiveMessages int64 `json:"total_receive_messages" bson:"total_receive_messages"`
		TotalSendMessages    int64 `json:"total_send_messages" bson:"total_send_messages"`
		Score                int64 `json:"score" bson:"score"`
		TotalLikes           int64 `json:"total_likes" bson:"total_likes"`
		TotalFavorites       int64 `json:"total_favorites" bson:"total_favorites"`
		TotalProfileVisit    int64 `json:"total_profile_visit" bson:"total_profile_visit"`
	}
)
