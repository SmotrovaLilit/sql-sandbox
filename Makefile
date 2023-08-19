build:
	go build ./main.go
fill-mysql: build
	./main --dbDriver=mysql fillDatabase
fill-postgres: build
	./main --dbDriver=postgres fillDatabase
fill-mssql: build
	./main --dbDriver=mssql fillDatabase
fill: fill-mysql fill-postgres fill-mssql

run-mysql-insert: build
	./main --dbDriver=mysql runQuery "SELECT * FROM users;INSERT into users (username, password) VALUES ('dsadas15', 'password')"
run-postgres-insert: build
	./main --dbDriver=postgres runQuery "SELECT * FROM users;INSERT into users (username, password) VALUES ('dsadas12', 'password');"
run-postgres-select: build
	./main --dbDriver=postgres runQuery "SELECT name, author  FROM books where id=1;SELECT username as name, password as author FROM users;"
run-mssql-select: build
	./main --dbDriver=mssql runQuery "SELECT name, author  FROM books where id=1;SELECT username as name, password as author FROM users;"
run-mssql-insert: build
	./main --dbDriver=mssql runQuery "SELECT * FROM users;INSERT into users (username, password) VALUES ('dsadas15', 'password')"


