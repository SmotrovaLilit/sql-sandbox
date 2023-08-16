# Useful commands for SQL Injections

|         | Version                   | Current user                     | Databases                                                                     | Tables                                                                                            | Columns                                                                                                                   |     
|---------|---------------------------|----------------------------------|-------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| Mysql   | `version()`/ `@@version`  | `current_user()` / `current_user`| `select table_schema from information_schema.tables group by table_schema;`   | `select table_name from information_schema.tables where table_schema='sandbox-mysql';`            | `select column_name, data_type from information_schema.columns where table_schema='sandbox-mysql' and table_name='users'` |
| Postgres| `version()`               | `current_user`                   | `select datname from pg_database;`                                            | `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';`                 | `select column_name, data_type from information_schema.columns where table_schema='public' and table_name='users'`        |
| MSSQL   | `@@version`               | `current_user`  / `system_user`  | `select name from sys.databases;`                                             | `select table_name from sandboxmssql.information_schema.tables;`                                  | `select COLUMN_NAME, DATA_TYPE from sandboxmssql.information_schema.columns where TABLE_NAME = 'users';`                  |
| Oracle  | `v$version`               | `select user from dual`          | List of schemas: `select owner from all_tables group by owner`                | `select table_name from all_tables where owner = 'SYS' order by table_name;`                      | `select column_name, data_type from all_tab_columns where table_name = 'MENU';`                                           |  



## Postgres
### Get tables from database Postgres
You can get tables from only current database without additional tools dblink. 
SQL bellow returns an error if current database is not `sandbox-postgres`
```sql
select table_name from sandbox-postgres.information_schema.tables where table_schema = 'public'
```

### List of schemas
```sql
SELECT schema_name
FROM information_schema.schemata;
```

## Oracle
### List of schemas
The database creates schema for each user. Schema is a set of tables, views, procedures, functions, etc.
Owner it is a name of schema.
```sql
select owner from all_tables group by owner;
```
