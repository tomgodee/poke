# poke
## Written in Golang
## Using Gin for routing 
## Postgresql as database, 
## The migrate package helps with migration 
## Using delve for debugging

### To run migration up run in root (application level) this command (Run migration with driver to postgresql, username=postgresql, password=zxc321, port=5432, database name = poke_development, sslmode = disabled)
```
migrate -database postgres://postgres:zxc321@localhost:5432/poke_development?sslmode=disable -path db/migrations up
```
### To run migration down
```
migrate -database postgres://postgres:zxc321@localhost:5432/poke_development?sslmode=disable -path db/migrations down
```
### If migration fails then all migrations will be flaged to be dirty, to force back to a clean version run (V as the number of the clean version)
```
migrate -database postgres://postgres:zxc321@localhost:5432/poke_development?sslmode=disable -path db/migrations force V
```

### To debug 
```
cd /cmd/poke
dlv debug
break main.main
break [folder_name] [func_name]
```
