
# Go REST Api

### Assignment

Product management API

1. Create golang server app that provides API for managing list of products
   1. Product(name, price, amount)
   2. Operations - get, list, update, delete
   4. Backend - SQL or NOSQL database (Redis, Mongo, Cassandra)
   3. Include simple testing client in your repository

# Get all products - vrací aktuální seznam produktů
curl localhost:8080/products
# Add a product - bere na vstupu .json file -> v repozitáři předem vytvořen .json file "data1.json" a "data2.json"
curl localhost:8080/products --include --header "Content-Type: application/json" -d @data1.json --request "POST"
# Update a product - aktualizuje product s definovaným id (localhost:8080/products/1) podle .json souboru "patch.json"
curl localhost:8080/products/2 --include --header "Content-Type: application/json" -d @patch.json --request "PATCH"
# Delete a product - smaže product se specifikovaným id
curl localhost:8080/products/1 --request "DELETE"
# Delete all products
curl localhost:8080/products --request "DELETE"