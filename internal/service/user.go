package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"synapsis/internal/model"
	"synapsis/internal/repository"
	"time"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Register(user model.Register) error
	Login(login model.Login) (string, error)
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

func (s *userService) Login(login model.Login) (string, error) {
	password, err := s.userRepository.GetPasswordByEmail(login.Email)
	if err != nil {
		log.Println("error: " + err.Error())
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
	if err != nil {
		log.Println("error: " + err.Error())
		return "", err
	}

	user, err := s.userRepository.GetUserByEmail(login.Email)
	if err != nil {
		log.Println("error: " + err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRETKEY")))
	if err != nil {
		log.Println("error: " + err.Error())
		return "", err
	}

	return tokenString, nil
}
