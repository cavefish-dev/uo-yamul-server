//go:build integration

package integration

import (
	"context"
	"fmt"
	"net"
	"strings"
	"testing"

	"google.golang.org/grpc"
	services "yamul-gateway/backend/services"
)

func trimNullBytes(s string) string {
	return strings.TrimRight(s, "\x00")
}

type mockLoginServer struct {
	services.UnimplementedLoginServiceServer
	credentials map[string]string
}

func (m *mockLoginServer) ValidateLogin(_ context.Context, req *services.LoginRequest) (*services.LoginResponse, error) {
	pass, ok := m.credentials[req.Username]
	if ok && trimNullBytes(pass) == trimNullBytes(req.Password) {
		return &services.LoginResponse{Value: services.LoginResponse_valid}, nil
	}
	return &services.LoginResponse{Value: services.LoginResponse_invalid}, nil
}

func (m *mockLoginServer) AddCredential(username, password string) {
	m.credentials[username] = password
}

type mockCharacterServer struct {
	services.UnimplementedCharacterServiceServer
}

func (m *mockCharacterServer) GetCharacterList(_ context.Context, _ *services.Empty) (*services.CharacterListResponse, error) {
	return &services.CharacterListResponse{
		CharacterLogins: []*services.LoginRequest{
			{Username: "testchar", Password: ""},
		},
	}, nil
}

type mockGameServer struct {
	services.UnimplementedGameServiceServer
	Received chan *services.StreamPackage
	ToSend   chan *services.StreamPackage
}

func (m *mockGameServer) OpenGameStream(stream services.GameService_OpenGameStreamServer) error {
	errCh := make(chan error, 1)
	go func() {
		for {
			pkg, err := stream.Recv()
			if err != nil {
				errCh <- err
				return
			}
			select {
			case m.Received <- pkg:
			default:
				errCh <- fmt.Errorf("mockGameServer: Received channel is full, cannot forward incoming package")
				return
			}
		}
	}()
	for {
		select {
		case pkg := <-m.ToSend:
			if err := stream.Send(pkg); err != nil {
				return err
			}
		case err := <-errCh:
			return err
		}
	}
}

type mockBackend struct {
	Login         *mockLoginServer
	Character     *mockCharacterServer
	Game          *mockGameServer
	loginAddr     string
	characterAddr string
	gameAddr      string
}

func startMockBackend(t *testing.T) *mockBackend {
	t.Helper()

	backend := &mockBackend{
		Login:     &mockLoginServer{credentials: make(map[string]string)},
		Character: &mockCharacterServer{},
		Game: &mockGameServer{
			Received: make(chan *services.StreamPackage, 16),
			ToSend:   make(chan *services.StreamPackage, 16),
		},
	}

	backend.loginAddr = startGRPCServer(t, func(s *grpc.Server) {
		services.RegisterLoginServiceServer(s, backend.Login)
	})
	backend.characterAddr = startGRPCServer(t, func(s *grpc.Server) {
		services.RegisterCharacterServiceServer(s, backend.Character)
	})
	backend.gameAddr = startGRPCServer(t, func(s *grpc.Server) {
		services.RegisterGameServiceServer(s, backend.Game)
	})

	return backend
}

func startGRPCServer(t *testing.T, register func(*grpc.Server)) string {
	t.Helper()
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen: %v", err)
	}
	s := grpc.NewServer()
	register(s)
	go func() { _ = s.Serve(lis) }()
	t.Cleanup(s.GracefulStop)
	return lis.Addr().String()
}
