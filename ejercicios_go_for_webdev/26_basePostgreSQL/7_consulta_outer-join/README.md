# outer join

## left outer join
 Un left outer join te da todo en una tabla y también los registros coincidentes en otra tabla.
 
 Para las tablas A y B un left outer join daría todas las filas de la tabla "izquierda" (A), incluso si la condición de unión no encuentra ninguna fila coincidente en la tabla "derecha" (B).
  
 Esto significa que si la cláusula ON coincide con 0 (cero) filas en B (para una fila dada en A), la unión aún devolverá una fila en el resultado (para esa fila), pero con NULL en cada columna de B.
 
  ```
  SELECT person.NAME, sport.NAME FROM person LEFT OUTER JOIN sport ON person.ID = sport.P_ID;
  ```
 
## right outer join
Un right outer join es como un left outer join, pero para la tabla de la derecha.

```
INSERT INTO sport (NAME) VALUES ('Squirrel Suit Flying');
```

```
  SELECT person.NAME, sport.NAME FROM person RIGHT OUTER JOIN sport ON person.ID = sport.P_ID;
```

## full outer join
Un full outer join es como ejecutar ambos left outer join y un right outer join al mismo tiempo. Te da todo, desde todas las tablas, y lo que coincide con lo que coincide.

```
  SELECT person.NAME, sport.NAME FROM person FULL OUTER JOIN sport ON person.ID = sport.P_ID;
```
