package model

import (
	"fmt"
	"regexp"
)

// User just user
type User struct {
	UserID   int
	Username string
	Email    string
	Password string
	Posts    []Post
}

//Print prints the userdata
func (user User) Print() {
	fmt.Println("Username is ", user.Username)
	fmt.Println("Email is ", user.Email)
	fmt.Println("Password is ", user.Password)
}

//IsValidEmail return true if email is valid
func (user User) IsValidEmail() bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(user.Email) < 3 && len(user.Email) > 254 {
		return false
	}
	return emailRegex.MatchString(user.Email)
}

//IsValidPassword password validation
func (user User) IsValidPassword() string {
	if user.Password == "" {
		return "password cannot be empty"
	}
	if len(user.Password) < 5 {
		return "password length should be atleast 5 characters"
	}
	return ""
}
