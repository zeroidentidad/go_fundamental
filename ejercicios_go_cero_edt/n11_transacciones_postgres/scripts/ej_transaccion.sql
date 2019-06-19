-- transacciones
-- start/begin transaction inicia una transaccion

start transaction

-- confirmar la transaccion:
commit

-- revertir la transaccion:
rollback



start transaction

-- crear factura 
insert into invoices (client) values ('Cliente X') returning id;

--crear detalle
insert into invoices_details (invoice_id, product_id, amount)
values(5,1,2),(5,3, 1);

--rollback

commit