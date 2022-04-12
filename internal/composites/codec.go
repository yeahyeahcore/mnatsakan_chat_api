package composites

import (
	"mnatsakan_chat_api/pkg/codec"
	"mnatsakan_chat_api/pkg/utils"

	"github.com/golang-jwt/jwt"
)

// GetCodecInstance ...
func GetCodecInstance() *codec.Codec {
	secretKey := utils.GenerateRandomString(25)

	return codec.NewCodec(secretKey, jwt.SigningMethodHS256)
}
