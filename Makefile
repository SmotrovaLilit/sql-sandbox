MYSQL_DSN="sandbox-mysql:password@tcp(127.0.0.1:3307)/sandbox-mysql?charset=utf8mb4&parseTime=True&loc=Local"
POSTGRES_DSN="host=localhost user=postgres password=postgres dbname=sandbox-postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
MS_SQL_DSN="sqlserver://sa:Password123@localhost:1434?database=sandboxmssql"

build:
	go build ./application/main.go
restore-mysql: build
	DB_DRIVER=mysql DB_DSN=$(MYSQL_DSN) ./main
restore-postgres: build
	DB_DRIVER=postgres DB_DSN=$(POSTGRES_DSN) ./main
restore-mssql: build
	DB_DRIVER=mssql DB_DSN=$(MS_SQL_DSN) ./main

restore: restore-mysql restore-postgres

