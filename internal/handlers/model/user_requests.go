package request_model

import (
	"symphony-api/internal/persistence/model"
	"time"
)

type GetUserByUsernameRequest struct {
	Username string `schema:"username,required"`
}

type GetUserFriendsRequest struct {
	Username string `schema:"username,required"`
}

type GetUserFriendsResponse struct {
	Friends []*UserResponse `json:"friends" binding:"required"`
}

type ListUserCommunitiesRequest struct {
	Username string `schema:"username,required"`
}

type ListUserCommunitiesResponse struct {
	Communities []*CommunityDataResponse `json:"communities" binding:"required"`
}

type BaseUserModel struct {
	Username string `json:"username" binding:"required"`
	Fullname string  `json:"fullname" binding:"required"`
	Email string `json:"email" binding:"required"`
	Birth_date time.Time `json:"birth_date" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

type CreateUserRequest struct {
	*BaseUserModel
}

type CreateFriendshipRequest struct {
	Username1 string `json:"username1" binding:"required"`
	Username2 string `json:"username2" binding:"required"`
}

func NewBaseUserModel(user *model.User) *BaseUserModel {
	return &BaseUserModel {
		Username: user.Username,
		Fullname: user.Fullname,
		Email: user.Email,
		Birth_date: user.Birth_date,
		Telephone: user.Telephone,
	}
}

func (request *CreateUserRequest) ToUser() *model.User {
	return &model.User{
		Username: request.Username,
		Fullname: request.Fullname,
		Email: request.Email,
		Birth_date: request.Birth_date,
		Telephone: request.Telephone,
	}
}

type UserResponse struct {
	*BaseUserModel
	Id int32 `json:"id"`
	Register_date time.Time `json:"register_date"`
}

func NewUserResponse(user *model.User) *UserResponse {
	return &UserResponse {
		Id: user.UserId,
		Register_date: user.Register_date,
		BaseUserModel: NewBaseUserModel(user),
	}
}
