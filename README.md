# team-project

**Preparation before running the web server**
-
**Download and install all resources**

install `Golang`, `Redis`, `PostgreSQL`, `Docker`, `docker-compose`

run
* `git clone git@gitlab.com:golang-lv-388/team-project.git`
* `git clone https://gitlab.com/golang-lv-388/team-project.git`
* `go get -d -v ./...`
* `go install -v ./...`
* `go get -u golang.org/x/vgo`
* `vgo build`

**Web server start:**
-
* go `run main.go`
> listen http://localhost:8080

**Run web server in a docker container:**
 * remove `go.sum` and `go.mode` files
 * run `vgo build`
 * run `docker-compose up --build`
> listen http://localhost:8080
