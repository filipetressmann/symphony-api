package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	birthDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	u := NewUser(0, "johndoe", "John Doe", "john@example.com", birthDate, "123456789")

	assert.NotNil(t, u)
	assert.Equal(t, int64(0), u.UserId)
	assert.Equal(t, "johndoe", u.Username)
	assert.Equal(t, "John Doe", u.Fullname)
	assert.Equal(t, "john@example.com", u.Email)
	assert.Equal(t, birthDate, u.Birth_date)
	assert.Equal(t, "123456789", u.Telephone)
	assert.WithinDuration(t, time.Now(), u.Register_date, time.Second)
}

func TestUserToMap(t *testing.T) {
	birthDate := time.Date(1995, 6, 15, 0, 0, 0, 0, time.UTC)
	u := NewUser(0, "alice", "Alice Smith", "alice@example.com", birthDate, "987654321")

	m := u.ToMap()

	assert.Equal(t, u.UserId, m["id"])
	assert.Equal(t, u.Username, m["username"])
	assert.Equal(t, u.Fullname, m["fullname"])
	assert.Equal(t, u.Email, m["email"])
	assert.Equal(t, u.Register_date, m["register_date"])
	assert.Equal(t, u.Birth_date, m["birth_date"])
	assert.Equal(t, u.Telephone, m["telephone"])
}