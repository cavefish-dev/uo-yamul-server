package handlers

import (
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func handler_openChatWindow(client interfaces.ClientConnection) { // 0x22
	chatWindowName := client.ReadFixedString(63)
	listeners.OnOpenChatWindow.Trigger(client, chatWindowName)
}
