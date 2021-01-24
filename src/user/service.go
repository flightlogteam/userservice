package user

import (
	"fmt"
	"regexp"
	"userservice/common"

	"golang.org/x/crypto/bcrypt"
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

func (u *UserService) Login(username string, email string, password string) (*User, error) {
	var user *User
	var err error
	if len(username) != 0 {
		user, err = u.repo.UserByUsername(username)
	}

	if len(email) != 0 {
		user, err = u.repo.UserByEmail(username)
	}

	if err != nil {
		return nil, common.NewUserError("unable to find a user", common.SQL_ERROR_TYPE)
	}

	if user == nil {
		return nil, common.NewUserError("no such user", common.USER_DOES_NOT_EXISTS_TYPE)
	}

	credentials := user.Credentials[0]

	// Get user with highest ID
	for _, c := range user.Credentials {
		if c.ID > credentials.ID {
			credentials = c
		}
	}

	// Check if the hash is set
	if err := bcrypt.CompareHashAndPassword([]byte(credentials.PasswordHash), []byte(password)); err != nil {
		return nil, common.NewUserError("invalid credentials", common.BAD_CREDENTIALS_ERROR)
	}

	return user, nil
}
