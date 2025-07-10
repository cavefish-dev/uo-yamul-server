package clientEvents

import (
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces/mocks"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/listeners"
)

func TestOnCharacterPreLogin(t *testing.T) {
	type args struct {
		event          listeners.CommandEvent[commands.PreLogin]
		expectedBuffer []byte
		loginSuccess   bool
		loginResult    commands.LoginDeniedReason
	}

	const USERNAME = "name"
	const PASSWORD = "password"

	tests := []struct {
		name string
		args args
	}{
		{
			name: "base case",
			args: args{
				event: listeners.CommandEvent[commands.PreLogin]{
					Client: nil, // The client will be mocked later
					Command: commands.PreLogin{
						Name:          USERNAME,
						Password:      PASSWORD,
						Slot:          0,
						EncryptionKey: 0,
					},
				},
				loginSuccess:   true,
				loginResult:    0,
				expectedBuffer: []byte{}, // Replace with the expected buffer
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange: Setup login service mock and expectations
			var mockLogin = login.MockSetup(t)
			mockLogin.MockCheckUserCredentials(USERNAME, PASSWORD, tt.args.loginSuccess, tt.args.loginResult)
			defer login.MockClose()

			// Arrange: Create a mock client connection and set login details
			mockClient := mocks.CreateClientConnectionWriteBufferMock(t)
			mockClient.MockLoginDetails(&dtos.LoginDetails{
				Username: USERNAME,
				Password: PASSWORD,
			})

			// Arrange: Create a mock game service
			gameServiceMock := mocks.CreateGameServiceMock(t)

			// Expect: If login is successful, expect character selection message to be sent
			if tt.args.loginSuccess {
				body := services.MsgCharacterSelection{
					Username: USERNAME,
					Slot:     int32(tt.args.event.Command.Slot),
				}
				selection := services.Message_CharacterSelection{CharacterSelection: &body}
				gameServiceMock.MockSend(services.MsgType_TypeCharacterSelection, &services.Message{Msg: &selection})
			}

			// Arrange: Configure client to use game service and create game connection
			mockClient.MockGameService(gameServiceMock)
			mockClient.MockCreateGameConnection(nil)

			// Act: Assign mock client to event and execute pre-login event
			tt.args.event.Client = mockClient
			OnCharacterPreLogin(tt.args.event)

			// Assert: Check that the expected buffer was sent and game collection was created
			mockClient.AssertSentBuffer(tt.args.expectedBuffer)
			mockClient.AssertCreatedGameCollection(1)
			gameServiceMock.AssertAllSend()
		})
	}
}
