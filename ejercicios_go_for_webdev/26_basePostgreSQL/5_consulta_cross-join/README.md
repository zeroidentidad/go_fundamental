# consultas

## consultar desde más de una tabla

Esta consulta es un **inner join**.
```
SELECT employees.NAME, phonenumbers.PHONE FROM employees INNER JOIN phonenumbers ON employees.ID = phonenumbers.EMP_ID;
```

# consultas join

Consultas join nos permiten seleccionar registros de dos o más tablas.

Una consulta join combina columnas de una o más tablas: une un grupo de columnas de diferentes tablas.

Hay cinco tipos de joins en postgres:

## cross join
Un cross join devuelve el ** producto cartesiano ** de las filas de las tablas en la unión. En otras palabras, producirá filas que combinan cada fila de la primera tabla con cada fila de la segunda tabla.

```
CREATE TABLE person (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL
);
```

```
INSERT INTO person (NAME) VALUES ('Shen'), ('Daniel'), ('Juan'), ('Arin'), ('McLeod');
```

```
CREATE TABLE sport (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL,
   P_ID         INT      references person(ID)
);
```

```
INSERT INTO sport (NAME, P_ID) VALUES ('Surf',1),('Soccer',3),('Ski',3),('Sail',3),('Bike',3);
```

```
SELECT person.NAME, sport.NAME FROM person CROSS JOIN sport;
```
