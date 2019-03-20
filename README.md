# team-project

**Docker setup:**
- 
Compile app depend on OS 
 * `CGO_ENABLED=0 GOOS={OS} go build -a -installsuffix cgo -o {appname} .`
 * {OS} == `linux` , `windows` , `MacOS`
 * Building the image >>> `docker build -t {appname} .`
 * List all the available images >>> `docker image ls`
 * Running the Docker image >>> `docker run -d -p 8080:8080 {appname}`
 * Finding Running containers >>> `docker container ls`
 
Interacting with the app
 running inside the container >>>
 
 `curl http://localhost:8080?name=DockerWorl`
 > `Hello, DockerWorl`
 #