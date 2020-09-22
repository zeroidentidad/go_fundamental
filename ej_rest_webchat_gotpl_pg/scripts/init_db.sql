-->> CONEXION CON POSTGRES/SUPERUSER para definir:

-- USUARIO DB:
CREATE USER webchat PASSWORD 'x1234567' LOGIN SUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION; -- superuser por si exts

-- TABLESPACE (opcional): 
CREATE TABLESPACE ts_webchat OWNER webchat LOCATION E'C:\\tablespaces\\webchat'; -- previamente crear carpeta en disco

-- CREACIÓN DE LA BASE DE DATOS
CREATE DATABASE gowebchat OWNER = webchat TABLESPACE = ts_webchat;

----------------------------------------------------------------------------------

-->> CONEXION CON USUARIO 'webchat' para crear las tablas 
  -- y las funciones (pgsql_threads.sql, pgsql_users.sql):


create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null       
);

create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null  
);