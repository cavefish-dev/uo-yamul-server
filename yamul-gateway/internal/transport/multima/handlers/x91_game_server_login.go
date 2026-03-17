package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
	"yamul-gateway/utils/stringUtils"
)

func gameServerLogin(client interfaces.ClientConnection) { // 0xA0
	encriptionKey := client.ReadUInt()
	username := stringUtils.TrimRight(client.ReadFixedString(30))
	password := stringUtils.TrimRight(client.ReadFixedString(30))
	command := commands.GameLoginRequest{
		Username:      username,
		Password:      password,
		EncryptionKey: encriptionKey,
	}

	listeners.OnGameLoginRequest.Trigger(client, command)

}
