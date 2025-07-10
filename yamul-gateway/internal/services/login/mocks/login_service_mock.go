package mocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"yamul-gateway/internal/dtos/commands"
)

type LoginServiceMock struct {
	assert           *assert.Assertions
	expectedUsername string
	expectedPassword string
	resultBool       bool
	resultReason     commands.LoginDeniedReason
}

func (l *LoginServiceMock) CheckUserCredentials(username string, password string) (bool, commands.LoginDeniedReason) {
	l.assert.Equal(l.expectedUsername, username)
	l.assert.Equal(l.expectedPassword, password)
	return l.resultBool, l.resultReason
}

func (l *LoginServiceMock) MockCheckUserCredentials(username string, password string, result1 bool, result2 commands.LoginDeniedReason) {
	l.expectedUsername = username
	l.expectedPassword = password
	l.resultBool = result1
	l.resultReason = result2
}

func CreateLoginServiceMock(t *testing.T) *LoginServiceMock {
	return &LoginServiceMock{
		assert: assert.New(t),
	}
}
