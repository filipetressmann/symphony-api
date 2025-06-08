package request_model

import (
	"symphony-api/internal/persistence/model"
	"time"
)

type GetUserByUsernameRequest struct {
	Username string `json:"username"`
}

type BaseUserModel struct {
	Username string `json:"username"`
	Fullname string  `json:"fullname"`
	Email string `json:"email"`
	Birth_date time.Time `json:"birth_date"`
	Telephone string `json:"telephone"`
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

type CreateUserRequest struct {
	*BaseUserModel
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
