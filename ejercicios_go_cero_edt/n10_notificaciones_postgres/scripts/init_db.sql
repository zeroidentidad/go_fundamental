
psql -U postgres

------------------------------------------------------------
create user unotifica password 'x';

create database regs_notificacion owner=unotifica;

create table productos (
    id serial,
    producto character varying(100),
    precio numeric(10,2) not null,
    constraint productos_id_pk primary key (id),
    constraint productos_producto_uk unique (producto)
);
------------------------------------------------------------
-- TEST USO LISTEN --
-- habilitar canal de escucha --

LISTEN canal;

-- prueba enviar notificacion al canal --

NOTIFY canal, 'Hola mundo';

-- dejar de escuchar canal --

UNLISTEN canal;

------------------------------------------------------------

-- funcion de notificacion --

create function notify_event()
returns trigger language 'plpgsql'
as $$
declare
	datos json;
	notification json;
begin
	if (TG_OP='DELETE') then
		datos = row_to_json(old);
	else
		datos = row_to_json(new);
	end if;
	
	notification = jsonb_build_object(
	'tabla', TG_TABLE_NAME,
	'accion', TG_OP,
	'datos', datos
	);

	perform pg_notify('canal', notification::text);
	return null;
end;
$$

-- trigger que hara uso de la funcion --

create trigger productos_notificacion
	after insert or delete or update
	on productos
	for each row
	execute procedure notify_event();

-- test LISTEN en la DB desde terminal --

LISTEN canal;

-- insert desde editor u otra ventana de conexion a la BD --

insert into productos (producto, precio) values ('Computadora X', 12345.50);
insert into productos (producto, precio) values ('Computadora X2', 22345.50);