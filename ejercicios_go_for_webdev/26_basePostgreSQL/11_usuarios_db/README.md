# usuarios y privilegios

## ver usuario actual
```
SELECT current_user;
```

## detalles de los usuarios
```
\du
```

## crear usuario
```
CREATE USER james WITH PASSWORD 'password';
```

## conceder privilegios
privileges: SELECT, INSERT, UPDATE, DELETE, RULE, ALL
```
GRANT ALL PRIVILEGES ON DATABASE company to james;
```

## revocar privilegios
```
REVOKE ALL PRIVILEGES ON DATABASE company from james;
```

## alterar
```
ALTER USER james WITH SUPERUSER;
```

```
ALTER USER james WITH NOSUPERUSER;
```

## eliminar
```
DROP USER james;
```
