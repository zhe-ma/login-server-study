package util

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/cryto/bcrypt"
)

func GetRequestId(c *gin.Context) string {
	value, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}

	if requestId, ok := v.(string); ok {
		return requestId
	}

	return ""
}

func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
