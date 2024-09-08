package service

import (
	"blogging-platform/dto"
	"blogging-platform/helper"
	"blogging-platform/model"
	"blogging-platform/repository"
)

type UserService interface {
	Register(dto.UserRegisterDTO) (model.User, error)
	Login(dto.UserLoginDTO) (dto.Token, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{userRepository: repo}
}

func (s *userService) Register(userDTO dto.UserRegisterDTO) (model.User, error) {
	hashedPassword, err := helper.HashPassword(userDTO.Password)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: hashedPassword,
	}

	return s.userRepository.Register(user)
}

func (s *userService) Login(userDTO dto.UserLoginDTO) (dto.Token, error) {
	existUser, err := s.userRepository.FindByUsername(userDTO.Username)
	if err != nil {
		return dto.Token{}, err
	}

	err = helper.ComparePass(existUser.Password, userDTO.Password)
	if err != nil {
		return dto.Token{}, err
	}

	token, err := helper.GenerateToken(existUser)
	if err != nil {
		return dto.Token{}, err
	}

	return token, nil
}
