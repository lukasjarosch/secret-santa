package santa

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lukasjarosch/genki/types/nullable"

	"github.com/lukasjarosch/secret-santa/internal/models"
)

type UseCase struct {
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u UseCase) CreateGroup(ctx context.Context, name string, creatorName string, status models.GroupStatus) (models.Group, error) {
	//log := logger.WithMetadata(ctx)

	group := models.Group{
		GroupID:     uuid.New().String(),
		Name:        name,
		CreatorName: creatorName,
		Members:     nil,
		Status:      status,
		CreatedAt:   nullable.NewTimeFromUnix(time.Now().Unix()),
		UpdatedAt:   nullable.NewTimeFromUnix(time.Now().Unix()),
	}

	// TODO: validate

	return group, nil
}
