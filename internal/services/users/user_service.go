package userservice

import (
	"companypresence-api/internal/models"
	"companypresence-api/internal/repositories"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepository	
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	
	return &UserService{
		userRepository: userRepository,
	}
}

func(s *UserService)CreateUser(ctx context.Context, user *models.User)error{
	err := s.userRepository.CreateUSer(ctx, user)	
	return err
}

func(s *UserService)GetUserById(ctx context.Context, id string)(user models.User, err error){
	user, err = s.userRepository.GetUserById(ctx, id)
	return user, err
}

func(s *UserService)GetUserByMail(ctx context.Context, mail string)(user models.User, err error){
	user, err = s.userRepository.GetUserByMail(ctx, mail)
	return user, err
}

func (s *UserService)GetActiveUsers(ctx context.Context)(users []models.User, err error){	 
	return s.userRepository.GetUsers(ctx, true)
}
func (s *UserService)GetAllUsers(ctx context.Context)(users []models.User, err error){
	return s.userRepository.GetUsers(ctx, false)
}

func (s *UserService)ValidateCredentials(ctx context.Context, email, password string)(models.User, error){
	user, err :=s.userRepository.GetUserByMail(ctx, email)
	if err != nil {
		return models.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *UserService)UpdateUser(ctx context.Context, user *models.User)error{
	return s.userRepository.UpdateUser(ctx, user)
}

func (s *UserService)DeleteUser(ctx context.Context, id string)error {
	return s.userRepository.DeleteUser(ctx, id)
}