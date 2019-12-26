# sintaxis general

## select
```
SELECT <fields> FROM <table>;
```

```
SELECT * FROM employees;
```

```
SELECT name, score FROM employees;
```

## cross join
```
SELECT <fields> FROM <table1> CROSS JOIN <table2>;
```

```
SELECT person.NAME, sport.NAME FROM person CROSS JOIN sport;
```

## inner join
```
SELECT <fields> FROM <table> INNER JOIN <table>
ON <pkey> = <fkey>;
```

```
SELECT person.NAME, sport.NAME FROM person INNER JOIN sport
ON person.ID = sport.P_ID;
```

# inner join

Un inner join nos permite seleccionar registros de dos tablas.
 
Usamos un inner join arriba cuando solicitamos los números de teléfono asociados con un empleado:

```
SELECT employees.NAME, phonenumbers.PHONE FROM employees INNER JOIN phonenumbers ON employees.ID = phonenumbers.EMP_ID;
```

Podemos usar uno con las tablas ```people``` y ```sports``` tambien, si quisiéramos, ya que estas tablas están conectadas (recordar ```P_ID INT references person(ID)```).

```
SELECT person.NAME, sport.NAME FROM person INNER JOIN sport ON person.ID = sport.P_ID;
```

Así es como wikipedia explica un inner join:

Un inner join requiere que cada fila de las dos tablas unidas tenga filas coincidentes, y es una operación de combinación de uso común en aplicaciones, pero no se debe suponer que es la mejor opción en todas las situaciones.

Inner join crea una nueva tabla de resultados combinando valores de columna de dos tablas (A y B) basado en el predicado de unión. 

La consulta compara cada fila de A con cada fila de B para encontrar todos los pares de filas que satisfacen el predicado de unión. 

Cuando el predicado de unión se satisface haciendo coincidir valores no NULL, los valores de columna para cada par de filas coincidentes de A y B se combinan en una fila de resultados.
