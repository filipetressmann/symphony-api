package model

import "time"

type Community struct {
	Id int32
	CommunityName string
	Description string
	CreatedAt time.Time
}

func (community *Community) ToTableData() map[string]any {
	return map[string]any{
		"community_name": community.CommunityName,
		"description": community.Description,
	}
}

func NewCommunityFromMap(data map[string]any) *Community {
	return &Community{
		Id: data["id"].(int32),
		CommunityName: data["community_name"].(string),
		Description: data["description"].(string),
		CreatedAt: data["created_at"].(time.Time),
	}
}

func NewCommunity(
	id int32,
	communityName string,
	description string,
	createdAt time.Time,
) *Community {
	return &Community{
		Id: id,
		CommunityName: communityName,
		Description: description,
		CreatedAt: createdAt,
	}
}