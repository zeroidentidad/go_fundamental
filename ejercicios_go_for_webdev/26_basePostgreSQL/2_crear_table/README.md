# crear table
```
CREATE TABLE employees (
   ID INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   RANK           INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL DEFAULT 25500.00,
   BDAY			  DATE DEFAULT '1900-01-01'
);
```

## mostrar tablas en una base de datos (lista abajo)
```
\d
```

## mostrar detalles de una tabla
```
\d <table name>
```

## drop tabla
```
DROP TABLE <table name>;
```

## esquema
Los esquemas nos permiten organizar nuestra base de datos y código de base de datos.

Un esquema es como una carpeta.

En esta carpeta, puede poner tablas, vistas, índices, secuencias, tipos de datos, operadores y funciones.

Sin embargo, a diferencia de las carpetas, los esquemas no se pueden anidar.

Los esquemas proporcionan espacios de nombres.

[Leer más sobre esquemas](https://www.tutorialspoint.com/postgresql/postgresql_schema.htm)