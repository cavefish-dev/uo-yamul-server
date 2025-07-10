package mocks

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"yamul-gateway/internal/dtos"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/connection"
)

type ClientConnectionWriteBufferMock struct {
	assert                     *assert.Assertions
	mutexIsLocked              bool
	mutexAlreadyUsed           bool
	buffer                     []byte
	usedBufferLength           int
	logger                     interfaces.Logger
	loginDetails               *dtos.LoginDetails
	createGameConnectionResult error
	gameConnectionCreated      int
	mockedGameService          interfaces.GameService
}

func (c *ClientConnectionWriteBufferMock) Close() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) SendAnyData() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) ReceiveData() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) ProcessInputBuffer() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) ReadByte() byte {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) WriteByte(value byte) {
	c.assert.True(c.mutexIsLocked, "Mutex not locked")
	c.assert.Less(c.usedBufferLength, 1024, "Buffer overflow. Message too large.")
	c.buffer[c.usedBufferLength] = value
	c.usedBufferLength++
}

func (c *ClientConnectionWriteBufferMock) ReadUShort() uint16 {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) WriteUShort(value uint16) {
	c.WriteByte(byte(value >> 8))
	c.WriteByte(byte(value))
}

func (c *ClientConnectionWriteBufferMock) ReadUInt() uint32 {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) WriteUInt(value uint32) {
	c.WriteByte(byte(value >> 24))
	c.WriteByte(byte(value >> 16))
	c.WriteByte(byte(value >> 8))
	c.WriteByte(byte(value))
}

func (c *ClientConnectionWriteBufferMock) ReadFixedString(length int) string {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) ReadFixedBytes(length int) []byte {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) WriteFixedString(length int, value string) {
	for i := 0; i < length; i++ {
		if i < len(value) {
			c.WriteByte(value[i])
		} else {
			c.WriteByte(0)
		}
	}
}

func (c *ClientConnectionWriteBufferMock) UpdateEncryptionSeed(newSeed uint32) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) StartPacket() {
	c.assert.False(c.mutexIsLocked, "Mutex is already Locked")
	c.assert.False(c.mutexAlreadyUsed, "Mutex was already Locked")
	c.mutexIsLocked = true
}

func (c *ClientConnectionWriteBufferMock) EndPacket() {
	c.assert.True(c.mutexIsLocked, "Mutex is not Locked")
	c.mutexIsLocked = false
	c.mutexAlreadyUsed = true
}

func (c *ClientConnectionWriteBufferMock) CheckEncryptionHandshake() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) SetLogin(username string, password string) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) CreateGameConnection() error {
	c.gameConnectionCreated++
	return c.createGameConnectionResult
}

func (c *ClientConnectionWriteBufferMock) KillConnection(err error) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) IsConnectionHealthy() bool {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) GetLogger() interfaces.Logger {
	return c.logger
}

func (c *ClientConnectionWriteBufferMock) GetStatus() *dtos.ClientConnectionStatus {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) GetEncryptionState() *dtos.EncryptionConfig {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) GetLoginDetails() *dtos.LoginDetails {
	if c.loginDetails == nil {
		panic("Unimplemented mock behaviour")
	}
	return c.loginDetails
}

func (c *ClientConnectionWriteBufferMock) GetConnection() net.Conn {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionWriteBufferMock) GetGameService() interfaces.GameService {
	if c.mockedGameService == nil {
		panic("Unimplemented mock behaviour")
	}
	return c.mockedGameService
}

func (c *ClientConnectionWriteBufferMock) AssertSentBuffer(expected []byte) {
	c.assert.False(c.mutexIsLocked, "Mutex is Locked")
	c.assert.EqualValues(expected, c.buffer[0:c.usedBufferLength])
}

func (c *ClientConnectionWriteBufferMock) AssertDeclaredLength(lengthPosition uint16) {
	var expectedLength = uint16(c.buffer[lengthPosition])<<8 | uint16(c.buffer[lengthPosition+1])
	c.assert.EqualValues(expectedLength, uint16(c.usedBufferLength))
}

func (c *ClientConnectionWriteBufferMock) MockLoginDetails(loginDetails *dtos.LoginDetails) {
	c.loginDetails = loginDetails
}

func (c *ClientConnectionWriteBufferMock) MockCreateGameConnection(e error) {
	c.createGameConnectionResult = e
}

func (c *ClientConnectionWriteBufferMock) AssertCreatedGameCollection(i int) {
	c.assert.Equal(i, c.gameConnectionCreated)
}

func (c *ClientConnectionWriteBufferMock) MockGameService(mockGameService interfaces.GameService) {
	c.mockedGameService = mockGameService
}

func CreateClientConnectionWriteBufferMock(t *testing.T) *ClientConnectionWriteBufferMock {
	result := &ClientConnectionWriteBufferMock{
		assert:           assert.New(t),
		mutexIsLocked:    false,
		mutexAlreadyUsed: false,
		buffer:           make([]byte, 1024),
		usedBufferLength: 0,
		logger:           connection.CreateAnonymousLogger(t.Name()),
	}
	var _ interfaces.ClientConnection = result
	return result
}
