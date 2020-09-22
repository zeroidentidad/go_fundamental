package db

import (
	"time"
)

// Nota pg funcs:
// Postgres no devuelve automáticamente el último ID de inserción,
// porque sería asumir que siempre se esta usando una secuencia.
// Se necesita usar keyword RETURNING en la inserción para obtener esta info.

// crear una nueva sesión para un usuario existente
func (user User) CreateSession() (session Session, err error) {
	statement := `SELECT * FROM pkg_users__createSession($1, $2, $3, $4) AS s(i INTEGER, u VARCHAR, e VARCHAR, us INTEGER, c TIMESTAMP)`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// usar QueryRow para devolver una fila y escanear el id devuelto en la estructura de Session
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	return
}

// obtener la sesión para un usuario existente
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow(`SELECT * FROM pkg_users__getSession($1)`, user.Id).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	return
}

// verificar si la sesión es válida en la base de datos
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow(`SELECT * FROM pkg_users__checkSession($1)`, session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}

	return
}

// eliminar sesión de la base de datos
func (session *Session) DeleteByUUID() (err error) {
	statement := `SELECT pkg_users__deleteByUUID($1)`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)

	return
}

// obtener el usuario de la sesión
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow(`SELECT * FROM pkg_users__getUserByID($1)`, session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return
}

// eliminar todas las sesiones de la base de datos
func SessionDeleteAll() (err error) {
	statement := `DELETE FROM sessions`
	_, err = Db.Exec(statement)

	return
}

// crear un nuevo usuario, guardar información del usuario en la base de datos
func (user *User) Create() (err error) {
	statement := `SELECT * FROM pkg_users__createUser($1, $2, $3, $4, $5) AS u(i INTEGER, u VARCHAR, c TIMESTAMP)`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// usar QueryRow para devolver una fila y escanear id devuelto en la estructura de User
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)

	return
}

// eliminar usuario de la base de datos
func (user *User) Delete() (err error) {
	statement := `SELECT pkg_users__deleteUser($1)`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	return
}

// actualizar información del usuario en la base de datos
func (user *User) Update() (err error) {
	statement := `SELECT pkg_users__updateUser($1, $2, $3)`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email)

	return
}

// eliminar todos los usuarios de la base de datos
func UserDeleteAll() (err error) {
	statement := `DELETE FROM users`
	_, err = Db.Exec(statement)

	return
}

// obtener todos los usuarios en la base de datos y devolverlos
func Users() (users []User, err error) {
	rows, err := Db.Query(`SELECT * FROM pkg_users__allUsers()`)
	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()

	return
}

// obtener un usuario dado el correo electrónico
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow(`SELECT * FROM pkg_users__getUserByEmail($1)`, email).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return
}

// obtener un solo usuario dado el UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow(`SELECT * FROM pkg_users__getUserByUUID($1)`, uuid).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return
}
