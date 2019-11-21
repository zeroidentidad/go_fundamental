package models

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL UNIQUE, 
        password VARCHAR(64) NOT NULL,
        email VARCHAR(50),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users []User

// Constructores estructuras:

func NewUser(username, password string) *User {
	user := &User{Username: username, Password: password}
	return user
}

// Metodos:

func (this *User) Save() {
	sql := "INSERT users SET username=?, password=?"
	Exec(sql, this.Username, this.Password)
}
