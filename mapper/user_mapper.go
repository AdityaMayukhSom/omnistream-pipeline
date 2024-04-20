package mapper

import (
	"devstream.in/pixelated-pipeline/api/dto"
	"devstream.in/pixelated-pipeline/database/repositories"
	"devstream.in/pixelated-pipeline/services/models"
)

func UserModelToEntity(userModel models.User) repositories.UserEntity {
	return repositories.UserEntity{}
}

func UserEntityToModel(userEntity repositories.UserEntity) models.User {
	return models.User{}
}

func UserDtoToModel(userDto dto.UserDTO) models.User {
	return models.User{}
}

func UserModelToDto(userModel models.User) dto.UserDTO {
	return dto.UserDTO{}
}
