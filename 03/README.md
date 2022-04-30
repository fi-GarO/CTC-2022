
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
# Add a product - bere na vstupu .json file -> v repozitáři předem vytvořen .json file "body.json"
curl localhost:8080/products --include --header "Content-Type: application/json" -d @body.json --request "POST"
# Buy a product - odečítá -1 z produktu se zadaným id
curl localhost:8080/buy?id=2 --request "PATCH"
# Delete a product - smaže product se specifikovaným id
curl localhost:8080/products/delete?id=3 --request "DELETE"
