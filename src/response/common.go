package response

type (
	Regular struct {
		Message string `json:"message"`
	}
	Status struct {
		Status bool `json:"status"`
	}
)
