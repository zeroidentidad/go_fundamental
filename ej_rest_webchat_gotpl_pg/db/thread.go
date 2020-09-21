package db

import "time"

// formatear la fecha CreatedAt para que se vea mejor en pantalla
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// obtener la cantidad de publicaciones en un hilo
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT * FROM pkg_threads__numReplies($1)", thread.Id)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()

	return
}

// obtener publicaciones en un hilo
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT * FROM pkg_threads__getPosts($1)", thread.Id)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()

	return
}

// crear un hilo nuevo
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "SELECT * FROM pkg_threads__createThread($1, $2, $3, $4) AS t(i INTEGER, u VARCHAR, t TEXT, us INTEGER, c TIMESTAMP)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// usar QueryRow para devolver una fila y escanear id devuelto en la estructura de Session
	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)

	return
}

// crear una nueva publicación en un hilo
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "SELECT * FROM pkg_threads__createPost($1, $2, $3, $4, $5) AS p(i INTEGER, u VARCHAR, b TEXT, us INTEGER, th INTEGER, c TIMESTAMP)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// usar QueryRow para devolver una fila y escanear id devuelto en la estructura de Session
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)

	return
}

// obtener todos los hilos de la base de datos y devolverlos
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT * FROM pkg_threads__allThreads()")
	if err != nil {
		return
	}

	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()

	return
}

// obtener un hilo por UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT * FROM pkg_threads__getThreadByUUID($1);", uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)

	return
}

// obtener el usuario que inició el hilo.
func (thread *Thread) User() (user User) {
	user = User{}
	_ = Db.QueryRow("SELECT * FROM pkg_threads__getUser($1)", thread.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return
}

// obtener el usuario que escribió la publicación
func (post *Post) User() (user User) {
	user = User{}
	_ = Db.QueryRow("SELECT * FROM pkg_threads__getUser($1)", post.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return
}
