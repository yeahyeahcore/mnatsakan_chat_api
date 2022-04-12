package service

import (
	"fmt"
	"strconv"
	"time"

	"mnatsakan_chat_api/internal/dto"
	"mnatsakan_chat_api/internal/model"
	"mnatsakan_chat_api/internal/repository"
	"mnatsakan_chat_api/pkg/codec"
	"mnatsakan_chat_api/pkg/utils"

	"github.com/golang-jwt/jwt"
)

const (
	tokenTTL = 12 * time.Hour
)

// Auth ...
type Auth struct {
	repository repository.UserRepository
	codec      *codec.Codec
}

// NewAuthService ...
func NewAuthService(repository repository.UserRepository, codec *codec.Codec) *Auth {
	return &Auth{
		repository: repository,
		codec:      codec,
	}
}

// Login ...
func (receiver *Auth) Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	findUserArguments := map[string]string{"login": loginRequest.Login}

	user, err := receiver.repository.FindOne(findUserArguments)
	if err != nil {
		return nil, err
	}

	if isPasswordCorrect := utils.CheckPasswordHash(loginRequest.Password, user.Password); isPasswordCorrect == false {
		return nil, fmt.Errorf("password is incorrect")
	}

	token, err := receiver.generateToken(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}, nil
}

// Register ...
func (receiver *Auth) Register(registerRequest *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	passwordHash, err := utils.GeneratePasswordHash(registerRequest.Password)
	if err != nil {
		return nil, err
	}

	user, err := receiver.repository.Insert(&model.User{
		Email:    registerRequest.Email,
		Password: passwordHash,
		Username: registerRequest.Username,
		Login:    registerRequest.Login,
	})

	return &dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
	}, err
}

func (receiver *Auth) generateToken(data string) (string, error) {
	token := jwt.New(receiver.codec.SignMethod)

	standardClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	token.Claims = &dto.TokenClaims{
		StandardClaims: standardClaims,
		Data:           data,
	}

	return token.SignedString([]byte(receiver.codec.SecretKey))
}
