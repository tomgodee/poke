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

### ISSUE: dotenv lib doesn't works
### ISSUE: PUT seems to work just like PATCH somehow ??? => due to our query update working as intended while PUT and PATCH are simple HTTP verbs => still need more research about this problem => json-patch

### Heroku
To push our local db to the heroku's remote db we need to set up lke the following
```
SET PGUSER=[your_db_username] => in my case it's postgres
SET PGPASSWORD=[your_db_password] => in my case it's zxc321
```
Then to push
```
heroku pg:push [your_db_name] [remote_db_name] --app [your_app_name]
// For me:
// [your_db_name] = poke-development
// [remote_db_name] = DATABASE_URL
// [your_app_name] = postgresql-tapered-69911 
```