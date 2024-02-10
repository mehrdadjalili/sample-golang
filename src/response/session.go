package response

import (
	"time"
)

type (
	ConnectionStatus struct {
		IsOnline  bool      `json:"is_online"`
		LastVisit time.Time `json:"last_visit"`
	}
	Session struct {
		Id          string    `json:"id"`
		DeviceName  string    `json:"device_name"` //name and model of device
		CreatedAt   time.Time `json:"created_at"`
		Location    string    `json:"location"`
		ThisSession bool      `json:"this_session"`
	}
	SessionList struct {
		Data []Session `json:"data"`
	}
)
