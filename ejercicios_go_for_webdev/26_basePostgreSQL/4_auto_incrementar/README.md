# campo clave de incremento automático
En lugar de crear un ID único, podemos hacer que postgres incremente automáticamente este campo de ID.
 
  Para hacer esto, utilizamos los tipos de datos smallserial, serial o bigserial (no son tipos verdaderos pero por conveniencia).
 
  Esto es como AUTO_INCREMENT en otras bases de datos.

```
CREATE TABLE phonenumbers(
	ID  SERIAL PRIMARY KEY,
	PHONE           TEXT      NOT NULL
);
```

```
INSERT INTO phonenumbers (PHONE) VALUES ( '234-432-5234'), ('543-534-6543'), ('312-123-5432');
```

```
\d phonenumbers
```

```
SELECT * FROM phonenumbers;
```
