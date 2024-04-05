package mapper

import (
	"devstream.in/pixelated-pipeline/api/dto"
	"devstream.in/pixelated-pipeline/database/entities"
	"devstream.in/pixelated-pipeline/services/models"
)

func UserModelToEntity(userModel models.User) entities.User {
	return entities.User{}
}

func UserEntityToModel(userEntity entities.User) models.User {
	return models.User{}
}

func UserDtoToModel(userDto dto.UserDTO) models.User {
	return models.User{}
}

func UserModelToDto(userModel models.User) dto.UserDTO {
	return dto.UserDTO{}
}
