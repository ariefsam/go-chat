package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"
)

func TestGenerateToken(t *testing.T) {
	var u usecase.Usecase
	var mockTokenUser mockdependency.TokenUserService
	u.TokenUserService = &mockTokenUser
	user := entity.User{
		ID:          "uid",
		PhoneNumber: "625363",
	}

	expectedToken := "xxjjjjjsss"
	mockTokenUser.On("Create", user).Return(expectedToken)
	token := u.GenerateToken(user)
	assert.Equal(t, expectedToken, token)
}
