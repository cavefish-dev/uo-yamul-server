package mocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/interfaces"
)

type MockedGameService struct {
	assert               *assert.Assertions
	isClosed             bool
	expectedArguments    []sendArguments
	expectedArgumentsIdx int
	expectedArgumentsLen int
}

type sendArguments struct {
	_type   services.MsgType
	message *services.Message
}

func (m *MockedGameService) Close() {
	m.assert.False(m.isClosed, "MockedGameService already closed")
	m.isClosed = true
}

func (m *MockedGameService) Send(_type services.MsgType, message *services.Message) {
	m.assert.Lessf(m.expectedArgumentsIdx, m.expectedArgumentsLen, "Unexpected call to Send. Expected %d calls, got %d", len(m.expectedArguments), m.expectedArgumentsIdx+1)
	m.assert.NotNil(m.expectedArguments[m.expectedArgumentsIdx], "Unexpected call to Send. Expected %d calls, got %d", len(m.expectedArguments), m.expectedArgumentsIdx+1)
	m.assert.Equalf(m.expectedArguments[m.expectedArgumentsIdx]._type, _type, "Unexpected message type. Expected %s, got %s", m.expectedArguments[m.expectedArgumentsIdx]._type, _type)
	m.assert.Equalf(m.expectedArguments[m.expectedArgumentsIdx].message, message, "Unexpected message. Expected %v, got %v", m.expectedArguments[m.expectedArgumentsIdx].message, message)
	m.expectedArgumentsIdx++
}

func (m *MockedGameService) MockSend(_type services.MsgType, message *services.Message) {
	m.expectedArguments[m.expectedArgumentsLen] = sendArguments{
		_type:   _type,
		message: message,
	}
	m.expectedArgumentsLen++
}

func (m *MockedGameService) AssertAllSend() {
	m.assert.Equal(m.expectedArgumentsIdx, m.expectedArgumentsLen, "Not all expected Send calls were made. Expected %d, got %d", m.expectedArgumentsLen, m.expectedArgumentsIdx)
}

func CreateGameServiceMock(t *testing.T) *MockedGameService {
	result := &MockedGameService{
		assert:               assert.New(t),
		isClosed:             false,
		expectedArgumentsIdx: 0,
		expectedArgumentsLen: 0,
		expectedArguments:    make([]sendArguments, 128),
	}
	var _ interfaces.GameService = result // Ensure MockedGameService implements GameService interface
	return result
}
