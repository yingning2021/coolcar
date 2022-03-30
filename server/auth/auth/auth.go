package auth

import (
	"context"
	"coolcar/auth/api/gen/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Service implement auth service
type Service struct {
	OpenIdResolver OpenIdResolver
	authpb.UnimplementedAuthServiceServer
	Logger *zap.Logger
}
type OpenIdResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openID, err := s.OpenIdResolver.Resolve(request.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}
	s.Logger.Info("received code", zap.String("code", request.Code))
	return &authpb.LoginResponse{
		AccessToken: "token for open id" + openID,
		ExpiresIn:   7200,
	}, nil
}
