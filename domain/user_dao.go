package domain

import (
	"errors"
	"fmt"
)

var (
	user = map[uint64]*User {
		123: &User{uint64(123), "Fernando", "Liu"},
	}
)

func GetUser(userID uint64) (*User, error) {
	if user := user[userID]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User ID %v does not exist", userID))
}