package usecases

import (
	"context"
	"github.com/liguangsheng/codelab/clean-architecture-demo/entities"
	"github.com/liguangsheng/codelab/clean-architecture-demo/usecases/ports"
)

type UserService struct {
	UserRepository ports.IUserRepository
}

func (s *UserService) GetUser(ctx context.Context, userID uint64) (*entities.User, error) {
	return s.UserRepository.GetUserByID(ctx, userID)
}
