package request_model

import (
	"symphony-api/internal/persistence/model"
	"time"
)

type BaseCommunityData struct {
	CommunityName string `json:"community_name"`
	Description string `json:"description"`
}

type CreateCommunityRequest struct {
	*BaseCommunityData
}

type CommunityDataResponse struct {
	*BaseCommunityData
	CreatedAt time.Time `json:"created_at"`
}

type GetCommunityByNameRequest struct {
	CommunityName string `json:"community_name"`
}

func (request *CreateCommunityRequest) ToCommunity() *model.Community {
	return &model.Community{
		CommunityName: request.CommunityName,
		Description: request.Description,
	}
}

func NewBaseCommunityData(communityName string, description string) *BaseCommunityData {
	return &BaseCommunityData {
		CommunityName: communityName,
		Description: description,
	}
}

func NewCommunityDataResponse(community *model.Community) *CommunityDataResponse {
	return &CommunityDataResponse{
		BaseCommunityData: NewBaseCommunityData(community.CommunityName, community.Description),
		CreatedAt: community.CreatedAt,
	}
}