//go:build integration

package integration

import (
	"net"
	"sync"
	"testing"

	"yamul-gateway/internal/autoconfig"
	"yamul-gateway/internal/events"
	"yamul-gateway/internal/transport/multima"
	"yamul-gateway/internal/transport/multima/handlers"
)

var handlersOnce sync.Once

type gatewayHarness struct {
	client  *uoClient
	backend *mockBackend
}

func newGatewayHarness(t *testing.T) *gatewayHarness {
	t.Helper()

	backend := startMockBackend(t)

	t.Setenv("YAMUL_LOGIN_ADDR", backend.loginAddr)
	t.Setenv("YAMUL_CHARACTER_ADDR", backend.characterAddr)
	t.Setenv("YAMUL_GAME_ADDR", backend.gameAddr)

	handlersOnce.Do(func() {
		handlers.Setup()
		events.Setup()
	})

	if err := autoconfig.ResetLoginModule(); err != nil {
		t.Fatalf("ResetLoginModule: %v", err)
	}
	t.Cleanup(autoconfig.CloseLoginModule)

	serverConn, clientConn := net.Pipe()
	t.Cleanup(func() {
		_ = serverConn.Close()
		_ = clientConn.Close()
	})

	go multima.ClientConnectionLoop(serverConn)

	return &gatewayHarness{
		client:  newUOClient(clientConn),
		backend: backend,
	}
}
