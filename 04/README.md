
# Docker compose


# Get all products - vrací aktuální seznam produktů
curl localhost:8080/products
# Add a product - bere na vstupu .json file -> v repozitáři předem vytvořen .json file "data1.json" a "data2.json" (při spouštění commandu je nutno být ve složce s /data1.json)
curl localhost:8080/products --include --header "Content-Type: application/json" -d @data1.json --request "POST"
# Update a product - aktualizuje product s definovaným id (localhost:8080/products/1) podle .json souboru "patch.json"
curl localhost:8080/products/1 --include --header "Content-Type: application/json" -d @patch.json --request "PATCH"
# Delete a product - smaže product se specifikovaným id
curl localhost:8080/products/1 --request "DELETE"
# Delete all products
curl localhost:8080/products --request "DELETE"


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

