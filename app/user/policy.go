package user

import "gin-api/model"

type Policy struct{}

func (policy *Policy) ViewAny(user model.User) bool {
	return true
}

func (policy *Policy) View(user model.User, model model.User) bool {
	return true
}

func (policy *Policy) Create(user model.User) bool {
	return true
}

func (policy *Policy) Update(user model.User, model model.User) bool {
	return true
}

func (policy *Policy) Delete(user model.User, model model.User) bool {
	return true
}
