# team-project

**Preparation before running the web server**
-
**Download and install all resources**

install `Go`, `Redis`, `PostgreSQL`, `Docker`, `docker-compose`

run
* ssh `git clone git@gitlab.com:golang-lv-388/team-project.git`
* https `git clone https://gitlab.com/golang-lv-388/team-project.git`
* `go build`

**Server address**

> https://team-projectv1.herokuapp.com/

**Web server run local start:**
-
>For running local make changes to the file `database/database.go`
>Change second argument in method `sql.Open` to `fmt.Sprintf("host=%s port=%s 
user=%s "+
"password=%s dbname=%s sslmode=disable", configurations.Config.PgHost, 
configurations.Config.PgPort, configurations.Config.PgUser,
configurations.Config.PgPassword, configurations.Config.PgName)`
> In the `main.go` file reassign the variable `port` to **string** `"8080"` 
>
* `go run main.go`
> listen http://localhost:8080

**Run web server in a docker container:**
 * run `docker-compose up --build`
> listen http://localhost:8080