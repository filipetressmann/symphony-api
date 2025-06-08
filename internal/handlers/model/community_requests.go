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

type GetCommunityByNameRequest struct {
	CommunityName string `json:"community_name"`
}

type AddUserToCommunityRequest struct {
	CommunityName string `json:"community_name"`
	Username string `json:"username"`
}

type ListUsersOfCommunityRequest struct {
	CommunityName string `json:"community_name"`
	Username string `json:"username"`
}

type ListUsersOfCommunityResponse struct {
	Users []*UserResponse `json:"users"`
}

type CommunityDataResponse struct {
	*BaseCommunityData
	CreatedAt time.Time `json:"created_at"`
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