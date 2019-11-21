package models

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL UNIQUE, 
        password VARCHAR(64) NOT NULL,
        email VARCHAR(50),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

// Constructores estructuras:

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Metodos DB:

func (this *User) Save() {
	if this.ID == 0 {
		this.insert()
	} else {
		this.update()
	}
}

func (this *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := Exec(sql, this.Username, this.Password, this.Email)
	this.ID, _ = result.LastInsertId() //int64
}

func (this *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=?"
	Exec(sql, this.Username, this.Password, this.Email)
}

func (this *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	Exec(sql, this.ID)
}

func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	}

	return user
}
