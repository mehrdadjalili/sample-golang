package response

type (
	User struct {
		ID          string `json:"id"`
		Username    string `json:"username"`
		Status      bool   `json:"status"`
		PhoneStatus bool   `json:"phone_status"`
		EmailStatus bool   `json:"email_status"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		IsAgent     bool   `json:"is_agent"`
		Gender      string `json:"gender"`
		Age         int    `json:"age"`
	}
)
