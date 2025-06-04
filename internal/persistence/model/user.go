package model

import (
	"encoding/json"
	"net/http"
	"time"
)

// User represents a user in the system.
// It contains fields for user ID, username, full name, email, registration date,
// birth date, and telephone number.
// The User struct is designed to be used with JSON serialization and deserialization.
// The `json` tags specify how the fields should be serialized to and from JSON.
// The `Register_date` field is automatically set to the current time when a new User is created.
// The `Birth_date` field is expected to be provided in the format of a time.Time object.
// The `UserId` field is an int64 that uniquely identifies the user.
type User struct {
	UserId int64
	Username string `json:"username"`
	Fullname string  `json:"full_name"`
	Email string `json:"email"`
	Register_date time.Time
	Birth_date time.Time `json:"birthdate"`
	Telephone string `json:"telephone"`
}

// NewUser creates a new User instance with the provided parameters.
// It initializes the UserId, Username, Fullname, Email, Birth_date, and Telephone fields.
// The Register_date field is set to the current time.
// The function returns a pointer to the newly created User instance.
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

// ToMap converts the User instance to a map[string]any.
// It returns a map where the keys are the field names and the values are the corresponding field values.
// This method is useful for serializing the User instance to a format that can be easily converted to JSON or used in other contexts.
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

// UserFromRequest decodes a JSON request body into a User instance.
// It reads the request body, unmarshals the JSON data into a User struct,
// and returns a pointer to the User instance along with any error encountered during the decoding process.
func UserFromRequest(request *http.Request) (*User, error) {
	var user User
    err := json.NewDecoder(request.Body).Decode(&user)
    return &user, err
}

// MapToUser converts a map[string]any to a User instance.
// It expects the map to contain keys corresponding to the User fields.
// The function retrieves the values from the map, casts them to the appropriate types,
// and returns a pointer to a new User instance.
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
