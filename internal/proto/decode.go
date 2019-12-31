package proto

import (
	v1 "github.com/lukasjarosch/secret-santa/api/v1"
	"github.com/lukasjarosch/secret-santa/internal/models"
)

func DecodeCreateGroupRequest(pbReq *v1.CreateGroupRequest) *models.Group {
	if pbReq == nil {
		return &models.Group{}
	}

	return &models.Group{
		Name:        pbReq.Name,
		CreatorName: pbReq.CreatorName,
		Status:      models.UndefinedStatus,
	}
}
