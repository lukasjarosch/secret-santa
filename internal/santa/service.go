package santa

import (
	"context"

	"github.com/lukasjarosch/secret-santa/internal/models"
)

type Service interface {
	CreateGroup(ctx context.Context, name string, creatorName string, status models.GroupStatus) (models.Group, error)
}