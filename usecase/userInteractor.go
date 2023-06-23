package usecase

import (
	"echo-clean-architecture/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)

	return
}

func (interactor *UserInteractor) UserByEmail(email string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmail(email)

	return
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)

	return
}

func (interactor *UserInteractor) DeleteByUserKey(u domain.User) (err error) {
	err = interactor.UserRepository.DeleteByUserKey(u)
	
	return
}
