package response

import "time"

type (
	UserStatistics struct {
		TotalLives           int64 `json:"total_lives"`
		LiveDuration         int64 `json:"live_duration"`
		TotalReceiveMessages int64 `json:"total_receive_messages"`
		TotalSendMessages    int64 `json:"total_send_messages"`
		Score                int64 `json:"score" bson:"score"`
		TotalLikes           int64 `json:"total_likes"`
		TotalFavorites       int64 `json:"total_favorites"`
		TotalProfileVisit    int64 `json:"total_profile_visit"`
	}
	Profile struct {
		ID               string           `json:"id"`
		Username         string           `json:"username"`
		PhoneStatus      bool             `json:"phone_status"`
		EmailStatus      bool             `json:"email_status"`
		FirstName        string           `json:"first_name"`
		LastName         string           `json:"last_name"`
		Gender           string           `json:"gender"`
		Age              int              `json:"age"`
		Email            string           `json:"email"`
		Phone            string           `json:"phone"`
		IsAgent          bool             `json:"is_agent"`
		Statistics       UserStatistics   `json:"statistics"`
		ConnectionStatus ConnectionStatus `json:"connection_status"`
		CreatedAt        time.Time        `json:"created_at"`
	}
)
