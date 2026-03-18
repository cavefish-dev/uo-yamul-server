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

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen: %v", err)
	}
	t.Cleanup(func() { _ = listener.Close() })

	var serverConn net.Conn
	connCh := make(chan net.Conn, 1)
	go func() {
		c, err := listener.Accept()
		if err == nil {
			connCh <- c
		}
	}()

	clientConn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("net.Dial: %v", err)
	}
	serverConn = <-connCh

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
