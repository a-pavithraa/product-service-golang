# Project Name
Products Service
## Description

POC in Golang to check KeyCloak and Migrate usage


## Usage


1. Clone the repository.
2. Spin up KeyCloak and PostgreSQL containers by using `docker-compose up -d`.
3. Start the application by using `go run main.go`.
4. Import the Postman collection from Products.postman_collection.json

## Additional Details

1. KeyCloak Realm is imported on startup. 
2. Table creation and populating inital records for the table are taken care by Migrate. If any changes need to be done in schema, add a new migration file in db/migrations folder following the convention mentioned in https://github.com/golang-migrate/migrate

