package ports

import (
	"context"
	"github.com/liguangsheng/codelab/clean-architecture-demo/entities"
)

type IUserRepository interface {
	GetUser(ctx context.Context, userID uint64) (*entities.User, error)
}