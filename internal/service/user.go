package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"synapsis/internal/model"
	"synapsis/internal/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Register(user model.Register) error
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Register(user model.Register) error {
	emailDB, err := s.userRepository.CheckEmail(user.Email)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	fmt.Println(emailDB)

	if emailDB != "" {
		log.Println("error: " + "email has been registered")
		return errors.New("email has been registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	newUser := model.User{
		Id:       uuid.New(),
		Email:    user.Email,
		Password: string(hashedPassword),
		Name:     user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
	}

	if err := s.userRepository.InsertUser(newUser); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}
