package model

import (
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	UserId int64
	Username string `json:"username"`
	Fullname string  `json:"full_name"`
	Email string `json:"email"`
	Register_date time.Time
	Birth_date time.Time `json:"birthdate"`
	Telephone string `json:"telephone"`
}

func NewUser(
	userId int64,
	username string,
	fullname string,
	email string,
	birthdate time.Time,
	telephone string,
) *User {
	return &User{
		UserId: userId,
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
		"id": user.UserId,
		"username": user.Username,
		"fullname": user.Fullname,
		"email": user.Email,
		"register_date": user.Register_date,
		"birth_date": user.Birth_date,
		"telephone": user.Telephone,
	}
}

func UserFromRequest(request *http.Request) (*User, error) {
	var user User
    err := json.NewDecoder(request.Body).Decode(&user)
    return &user, err
}

func MapToUser(data map[string]any) *User {
	return &User{
		UserId: data["id"].(int64),
		Username: data["username"].(string),
		Fullname: data["fullname"].(string),
		Email: data["email"].(string),
		Register_date: data["register_date"].(time.Time),
		Birth_date: data["birth_date"].(time.Time),
		Telephone: data["telephone"].(string),
	}
}
