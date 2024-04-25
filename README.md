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

KeyCloak Realm is imported on startup. Intial schema and records for Database are created using Migrate

