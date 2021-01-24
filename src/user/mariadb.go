package user

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MariadbRepo holds a connection
type MariadbRepo struct {
	context *gorm.DB
}

func (r *MariadbRepo) UserByUsername(username string) (*User, error) {
	var user User
	result := r.context.
		Preload("Credentials").
		Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (r *MariadbRepo) UserByEmail(email string) (*User, error) {
	var user User
	result := r.context.
		Preload("Credentials").
		Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (r *MariadbRepo) UserById(id string) (*User, error) {
	var user User
	result := r.context.First(&user, id)
	return &user, result.Error
}

// NewRepository instansiates a repository obect
func NewRepository(username string, password string, hostname string, port string, database string) (*MariadbRepo, error) {
	repo := MariadbRepo{}
	context, err := gorm.Open(mysql.Open(createConnectionString(username, password, hostname, port, database)), &gorm.Config{})
	repo.context = context
	return &repo, err
}

func (r *MariadbRepo) Create(user *User) (string, error) {
	user.ID = uuid.New().String()
	result := r.context.Create(user)
	return user.ID, result.Error
}

func (r *MariadbRepo) Update(user *User) error {
	return r.context.Save(user).Error
}

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

func createConnectionString(username string, password string, hostname string, port string, database string) string {
	if len(hostname) > 0 { // Full config
		return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, hostname, port, database)
	}

	return fmt.Sprintf("%v:%v@/", username, password)
}
