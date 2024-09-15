package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

// Convierte un modelo de dominio a un DTO
func ToUserDTO(user *models.User) dtos.UserResponse {
	if user == nil {
		return dtos.UserResponse{}
	}
	return dtos.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToUserDTOs(users []models.User) []dtos.UserResponse {
	userDTOs := make([]dtos.UserResponse, len(users))
	for i, user := range users {
		// Mapear los campos directamente, aplicando transformaciones si es necesario
		userDTOs[i] = dtos.UserResponse{
			Id:    user.ID,
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
		}
	}
	return userDTOs
}

// Convierte un DTO en un modelo de dominio
func ToUserModel(user dtos.UserRequest) models.User {
	return models.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	}
}
