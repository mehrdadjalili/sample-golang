package response

import (
	"time"
)

type (
	SessionInfo struct {
		Id         string    `json:"id"`
		DeviceName string    `json:"device_name"` //name and model of device
		DeviceId   string    `json:"device_id"`
		MacAddress string    `json:"mac_address"`
		CreatedAt  time.Time `json:"created_at"`
		TimeStamp  int64     `json:"timestamp"`
		IP         string    `json:"ip"`
	}
	UserInfo struct {
		ID          string    `json:"id"`
		Username    string    `json:"username"`
		PhoneStatus bool      `json:"phone_status"`
		EmailStatus bool      `json:"email_status"`
		Phone       string    `json:"phone"`
		Email       string    `json:"email"`
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		IsAgent     bool      `json:"is_agent"`
		Gender      string    `json:"gender"`
		Status      bool      `json:"status"`
		Age         int       `json:"age"`
		CreatedAt   time.Time `json:"created_at"`
	}
	CheckToken struct {
		User    UserInfo    `json:"user"`
		Session SessionInfo `json:"session"`
	}
	ResultData struct {
		Result string      `json:"result"`
		Data   interface{} `json:"data"`
	}
)
