package token_user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/token_user"
)

func TestToken(t *testing.T) {
	var user entity.User
	user.ID = "id01"
	user.Name = "name user"
	user.PhoneNumber = "625232323"
	var tokenService token_user.Token
	tokenService.Secret = []byte("xxx")
	token := tokenService.Create(user)
	assert.NotEmpty(t, token)
	isValid, parseUser := tokenService.Parse(token)
	assert.True(t, isValid)
	assert.Equal(t, user, parseUser)
}
