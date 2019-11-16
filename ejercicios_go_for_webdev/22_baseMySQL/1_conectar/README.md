# Usando MySQL

1. Necesitaremos un controlador MySQL, ej:
  - go get github.com/go-sql-driver/mysql
  - [documentación](https://github.com/go-sql-driver/mysql#installation)
  - [todos los controladores SQL](https://github.com/golang/go/wiki/SQLDrivers)
  - [libro de Astaxie](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html)
2. Incluir controlador en importaciones
  - _ "github.com/go-sql-driver/mysql"
  - [documentación](https://github.com/go-sql-driver/mysql#usage)
3. Determinar el nombre del origen de datos
  - user:password@tcp(localhost:5555)/dbname?charset=utf8
  - [documentación](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
4. Abrir conexion
  - db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")

[package sql](https://godoc.org/database/sql)