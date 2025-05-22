## All Install Packages:

-  `go mod init bappa.com/rest`
-  `go get -u github.com/gin-gonic/gin`
-  `go get github.com/mattn/go-sqlite3`
- `go get github.com/fatih/color`  for color the console
- `go install github.com/gravityblast/fresh@latest` for auto restart


## Run the server

-` go run .` ||  `go run main.go` || `fresh`
- `fresh` for auto restart 


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

Preparing Statements vs Directly Executing Queries `(Prepare() vs Exec()/Query())`

We started sending SQL commands to the SQLite database.

And we did this by following different approaches:

- DB.Exec() (when we created the tables)

- Prepare() + stmt.Exec() (when we inserted data into the database)

- DB.Query() (when we fetched data from the database)

Using Prepare() is 100% optional! You could send all your commands directly via Exec() or Query().

The difference between those two methods then just is 

- whether you're fetching data from the database `(=> use Query())` or your manipulating the database / data in the database `(=> use Exec()).`

But what's the advantage of using `Prepare()?`

Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).

This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.

And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.

💡 But in order to show you the different ways of using the sql package,I decided to also include this preparation approach in this project 


✅  05-Beautify console and json data 🔆

- `go get github.com/fatih/color`

- use like this

import ("github.com/fatih/color")

color.Cyan("🔋 🚀 Server running at http://localhost:8080")

✅ 06-Add git ignore files