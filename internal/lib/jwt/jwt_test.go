package jwt

import (
	"app/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	user := models.User{
		ID:    123,
		Email: "test@gmail.com"}
	app := models.App{
		ID:     123,
		Name:   "service",
		Secret: "superSecret",
	}
	duration := time.Minute * 60

	token, err := NewToken(user, app, duration)

	assert.NoError(t, err, "failed to create token")
	assert.NotEmpty(t, token, "token is not supposed to be empty")

	t.Log("token:", token)

	parsedToken, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				t.Fatalf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(app.Secret), nil
		},
	)

	assert.NoError(t, err, "failed to parse token")
	assert.True(t, parsedToken.Valid, "token must be valid")

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	require.True(t, ok)

	assert.Equal(t, user.ID, int64(claims["uid"].(float64)))
	assert.Equal(t, user.Email, (claims["email"]).(string))
	assert.Equal(t, app.ID, int(claims["app_id"].(float64)))
	t.Log("parsedToken", parsedToken)

}
