package user

import (
	"fmt"
	"regexp"
	"userservice/common"
)

type UserService struct {
	regexExpression *regexp.Regexp
	repo            Repository
}

func NewUserService(repo Repository) *UserService {
	expression := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return &UserService{
		regexExpression: expression,
		repo:            repo,
	}
}

func (u *UserService) UserByEmail(email string) (*User, error) {
	return u.repo.UserByEmail(email)
}

func (u *UserService) UserById(id string) (*User, error) {
	return u.repo.UserById(id)
}

func (u *UserService) Create(user *User) (string, error) {
	user.Active = false

	// First validate the values. Do they seem reasonable?
	if len(user.Givenname) == 0 {
		return "", common.NewUserError("Missing givenname", common.VALIDATION_ERROR_TYPE)
	}

	if len(user.Familyname) == 0 {
		return "", common.NewUserError("Missing familyname", common.VALIDATION_ERROR_TYPE)
	}

	if len(user.Email) == 0 {
		return "", common.NewUserError("Missing email", common.VALIDATION_ERROR_TYPE)
	}
	if !u.regexExpression.Match([]byte(user.Email)) {
		return "", common.NewUserError("Invalid email", common.VALIDATION_ERROR_TYPE)
	}

	// Check if we have used that email or username before
	if param, err := u.repo.UserExists(user.Email, user.Username); err != nil || len(param) > 0 {
		if err != nil {
			return "", common.NewUserError(err.Error(), common.VALIDATION_ERROR_TYPE)
		}
		return "", common.NewUserError(fmt.Sprintf("%s already exists", param), common.VALIDATION_ERROR_TYPE)
	}

	userId, err := u.repo.Create(user)

	if err != nil {
		return "", common.NewUserError(err.Error(), common.SQL_ERROR_TYPE)
	}

	return userId, nil
}

func (u *UserService) Update(user *User) error {
	return u.repo.Update(user)
}
