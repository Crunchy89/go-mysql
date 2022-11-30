package payload

type (
	UserCreate struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		RoleID   uint   `json:"role_id" binding:"required"`
	}
	UserUpdate struct {
		ID       uint    `json:"id" binding:"required"`
		RoleID   *uint   `json:"role_id,omitempty"`
		Password *string `json:"password,omitempty"`
	}
	UserDelete struct {
		ID uint `json:"id" binding:"required"`
	}
)
