package service

import (
	"blogging-platform/dto"
	"blogging-platform/helper"
	"blogging-platform/model"
	"blogging-platform/repository"
)

type UserService interface {
	Register(dto.UserRegisterDTO) (model.User, error)
	Login(dto.UserLoginDTO) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{userRepository: repo}
}

func (us *userService) Register(userDTO dto.UserRegisterDTO) (model.User, error) {
	hashedPassword, err := helper.HashPassword(userDTO.Password)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: hashedPassword,
	}

	return us.userRepository.Register(user)
}

func (us *userService) Login(userDTO dto.UserLoginDTO) (model.User, error) {
	existUser, err := us.userRepository.FindByUsername(userDTO.Username)
	if err != nil {
		return model.User{}, err
	}

	err = helper.ComparePass(existUser.Password, userDTO.Password)
	if err != nil {
		return model.User{}, err
	}

	return existUser, nil
}
