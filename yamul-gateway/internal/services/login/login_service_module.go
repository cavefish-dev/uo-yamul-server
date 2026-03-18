package login

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"testing"
	backendServices "yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/services/login/mocks"
)

type LoginService struct {
	dial   *grpc.ClientConn
	client backendServices.LoginServiceClient
}

var service *LoginService
var serviceInterface LoginServiceInterface

var Module interfaces.Module = module{}

type module struct{}

func (m module) Setup() error {
	var err error
	service, err = newLoginService()
	serviceInterface = service
	return err
}

func MockSetup(t *testing.T) *mocks.LoginServiceMock {
	mock := mocks.CreateLoginServiceMock(t)
	serviceInterface = mock
	return mock
}

func MockClose() {
	serviceInterface = nil
}

func (m module) Close() {
	if service != nil {
		service.close()
	}
}

func loginServiceAddress() string {
	if addr := os.Getenv("YAMUL_LOGIN_ADDR"); addr != "" {
		return addr
	}
	return "localhost:8087"
}

func newLoginService() (*LoginService, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial(loginServiceAddress(), opts...)
	if err != nil {
		return nil, err
	}
	service := &LoginService{
		dial:   dial,
		client: backendServices.NewLoginServiceClient(dial),
	}
	var _ LoginServiceInterface = service
	return service, nil
}

func (s LoginService) close() {
	_ = s.dial.Close()
}

func (s LoginService) CheckUserCredentials(username string, password string) (bool, commands.LoginDeniedReason) {
	response, err := s.client.ValidateLogin(context.Background(), &backendServices.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return false, commands.LoginDeniedReason_CommunicationProblem
	}
	if response.Value == backendServices.LoginResponse_valid {
		return true, 0
	}
	return false, commands.LoginDeniedReason_IncorrectUsernamePassword
}
