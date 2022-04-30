
# Docker compose

### Assignment

1. Bundle application from assignment 3 with database as docker compose file
   1. All images must be public 
2. User volumes for data storage
3. Use separate docker network bridge


### Example way hot to build container containing Go binary

1. (Windows Only) [Cross compile to linux](https://stackoverflow.com/a/43945772)
   ```shell
   set GOARCH=amd64
   set GOOS=linux
   ```
1. Build
   ```shell
   go build -o myapp
   ```
2. Dockerfile
   ```Dockerfile
   FROM alpine:3.15.4
   COPY myapp myapp
   ENTRYPOINT ["./myapp"]
   ```
3. Build image and push
   ```shell
   docker build -t mydockeraccount/myapp:1.0.0 .
   docker push mydockeraccount/myapp:1.0.0
   ```

