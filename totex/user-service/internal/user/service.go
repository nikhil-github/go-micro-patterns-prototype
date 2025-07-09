package user

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	userv1 "github.com/yourusername/schema/gen/user/v1"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(ctx context.Context, req *connect.Request[userv1.CreateUserRequest]) (*connect.Response[userv1.CreateUserResponse], error) {
	id := uuid.NewString()
	now := time.Now().UTC().Format(time.RFC3339)
	resp := &userv1.CreateUserResponse{
		UserId:    id,
		Email:     req.Msg.Email,
		Name:      req.Msg.Name,
		CreatedAt: now,
	}
	return connect.NewResponse(resp), nil
}
