package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

// Convierte un modelo de dominio a un DTO
func ToUserDTO(user *models.User) dtos.User {
	if user == nil {
		return dtos.User{}
	}
	return dtos.User{
		Id:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToUserDTOs(users []models.User) []dtos.User {
	userDTOs := make([]dtos.User, len(users))
	for i, user := range users {
		// Mapear los campos directamente, aplicando transformaciones si es necesario
		userDTOs[i] = dtos.User{
			Id:    user.ID,
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
		}
	}
	return userDTOs
}

// Convierte un DTO en un modelo de dominio
func ToUserModel(user dtos.User) models.User {
	return models.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	}
}
