package auth

import (
	"context"
	"coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Service implement auth service
type Service struct {
	OpenIdResolver OpenIdResolver
	Mongo          *dao.Mongo
	Logger         *zap.Logger
	authpb.UnimplementedAuthServiceServer
}
type OpenIdResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(c context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openID, err := s.OpenIdResolver.Resolve(request.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}

	accountID, err := s.Mongo.ResolveAccountID(c, openID)
	if err != nil {
		s.Logger.Error("cannot resolve account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: "token for open id" + accountID,
		ExpiresIn:   7200,
	}, nil
}
