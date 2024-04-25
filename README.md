# Project Name
Products Service
## Description

POC in Golang to check KeyCloak and Migrate usage


## Usage


1. Clone the repository.
2. Spin up KeyCloak and PostgreSQL containers by using `docker-compose up -d`.
3. Start the application by using `go run main.go`.
4. Import the Postman collection from Products.postman_collection.json.
5. In the collection, click on Get New Access token. When asked for credentials, provide admin/admin for admin user and user1/user1 for regular user. Click on Use token once authenticated.
6. All the endpoints can be accessed only when step 5 is done. Otherwise 401 error is thrown.
7. Only the admin user has access to POST/Product endpoint

## Additional Details

1. KeyCloak Realm is imported on startup. 
2. Table creation and populating inital records for the table are taken care by Migrate. If any changes need to be done in schema, add a new migration file in db/migrations folder following the convention mentioned in https://github.com/golang-migrate/migrate

