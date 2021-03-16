package services

import (
	domain "github.com/Fernando-liu-Melbourne/microservices/domain"
	"github.com/Fernando-liu-Melbourne/microservices/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}