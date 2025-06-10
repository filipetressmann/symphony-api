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
	communityName string,
	description string,
	createdAt time.Time,
) *Community {
	return &Community{
		CommunityName: communityName,
		Description: description,
		CreatedAt: createdAt,
	}
}

func MapArrayToCommunity(data []map[string]any) []*Community {
	communities := make([]*Community, 0)

	for _, community := range data {
		communities = append(communities, NewCommunityFromMap(community))
	}

	return communities
}
