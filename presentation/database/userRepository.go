package database

import (
	"echo-clean-architecture/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}

	return
}

func (repo *UserRepository) FindByEmail(email string) (user domain.User, err error) {
	if err = repo.Where("email = ?", email).Find(&user).Error; err != nil {
		return
	}

	return
}

func (repo *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = repo.Create(&u).Error; err != nil { return }
	user = u

	return
}

func (repo *UserRepository) DeleteByUserKey(user domain.User) (err error) {
	if err = repo.Where("user_key = ?", user.UserKey).Delete(&user).Error; err != nil {
		return
	}
	
	return
}
