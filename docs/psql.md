To clear terminal: ```\! cls```
To create a db: ```CREATE DATABASE [dbname] ```
To list all db: ``` \l ```
To connect to a db: ``` \c [dbname] ```
To drop a db: ```DROP DATABASE [ IF EXISTS ] name```
To create a table: 
```
CREATE TABLE table_name(
   column1 datatype,
   column2 datatype,
   column3 datatype,
   .....
   columnN datatype,
   PRIMARY KEY( one or more columns )
);
```
To drop a table: ```DROP TABLE table_name;```

### To backup a db
```
pg_dump -U [user_name]-W -F t [db_name_needed_backup] > [backup_file_name_path]
pg_dump -U postgres -W -F t poke_development > d:\pgbackup\poke.tar
```
### To restore a db
```
pg_restore -U [user_name] --dbname=[target_db_name] --verbose [file_name_path]
pg_restore -U postgres --dbname=poke_test --verbose d:\pgbackup\poke.tar
```


### ISSUE: dotenv lib doesn't works
### ISSUE: PUT seems to work just like PATCH somehow ??? => due to our query update working as intended while PUT and PATCH are simple HTTP verbs => still need more research about this problem => json-patch

### Learn about type in postgresql, how you could define a new type, assign a field to have value of that type, how to drop that type when migrating down. 