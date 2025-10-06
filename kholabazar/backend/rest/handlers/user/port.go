package user

import "kholabazar/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, password string) (*domain.User, error)
}
