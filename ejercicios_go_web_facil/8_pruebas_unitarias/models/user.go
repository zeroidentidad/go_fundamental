package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL UNIQUE, 
        password VARCHAR(64) NOT NULL,
        email VARCHAR(50),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	createdAt time.Time
}

type Users []User

// Constructores estructuras:

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Email: email}
	user.SetPassword(password)
	return user
}

func CreateUser(username, password, email string) (*User, error) {
	user := NewUser(username, password, email)
	err := user.Save()
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

func (this *User) SetPassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	this.Password = string(hash)
}

func Login(username, password string) bool {
	user := GetUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}

func GetUserByUsername(username string) *User {
	sql := "SELECT id, username, password, email, created_date FROM users WHERE username=?"
	return GetUser(sql, username)
}

func GetUserById(id int) *User {
	sql := "SELECT id, username, password, email, created_date FROM users WHERE id=?"
	return GetUser(sql, id)
}
