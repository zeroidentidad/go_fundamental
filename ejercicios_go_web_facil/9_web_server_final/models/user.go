package models

import (
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL UNIQUE, 
        password VARCHAR(64) NOT NULL,
        email VARCHAR(50),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	createdAt time.Time
}

type Users []User

// Constructores estructuras:

func NewUser(username, password, email string) (*User, error) {
	user := &User{Username: username, Email: email}

	if err := user.Valid(); err != nil {
		return &User{}, err
	}

	err := user.SetPassword(password)
	return user, err
}

func CreateUser(username, password, email string) (*User, error) {
	user, err := NewUser(username, password, email)

	if err != nil {
		return &User{}, err
	}

	err = user.Save()

	if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
		err = errorUsernameExists
	}

	return user, err
}

// Metodos DB:

func (this *User) Save() error {
	if this.ID == 0 {
		return this.insert()
	} else {
		return this.update()
	}
}

func (this *User) insert() error {
	sql := "INSERT users SET username=?, password=?, email=?"
	id, err := InsertData(sql, this.Username, this.Password, this.Email)
	//this.ID, err = result.LastInsertId() //int64
	this.ID = id
	return err
}

func (this *User) update() error {
	sql := "UPDATE users SET username=?, password=?, email=?"
	_, err := Exec(sql, this.Username, this.Password, this.Email)
	return err
}

func (this *User) Delete() error {
	sql := "DELETE FROM users WHERE id=?"
	_, err := Exec(sql, this.ID)
	return err
}

func GetUser(sql string, conditional interface{}) *User {
	user := &User{}
	rows, err := Query(sql, conditional)
	if err != nil {
		return user
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.createdAt)
	}

	return user
}

func GetUsers() Users {
	sql := "SELECT id, username, password, email, created_at FROM users"
	users := Users{}
	rows, _ := Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.createdAt)
		users = append(users, user)
	}

	return users
}

func (this *User) GetCreatedDate() time.Time {
	return this.createdAt
}

func (this *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errorPasswordEncryption
	}
	this.Password = string(hash)
	return nil
}

func Login(username, password string) (*User, error) {
	user := GetUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return &User{}, errorLogin
	}

	return user, nil
}

func GetUserByUsername(username string) *User {
	sql := "SELECT id, username, password, email, created_at FROM users WHERE username=?"
	return GetUser(sql, username)
}

func GetUserById(id int) *User {
	sql := "SELECT id, username, password, email, created_at FROM users WHERE id=?"
	return GetUser(sql, id)
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errorEmail
	}
	return nil
}

func (this *User) Valid() error {
	if err := ValidEmail(this.Email); err != nil {
		return err
	}

	if err := ValidUsername(this.Username); err != nil {
		return err
	}

	return nil
}

func ValidUsername(username string) error {
	if username == "" {
		return errorUsername
	}

	if len(username) < 5 {
		return errorShortUsername
	}

	if len(username) > 30 {
		return errorLargeUsername
	}

	return nil
}
