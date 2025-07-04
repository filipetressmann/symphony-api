package model

import (
	"time"
)

type User struct {
	UserId int32
	Username string
	Fullname string
	Email string
	Register_date time.Time
	Birth_date time.Time
	Telephone string
}

func NewUser(
	username string,
	fullname string,
	email string,
	birthdate time.Time,
	telephone string,
) *User {
	return &User{
		Username: username,
		Fullname: fullname,
		Email: email,
		Birth_date: birthdate,
		Register_date: time.Now(),
		Telephone: telephone,
	}
}

func (user *User) ToMap() map[string]any {
	return map[string]any{
		"username": user.Username,
		"fullname": user.Fullname,
		"email": user.Email,
		"birth_date": user.Birth_date,
		"telephone": user.Telephone,
	}
}

func MapToUser(data map[string]any) *User {
	return &User{
		UserId: data["id"].(int32),
		Username: data["username"].(string),
		Fullname: data["fullname"].(string),
		Email: data["email"].(string),
		Register_date: data["register_date"].(time.Time),
		Birth_date: data["birth_date"].(time.Time),
		Telephone: data["telephone"].(string),
	}
}

func MapArrayToUsers(data []map[string]any) []*User {
	users := make([]*User, 0)

	for _, user := range data {
		users = append(users, MapToUser(user))
	}

	return users
}
