package payload

type (
	UserResponse struct {
		UUID     string `json:"uuid,omitempty"`
		Username string `json:"username,omitempty"`
	}
)
