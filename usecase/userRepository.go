package usecase

import (
    "echo-clean-architecture/domain"
)

type UserRepository interface {
	FindById(id int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Store(domain.User) (domain.User, error)
	DeleteByUserKey(domain.User) error
}
