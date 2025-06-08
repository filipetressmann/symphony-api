package request_model

import (
	"symphony-api/internal/persistence/model"
	"time"
)

type BaseCommunityData struct {
	CommunityName string `json:"community_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateCommunityRequest struct {
	*BaseCommunityData
}

type GetCommunityByNameRequest struct {
	CommunityName string `json:"community_name" binding:"required"`
}

type AddUserToCommunityRequest struct {
	CommunityName string `json:"community_name" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type ListUsersOfCommunityRequest struct {
	CommunityName string `json:"community_name" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type ListUsersOfCommunityResponse struct {
	Users []*UserResponse `json:"users" binding:"required"`
}

type CommunityDataResponse struct {
	*BaseCommunityData
	CreatedAt time.Time `json:"created_at" binding:"required"`
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