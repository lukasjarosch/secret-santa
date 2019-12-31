package proto

import (
	"github.com/lukasjarosch/genki/types/nullable"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/lukasjarosch/secret-santa/api/v1"
	"github.com/lukasjarosch/secret-santa/internal/models"
)

func EncodeError(err error) error {
	switch err {
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}

func EncodeGroup(group models.Group) *v1.Group {
	return &v1.Group{
		Id:                   group.GroupID,
		Name:                 group.Name,
		CreatorName:          group.CreatorName,
		Members:              EncodeGroupMemberList(group.Members),
		CreatedAt:            NullableTimeToProtoDate(group.CreatedAt),
	}
}

func EncodeGroupMemberList(member []models.GroupMember) []*v1.GroupMember {
	var members []*v1.GroupMember

	for _, mem := range member {
		members = append(members, EncodeGroupMember(mem))
	}

	return members
}

func EncodeGroupMember(member models.GroupMember) *v1.GroupMember {
	return &v1.GroupMember{
		Id:                   member.MemberID,
		Name:                 member.Name,
	}
}

func NullableTimeToProtoDate(time nullable.Time) *date.Date {
	if !time.Valid {
		return &date.Date{}
	}

	return &date.Date{
		Year:                 int32(time.Time.Year()),
		Month:                int32(time.Time.Month()),
		Day:                  int32(time.Time.Day()),
	}
}