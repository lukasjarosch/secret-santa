package proto

import (
	"context"

	. "github.com/lukasjarosch/secret-santa/api/v1"
	"github.com/lukasjarosch/secret-santa/internal/santa"
)

type SecretSantaService struct {
	useCase santa.Service
}

func NewSecretSantaService(useCase santa.Service) *SecretSantaService {
	return &SecretSantaService{useCase:useCase}
}

func (svc *SecretSantaService) CreateGroup(ctx context.Context, in *CreateGroupRequest) (*CreateGroupResponse, error) {
	req := DecodeCreateGroupRequest(in)
	group, err := svc.useCase.CreateGroup(ctx, req.Name, req.CreatorName, req.Status)
	if err != nil {
		return nil, EncodeError(err)
	}
	return &CreateGroupResponse{
		Group: EncodeGroup(group),
	}, nil
}

func (svc *SecretSantaService) GetGroup(ctx context.Context, in *GetGroupRequest) (*GetGroupResponse, error) {
	return &GetGroupResponse{}, nil
}

func (svc *SecretSantaService) AddGroupMember(ctx context.Context, in *AddGroupMemberRequest) (*AddGroupMemberResponse, error) {
	return &AddGroupMemberResponse{}, nil
}

func (svc *SecretSantaService) SendInviteEmails(ctx context.Context, in *SendInviteEmailsRequest) (*SendInviteEmailsResponse, error) {
	return &SendInviteEmailsResponse{}, nil
}

func (svc *SecretSantaService) SaveWishes(ctx context.Context, in *SaveWishesRequest) (*SaveWishesResponse, error) {
	return &SaveWishesResponse{}, nil
}

func (svc *SecretSantaService) CreateMapping(ctx context.Context, in *CreateMappingRequest) (*CreateMappingResponse, error) {
	return &CreateMappingResponse{}, nil
}

func (svc *SecretSantaService) SendSantaEmails(ctx context.Context, in *SendSantaEmailsRequest) (*SendSantaEmailsResponse, error) {
	return &SendSantaEmailsResponse{}, nil
}
