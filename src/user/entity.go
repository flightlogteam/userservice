package user

import (
	"fmt"
	"time"
)

// User models a db user
type User struct {
	ID          string        `gorm:"primaryKey;column:id;size:32"`
	Givenname   string        `gorm:"column:givenname;size:64"`
	Familyname  string        `gorm:"column:familyname;size:64"`
	Email       string        `gorm:"column:email;size:320"`
	Active      Bit           `gorm:"column:active"`
	Username    string        `gorm:"column:username;size:32"`
	Credentials []Credentials `gorm:"foreignKey:userId"`
	Privacy     PrivacyLevel  `gorm:"column:privacylevel"`
	UpdatedAt   *time.Time    `gorm:"column:updatedat;type:time"`
	CreatedAt   *time.Time    `gorm:"column:createdat;type:time"`
	DeletedAt   *time.Time    `gorm:"column:deletedat;type:time"`
}

type Credentials struct {
	ID           int        `gorm:"primaryKey;column:id"`
	UserId       string     `gorm:"column:userId;size:32"`
	PasswordHash string     `gorm:"column:passwordHash;size:100"`
	UpdatedAt    *time.Time `gorm:"column:updatedat;type:time"`
	CreatedAt    *time.Time `gorm:"column:createdat;type:time"`
	DeletedAt    *time.Time `gorm:"column:deletedat;type:time"`
}

func (Credentials) TableName() string {
	return "Credentials"
}

func (User) TableName() string {
	return "User"
}

type Bit bool

func (b *Bit) Scan(src interface{}) error {
	unt, ok := src.([]uint8)

	if ok {
		*b = unt[0] > 0
		return nil
	}

	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("unexpected type for bit: %T", src)
	}
	switch str {
	case "\x00":
		*b = Bit(false)
	case "\x01":
		*b = Bit(true)
	}
	return nil
}

// PrivacyLevel is a type that describes a users level of privacy
type PrivacyLevel int

const (
	// ProtectedLevel describes a level where a user can be seen by anyone who has a flightlog login
	ProtectedLevel PrivacyLevel = 0
	// PrivateLevel describes a level where a user needs to grant access to those who wants access
	PrivateLevel PrivacyLevel = 1
	// PublicLevel describes a user that anyone can watch with or without login
	PublicLevel PrivacyLevel = 2
)
