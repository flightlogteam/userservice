package user

import (
	"gorm.io/gorm"
	"log"
	"time"
)

// MariadbRepo holds a connection
type MariadbRepo struct {
	context *gorm.DB
}

// UserByUsername gets a user by its username
func (r *MariadbRepo) UserByUsername(username string) (*User, error) {
	var user User
	result := r.context.
		Where("username = ?", username).First(&user)
	return &user, result.Error
}

// UserByEmail get an user by its email
func (r *MariadbRepo) UserByEmail(email string) (*User, error) {
	var user User
	result := r.context.
		Where("email = ?", email).First(&user)
	return &user, result.Error
}

// UserById fetches a user by the userId
func (r *MariadbRepo) UserById(id string) (*User, error) {
	var user User
	result := r.context.Where("id = ?", id).First(&user)
	log.Println(user)
	return &user, result.Error
}

// NewRepository instansiates a repository obect
func NewRepository(context *gorm.DB) *MariadbRepo {
	repo := MariadbRepo{context: context}
	return &repo
}

// Create creates a user
func (r *MariadbRepo) Create(user *User) (string, error) {
	timestamp := time.Now()
	user.CreatedAt = &timestamp
	user.UpdatedAt = &timestamp

	result := r.context.Create(user)
	return user.ID, result.Error
}

// Update updates a user
func (r *MariadbRepo) Update(user *User) error {
	return r.context.Save(user).Error
}

// UserExists returns an userId if the user exists
func (r *MariadbRepo) UserExists(email string, username string) (string, error) {
	// Email and username are unique. Thereby
	var users []User = make([]User, 2)
	if res := r.context.Where("email = ?", email).
		Or("username = ?", username).
		Find(&users); res.Error != nil {
		return "", nil
	}
	for _, user := range users {
		if user.Username == username {
			return "username", nil
		}

		if user.Email == email {
			return "email", nil
		}
	}
	return "", nil
}
