# poke
## Written in Golang
## Using Gin for routing 
## Postgresql as database, 
## The migrate package helps with migration 

### To run migration up run in root this command (Run migration with driver to postgresql, username=postgresql, password=zxc321, port=5432, database name = example, sslmode = disabled)
```
migrate -database postgres://postgres:zxc321@localhost:5432/example?sslmode=disable -path db/migrations up
```
### To run migration down
```
migrate -database postgres://postgres:zxc321@localhost:5432/example?sslmode=disable -path db/migrations down
```