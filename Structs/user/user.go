package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	birthdate string
	createdAt time.Time
}

// Struct embedding
type Admin struct {
	email    string
	password string
	User
}

// Methods for struct
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstname, u.lastname, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstname = ""
	u.lastname = ""
}

// Constructor
func New(firstname, lastname, birthdate string) (*User, error) {

	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstname: firstname,
		lastname:  lastname,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstname: "ADMIN",
			lastname:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
