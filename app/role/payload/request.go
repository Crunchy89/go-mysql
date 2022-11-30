package payload

type (
	RoleCreate struct {
		Role string `json:"role" binding:"required"`
	}
	RoleUpdate struct {
		UUID string `json:"uuid" binding:"required"`
		Role string `json:"role" binding:"required"`
	}
	RoleDelete struct {
		UUID string `json:"uuid" binding:"required"`
	}
)
