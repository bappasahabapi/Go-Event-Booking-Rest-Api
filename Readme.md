## Run the server

-` go run .` ||  `go run main.go`


### Steps:

✅ 01-project initialized with server and gin framework 

-  `go mod init bappa.com/rest`
-  `go get -u github.com/gin-gonic/gin`

✅ 02-Add and Test Create and Get event 


✅ 03-Initilazation of SQL Database 

 -`https://github.com/mattn/go-sqlite3`
    - `go get github.com/mattn/go-sqlite3`


🔋 Database Connection Part

```go
var DB *sql.DB
 
func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "api.db")
 
    if err != nil {
        panic("Could not connect to database.")
    }
 
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)
 
    createTables()
}
```
✅ 04-Insert & Read Data from Database



💡🌟🔆 05-Beautify console and json data

- `go get github.com/fatih/color`

- use like this

import ("github.com/fatih/color")

color.Cyan("🔋 🚀 Server running at http://localhost:8080")