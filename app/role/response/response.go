package response

import (
	"github.com/Crunchy89/go-mysql/app/role/payload"
	"github.com/Crunchy89/go-mysql/domain"
)

func SingleResponse(data *domain.Role) *payload.RoleResponse {
	response := &payload.RoleResponse{
		UUID: data.UUID,
		Role: data.Role,
	}
	return response
}

func BatchResponse(datas []*domain.Role) []*payload.RoleResponse {
	response := []*payload.RoleResponse{}
	for _, data := range datas {
		newData := &payload.RoleResponse{
			UUID: data.UUID,
			Role: data.Role,
		}
		response = append(response, newData)
	}
	return response
}
