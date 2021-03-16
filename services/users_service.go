package services

import (
	domain "github.com/Fernando-liu-Melbourne/microservices/domain"
)

func GetUser(userId uint64) (*domain.User, error) {
	return domain.GetUser(userId)
}