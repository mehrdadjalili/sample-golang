package request

type (
	User struct {
		ID          string `json:"id"`
		Status      bool   `json:"status"`
		PhoneStatus bool   `json:"phone_status"`
		EmailStatus bool   `json:"email_status"`
		Password    string `json:"password"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		IsAgent     bool   `json:"is_agent"`
		Gender      string `json:"gender"`
		Age         int    `json:"age"`
		Username    string `json:"username"`
	}
)
