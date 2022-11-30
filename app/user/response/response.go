package response

import (
	"github.com/Crunchy89/go-mysql/app/user/payload"
	"github.com/Crunchy89/go-mysql/domain"
)

func SingleResponse(data *domain.User) *payload.UserResponse {
	response := &payload.UserResponse{
		UUID:     data.UUID,
		Username: data.Username,
	}
	return response
}

func BatchResponse(datas []*domain.User) []*payload.UserResponse {
	response := []*payload.UserResponse{}
	for _, data := range datas {
		newData := &payload.UserResponse{
			UUID:     data.UUID,
			Username: data.Username,
		}
		response = append(response, newData)
	}
	return response
}
