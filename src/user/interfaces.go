package user

// Writer describes write related actions for a user
type Writer interface {
	Create(user *User) (string, error)
	Update(user *User) error
}

// Reader describes read related actions for a user
type Reader interface {
	UserByEmail(email string) (*User, error)
	UserById(id string) (*User, error)
}

type Repository interface {
	Reader
	Writer
	// UserExists checks if there is any row with a matching email or username
	UserExists(email string, username string) (string, error)
}

// Services define user service layer
type Service interface {
	Reader
	Writer
}
