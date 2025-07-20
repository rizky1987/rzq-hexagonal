package response

import "rzq-hexagonal/domain/entity"

type RegisterResponse struct {
	Id    string `json:"Id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func ConvertUserFromEntity(user *entity.User) RegisterResponse {

	return RegisterResponse{
		Id:    user.Id.String(),
		Email: user.Email,
		Name:  user.Name,
	}
}
