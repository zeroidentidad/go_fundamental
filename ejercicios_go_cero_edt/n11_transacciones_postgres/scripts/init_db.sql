-- crear db:
create database transacciones_go;

-- crear tablas:

-- crear tabla de productos
create table products (
id serial not null,
name varchar(50) not null,
price numeric(10,2) not null,
amount int not null,
constraint products_id_pk primary key (id),
constraint products_amount_ck check (amount >= 0)
);

-- crear tabla  de facturacion (encabezado)
create table invoices(
id serial not null,
client varchar(100) not null,
invoice_date date not null default now(),
constraint invoices_id_pk primary key (id)
);

-- crear tabla de detalle de facturacion
create table invoices_details(
id serial not null, 
invoice_id int not null,
product_id int not null,
amount int not null,
constraint invoices_details_id_pk primary key (id),
constraint invoices_details_invoice_id_fk foreign key (invoice_id) references invoices (id) on update restrict on delete restrict,
constraint invoices_details_product_id_fk foreign key (product_id) references products (id) on update restrict on delete restrict
);

-- crear funciones - triggers:

-- crear funcion de actualizacion del stock:

create or replace function stock_update()
returns trigger as
$BODY$
begin
	update products set amount = amount - new.amount where id = new.product_id;
	return new;
end;
$BODY$
language plpgsql;

-- crear trigger para uso de la funcion de actualizacion:
create trigger stock_update
after insert
on invoices_details
for each row execute procedure stock_update();