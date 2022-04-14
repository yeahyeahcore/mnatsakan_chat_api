package service

import (
	"errors"
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

// Auth service
type Auth struct {
	repository repository.UserRepository
	codec      *codec.Codec
}

// NewAuthService - returns new auth service instance
func NewAuthService(repository repository.UserRepository, codec *codec.Codec) *Auth {
	return &Auth{
		repository: repository,
		codec:      codec,
	}
}

// Login user
func (receiver *Auth) Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	findUserArguments := map[string]string{"login": loginRequest.Login}

	user, err := receiver.repository.FindOne(findUserArguments)
	if err != nil {
		return nil, err
	}

	if isPasswordCorrect := utils.CheckPasswordHash(loginRequest.Password, user.Password); isPasswordCorrect == false {
		return nil, fmt.Errorf("password is incorrect")
	}

	token, err := receiver.GenerateToken(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}, nil
}

// Register user
func (receiver *Auth) Register(registerRequest *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	passwordHash, err := utils.GeneratePasswordHash(registerRequest.Password)
	if err != nil {
		return nil, err
	}

	user, err := receiver.repository.FindOne(map[string]string{"login": registerRequest.Login})
	if user != nil {
		return nil, fmt.Errorf("user exist")
	}

	newUser, err := receiver.repository.Insert(&model.User{
		Email:    registerRequest.Email,
		Password: passwordHash,
		Username: registerRequest.Username,
		Login:    registerRequest.Login,
	})

	return &dto.RegisterResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
	}, err
}

// ParseToken - parsing token and returns users id
func (receiver *Auth) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &dto.TokenClaims{}, receiver.checkTokenSignMethod)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*dto.TokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *dto.TokenClaims")
	}

	return claims.UserID, nil
}

// GenerateToken - generates token
func (receiver *Auth) GenerateToken(userID string) (string, error) {
	token := jwt.New(receiver.codec.SignMethod)

	standardClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	token.Claims = &dto.TokenClaims{
		StandardClaims: standardClaims,
		UserID:         userID,
	}

	return token.SignedString([]byte(receiver.codec.SecretKey))
}

func (receiver *Auth) checkTokenSignMethod(token *jwt.Token) (interface{}, error) {
	if token.Method.Alg() != receiver.codec.SignMethod.Name {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
	}

	return []byte(receiver.codec.SecretKey), nil
}
