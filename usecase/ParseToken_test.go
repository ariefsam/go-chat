package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"
)

func TestParseToken(t *testing.T) {
	var u usecase.Usecase
	var mockTokenUserService mockdependency.TokenUserService
	u.TokenUserService = &mockTokenUserService
	expectedUser := entity.User{
		ID:          "uid",
		PhoneNumber: "625363",
	}
	expectedIsValid := true
	token := "xxdddkkkkjjj"
	mockTokenUserService.On("Parse", token).Return(expectedIsValid, expectedUser)
	isValid, user := u.ParseToken(token)
	assert.Equal(t, expectedIsValid, isValid)
	assert.Equal(t, expectedUser, user)
}
