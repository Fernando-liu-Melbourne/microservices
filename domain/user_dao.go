package domain

import (
	"fmt"
	"github.com/Fernando-liu-Melbourne/microservices/utils"
	"net/http"
)

var (
	user = map[uint64]*User {
		123: &User{uint64(123), "Fernando", "Liu"},
	}
)

func GetUser(userID uint64) (*User, *utils.ApplicationError) {
	if user := user[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User ID %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code: "Not Found",
	}
}