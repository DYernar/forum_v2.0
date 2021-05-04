package db

import (
	"database/sql"
	"errors"
	"fmt"

	"forum_v2.0/model"
	"golang.org/x/crypto/bcrypt"
)

//GetLastPost returns the id of the newest post idk why we need it
func GetLastPost() int {
	var i = 0
	db, _ := ConnectDb()
	row, err := db.Query("select rowid from posts")
	fmt.Println(err)
	defer row.Close()
	for row.Next() {
		i++
	}
	return i
}

//GetAllUsers returns the all list of users also useless
func GetAllUsers() []model.User {
	db, _ := ConnectDb()

	row, err := db.Query("select username, email, password from users")
	fmt.Println(err)
	defer row.Close()
	var all []model.User
	for row.Next() {
		var u model.User
		row.Scan(&u.Username, &u.Email, &u.Password)
		all = append(all, u)
	}
	db.Close()
	return all
}

//CheckCredentialsByEmail true if credentials are correct otherwise false
func CheckCredentialsByEmail(email, password string) bool {
	if email == "" || password == "" {
		return false
	}
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return false
	}

	var pass string
	err = db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&pass)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("other error", err.Error())
			db.Close()
			return false
		}
		db.Close()
		return false
	}
	db.Close()
	return CheckPasswordHash(password, pass)

}

//CheckCredentialsByUsername true if credentials are correct otherwise false
func CheckCredentialsByUsername(username, password string) bool {
	if username == "" || password == "" {
		return false
	}
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return false
	}

	var pass string
	err = db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&pass)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("other error", err.Error())
			db.Close()
			return false
		}
		db.Close()
		return false
	}
	db.Close()
	return CheckPasswordHash(password, pass)
}

//CheckCredentials true if credentials are correct otherwise false
func CheckCredentials(usernameOrEmail, password string) bool {
	byEmail := CheckCredentialsByEmail(usernameOrEmail, password)
	byUsername := CheckCredentialsByUsername(usernameOrEmail, password)
	return byEmail || byUsername
}

// CheckPasswordHash compares hash and password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetUserByUsernameCredentials return the user by username and password
func GetUserByUsernameCredentials(username, password string) model.User {
	var user model.User
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return user
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE username = $1 and password = $2", username, password)
	if er != nil {
		fmt.Println(er)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}

//GetUserByEmailCredentials return the user by his email
func GetUserByEmailCredentials(email, password string) model.User {
	var user model.User
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return user
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE email = $1 and password = $2", email, password)
	if er != nil {
		fmt.Println(er)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}

//GetUser return the user by his email
func GetUser(nameOrEmail, password string) model.User {
	byEmail := GetUserByEmailCredentials(nameOrEmail, password)
	byName := GetUserByUsernameCredentials(nameOrEmail, password)
	if byEmail.Username == "" {
		return byName
	}
	return byEmail
}

//GetUserByName returns the user
func GetUserByName(username string) model.User {
	var user model.User
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		fmt.Println("db connection error: ", err.Error())
		return user
	}
	row, er := db.Query("SELECT rowid, username, email FROM users WHERE username = $1 ", username)
	if er != nil {
		db.Close()
		fmt.Println(er)
		return user
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}

//GetUserByEmail returns the user by his email
func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		fmt.Println("db connection error: ", err.Error())
		return user, err
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE email = $1 ", email)
	if er != nil {
		db.Close()
		fmt.Println(er)
		return user, err
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	if user.Username == "" {
		return user, errors.New("not found")
	}
	return user, nil
}

//GetUserByID returns the user by his id
func GetUserByID(userid int) model.User {
	var user model.User
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		fmt.Println("db connection error: ", err.Error())
		return user
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE rowid = $1 ", userid)
	if er != nil {
		db.Close()
		fmt.Println(er)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}

//UsernameExists returns true if username is already in use
func UsernameExists(name string) (bool, error) {
	if name == "" {
		return false, nil
	}
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return false, err
	}

	err = db.QueryRow("SELECT username FROM users WHERE username = $1", name).Scan(&name)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("other error")
			defer db.Close()
			return false, err
		}
		fmt.Println(err)
		defer db.Close()
		return false, nil
	}
	defer db.Close()
	return true, nil
}

//EmailExists returns true if email is occupied
func EmailExists(email string) (bool, error) {
	if email == "" {
		return false, nil
	}
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return false, err
	}

	err = db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("other error")
			defer db.Close()
			return false, err
		}
		fmt.Println(err)
		defer db.Close()
		return false, nil
	}
	defer db.Close()
	return true, nil
}

//SignupUser inserts user into users table
func SignupUser(user model.User) bool {
	db, err := ConnectDb()

	if err != nil {
		return false
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	users, err := db.Prepare("insert into users (username, email, password) values (?, ?, ?)")

	if err != nil {
		fmt.Print("usertable creation err!: ")
		fmt.Println(err)
		return false
	}

	users.Exec(user.Username, user.Email, hashedPassword)

	db.Close()
	return true
}
