#  crear base de datos
```
CREATE DATABASE employees;
```

## lista de bases de datos
```
\l
```

## conectarse a una base de datos
```
\c <database name>
```

## volver a postgres database
```
\c postgres
```

## ver usuario actual
```
SELECT current_user;
```

## ver base de datos actual
```
SELECT current_database();
```

## drop (remove, delete) base de datos
```
DROP DATABASE <database name>;
```