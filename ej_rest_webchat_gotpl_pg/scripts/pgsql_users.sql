-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__allUsers(
)RETURNS SETOF users AS
$BODY$

BEGIN
	-- BUSCAR USUARIOS
	RETURN QUERY SELECT id, uuid, 
						name, email, 
						password, created_at 
						FROM users;
	
	IF NOT FOUND THEN
		--NADA
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__allUsers();

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__getUserByID(
_id INTEGER
)RETURNS TABLE (
id INTEGER,
uuid VARCHAR,
name VARCHAR,
email VARCHAR,
created_at TIMESTAMP
) AS
$BODY$

BEGIN
	-- BUSCAR USER POR ID
	RETURN QUERY SELECT u.id, u.uuid, 
						u.name, u.email, 
						u.created_at 
						FROM users u
						WHERE u.id = _id;
	
	IF NOT FOUND THEN
		--RAISE EXCEPTION 'Sin datos.';
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__getUserByID(1::INTEGER);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__getUserByUUID(
_uuid VARCHAR
)RETURNS SETOF users AS
$BODY$

BEGIN
	-- BUSCAR USER POR UUID
	RETURN QUERY SELECT id, uuid, 
						name, email, 
						password, created_at 
						FROM users 
						WHERE uuid = _uuid;
	
	IF NOT FOUND THEN
		--RAISE EXCEPTION 'Sin datos.';
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__getUserByUUID('e3a9616d-e165-4aac-56d1-92a544955495'::VARCHAR);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__getUserByEmail(
_email VARCHAR
)RETURNS SETOF users AS
$BODY$

BEGIN
	-- BUSCAR USER POR UUID
	RETURN QUERY SELECT id, uuid, 
						name, email, 
						password, created_at 
						FROM users 
						WHERE email = _email;
	
	IF NOT FOUND THEN
		--RAISE EXCEPTION 'Sin datos.';
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__getUserByEmail('zero@mail.com'::VARCHAR);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__updateUser(
_id INTEGER,
_name VARCHAR DEFAULT NULL,
_email VARCHAR DEFAULT NULL
)RETURNS VOID AS
$BODY$

BEGIN
	-- SE ACTUALIZA DATOS DEL USUARIO
	UPDATE users 
	SET name = COALESCE(_name, name), email = COALESCE(_email, email)
	WHERE id = _id;

	IF NOT FOUND THEN
		RAISE EXCEPTION 'No fue posible actualizar datos.';
	END IF;
	
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT pkg_users__updateUser(
1::INTEGER,
'zero'::VARCHAR,
'zero@mail.com'::VARCHAR
)

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__deleteUser(
_id INTEGER
)RETURNS VOID AS
$BODY$
DECLARE
	-- NA
BEGIN
	-- SE ELIMINA EL REGISTRO
	DELETE FROM users WHERE id = _id;
	
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Ocurrio error eliminando usuario.';
	END IF;
 
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT pkg_users__deleteUser(2::INTEGER);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__deleteByUUID(
_uuid VARCHAR
)RETURNS VOID AS
$BODY$
DECLARE
	-- NA
BEGIN
	-- SE ELIMINA EL REGISTRO
	DELETE FROM sessions WHERE uuid = _uuid;
	
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Ocurrio error eliminando sesion.';
	END IF;
 
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT pkg_users__deleteByUUID('7cd8635a-dd06-4bc5-5998-dc0b1700d6f1'::VARCHAR);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__getSession(
_user_id INTEGER
)RETURNS SETOF sessions AS
$BODY$

BEGIN
	-- BUSCAR SESSION POR USER_ID
	RETURN QUERY SELECT id, uuid, 
						email, user_id, 
						created_at 
						FROM sessions 
						WHERE user_id = _user_id;
	
	IF NOT FOUND THEN
		--RAISE EXCEPTION 'Sin datos.';
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__getSession(1::INTEGER);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__checkSession(
_uuid VARCHAR
)RETURNS SETOF sessions AS
$BODY$

BEGIN
	-- BUSCAR SESSION POR USER_ID
	RETURN QUERY SELECT id, uuid, 
						email, user_id, 
						created_at 
						FROM sessions 
						WHERE uuid = _uuid;
	
	IF NOT FOUND THEN
		--RAISE EXCEPTION 'Sin datos.';
	END IF;		
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__checkSession('7cd8635a-dd06-4bc5-5998-dc0b1700d6f1'::VARCHAR);


-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__createUser(
_uuid VARCHAR, 
_name VARCHAR,
_email VARCHAR,
_password VARCHAR,
_created_at TIMESTAMP
) RETURNS RECORD AS
$BODY$
DECLARE
	id_ INTEGER;
	uuid_ VARCHAR;
	created_at_ TIMESTAMP;
BEGIN
	-- SE INSERTA NUEVO USUARIO
	INSERT INTO users (uuid, name, email, password, created_at) 
	VALUES (_uuid, _name, _email, _password, _created_at)
	RETURNING id, uuid, created_at
	INTO id_, uuid_, created_at_;

	IF FOUND THEN
		-- NADA POR EL MOMENTO
	ELSE
		RAISE EXCEPTION 'No fue posible insertar registro.';
	END IF;

	RETURN ROW(id_, uuid_, created_at_);
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__createUser(
'e161641b-6c37-43b8-7b59-70d1db90eee0'::VARCHAR,
'zero2'::VARCHAR,
'zero2@mail.com'::VARCHAR,
'pass123'::VARCHAR,
'2020-09-20 19:12:08'::TIMESTAMP
) AS u(i INTEGER, u VARCHAR, c TIMESTAMP);

-- ============================================================= --

-- FUNC DEF:
CREATE OR REPLACE FUNCTION pkg_users__createSession(
_uuid VARCHAR,
_email VARCHAR,
_user_id INTEGER,
_created_at TIMESTAMP
) RETURNS RECORD AS
$BODY$
DECLARE
	id_ INTEGER;
	uuid_ VARCHAR;
	email_ VARCHAR;
	user_id_ INTEGER;
	created_at_ TIMESTAMP;
BEGIN
	-- SE INSERTA NUEVA SESION
	INSERT INTO sessions (uuid, email, user_id, created_at) 
	VALUES (_uuid, _email, _user_id, _created_at) 
	RETURNING id, uuid, email, user_id, created_at
	INTO id_, uuid_, email_, user_id_, created_at_;

	IF FOUND THEN
		-- NADA POR EL MOMENTO
	ELSE
		RAISE EXCEPTION 'No fue posible insertar sesion.';
	END IF;

	RETURN ROW(id_, uuid_, email_, user_id_, created_at_);
END;
$BODY$
LANGUAGE plpgsql;

-- FUNC TEST:
SELECT * FROM pkg_users__createSession(
'e161641b-6c37-43b8-7b59-70d1db90eee0'::VARCHAR,
'zero@mail.com'::VARCHAR,
1::INTEGER,
'2020-09-20 19:12:08'::TIMESTAMP
) AS s(i INTEGER, u VARCHAR, e VARCHAR, us INTEGER, c TIMESTAMP);
