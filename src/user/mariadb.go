package user

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func NewRepository(username string, password string, hostname string, port string, database string) (*MariadbRepo, error) {
	err := ensureDatabaseExists(username, password, hostname, port, database)

	if err != nil {
		return nil, err
	}

	repo := MariadbRepo{}
	context, err := gorm.Open(mysql.Open(createConnectionString(username, password, hostname, port, database)), &gorm.Config{})

	context.AutoMigrate(User{})
	repo.context = context
	return &repo, err
}

func ensureDatabaseExists(username string, password string, hostname string, port string, database string) error {
	connectionWithoutDatabase, err := gorm.Open(mysql.Open(createConnectionString(username, password, hostname, port, "")), &gorm.Config{})

	if err != nil {
		return err
	}

	tx := connectionWithoutDatabase.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", database))
	tx.Commit()

	return nil
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

func createConnectionString(username string, password string, hostname string, port string, database string) string {
	connectionString := fmt.Sprintf("%v:%v@", username, password)

	if len(hostname) > 0 {
		connectionString += fmt.Sprintf("tcp(%v:%v)", hostname, port)
	}

	if len(database) > 0 {
		connectionString += fmt.Sprintf("/%v", database)
	} else {
		connectionString += "/"
	}

	return connectionString + "?parseTime=true"

}
