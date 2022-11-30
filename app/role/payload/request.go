package payload

type (
	RoleCreate struct {
		Role string `json:"role" binding:"required"`
	}
	RoleUpdate struct {
		ID   uint   `json:"id" binding:"required"`
		Role string `json:"role" binding:"required"`
	}
	RoleDelete struct {
		ID uint `json:"id" binding:"required"`
	}
)
