package auth

import (
	"context"
	"coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

//Service implement auth service
type Service struct {
	OpenIdResolver OpenIdResolver
	Mongo          *dao.Mongo
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Logger         *zap.Logger
	authpb.UnimplementedAuthServiceServer
}
type OpenIdResolver interface {
	Resolve(code string) (string, error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
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
	tkn, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
