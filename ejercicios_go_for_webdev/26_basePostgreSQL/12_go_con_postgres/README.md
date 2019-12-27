# Go y postgres

driver
```
go get github.com/lib/pq
```

## crear una db

```
CREATE DATABASE bookstore;
```

## crear usuario
```
CREATE USER bond WITH PASSWORD 'password';
```

## conceder privilegios
```
GRANT ALL PRIVILEGES ON DATABASE bookstore to bond;
```
