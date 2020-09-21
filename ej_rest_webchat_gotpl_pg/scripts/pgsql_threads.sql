-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_threads__allThreads(
)RETURNS SETOF threads AS
$BODY$

BEGIN
	-- BUSCAR HILOS
	RETURN QUERY SELECT id, uuid, 
						topic, user_id, 
						created_at 
						FROM threads 
						ORDER BY 
						created_at DESC;
	
	IF NOT FOUND THEN
		--NADA
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_threads__allThreads();

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_threads__getUser(
_id_user INTEGER
)RETURNS TABLE (
id INTEGER,
uuid VARCHAR,
name VARCHAR,
email VARCHAR,
created_at TIMESTAMP
) AS
$BODY$

BEGIN
	
	-- BUSCAR USUARIO
	RETURN QUERY SELECT u.id, u.uuid, 
						u.name, u.email, 
						u.created_at 
						FROM users u
						WHERE u.id = _id_user;

	IF NOT FOUND THEN
		--RAISE EXCEPTION 'No encontrado.';
	END IF;	
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_threads__getUser(1::INTEGER);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_threads__getThreadByUUID(
_uuid VARCHAR
)RETURNS SETOF threads AS
$BODY$

BEGIN
	-- BUSCAR HILO POR UUID
	RETURN QUERY SELECT id, uuid, 
						topic, user_id, 
						created_at 
						FROM threads 
						WHERE uuid = _uuid;
	
	IF NOT FOUND THEN
		--RAISE EXCEPTION 'Sin datos.';
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_threads__getThreadByUUID('e163944b-6c37-43b8-7b59-79d1db86eee6'::VARCHAR);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_threads__numReplies(
_thread_id INTEGER
) RETURNS INTEGER AS
$BODY$
DECLARE
	_count INTEGER;
BEGIN
	-- CUENTA PUBLICACIONES HILO
	SELECT COUNT(*) 
	INTO _count 
	FROM posts 
	WHERE thread_id = _thread_id;

	IF _count=0 THEN
		-- NADA
	END IF;

	RETURN _count;
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_threads__numReplies(1::INTEGER);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_threads__getPosts(
_thread_id INTEGER
)RETURNS SETOF posts AS
$BODY$

BEGIN
	-- BUSCAR PUBLICACIONES DE HILO
	RETURN QUERY SELECT id, uuid, 
						body, user_id, 
						thread_id, created_at 
						FROM posts 
						WHERE thread_id = _thread_id;
	
	IF NOT FOUND THEN
		--RAISE EXCEPTION 'Sin datos.';
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_threads__getPosts(1::INTEGER);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_threads__createThread(
_uuid VARCHAR, 
_topic TEXT,
_user_id INTEGER,
_created_at TIMESTAMP
) RETURNS RECORD AS
$BODY$
DECLARE
	id_ INTEGER;
	uuid_ VARCHAR;
	topic_ TEXT;
	user_id_ INTEGER;
	created_at_ TIMESTAMP;
BEGIN
	-- SE INSERTA NUEVO HILO
	INSERT INTO threads (uuid, topic, user_id, created_at) 
	VALUES (_uuid, _topic, _user_id, _created_at) 
	RETURNING id, uuid, topic, user_id, created_at 
	INTO id_, uuid_, topic_, user_id_, created_at_;

	IF FOUND THEN
		-- NADA POR EL MOMENTO
	ELSE
		RAISE EXCEPTION 'No fue posible insertar registro.';
	END IF;

	RETURN ROW(id_, uuid_, topic_, user_id_, created_at_);
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT pkg_threads__createThread(
'e162644b-6c37-43b8-7b59-79d1db80eee0'::VARCHAR,'Nuevas dudas'::TEXT,
1::INTEGER,'2020-09-20 19:12:08'::TIMESTAMP
);

SELECT * FROM pkg_threads__createThread(
'e161641b-6c35-43b8-7b59-70d1db80eee0'::VARCHAR,'Retos y proyectos'::TEXT,
1::INTEGER,'2020-09-20 19:12:08'::TIMESTAMP
) AS t(i INTEGER, u VARCHAR, t TEXT, us INTEGER, c TIMESTAMP);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_threads__createPost(
_uuid VARCHAR, 
_body TEXT,
_user_id INTEGER,
_thread_id INTEGER,
_created_at TIMESTAMP
) RETURNS RECORD AS
$BODY$
DECLARE
	id_ INTEGER;
	uuid_ VARCHAR;
	body_ TEXT;
	user_id_ INTEGER;
	thread_id_ INTEGER;
	created_at_ TIMESTAMP;
BEGIN
	-- SE INSERTA NUEVA PUBLICACION EN HILO
	INSERT INTO posts (uuid, body, user_id, thread_id, created_at) 
	VALUES (_uuid, _body, _user_id, _thread_id, _created_at)
	RETURNING id, uuid, body, user_id, thread_id, created_at
	INTO id_, uuid_, body_, user_id_, thread_id_, created_at_;

	IF FOUND THEN
		-- NADA POR EL MOMENTO
	ELSE
		RAISE EXCEPTION 'No fue posible insertar registro.';
	END IF;

	RETURN ROW(id_, uuid_, body_, user_id_, thread_id_, created_at_);
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_threads__createPost(
'e162644b-6c37-43b8-0b57-79d1dc80eue9'::VARCHAR,'go.dev sitio de docu de 3rd libs'::TEXT,
1::INTEGER, 4::INTEGER, 
'2020-09-20 19:12:08'::TIMESTAMP
) AS p(i INTEGER, u VARCHAR, b TEXT, us INTEGER, th INTEGER, c TIMESTAMP);
