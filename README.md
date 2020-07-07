# covid
Stats for covid-19

# Quick set up

## For macos and linux users

1. Set up log files  
 a. ```mkdir /var/log/covid``` \
 b. ```touch /var/log/covid/app.ndjson``` \
 c. ```chmod -R 0777 /var/log/covid/app.ndjson```. You can change the log's path by updating the config file.
2. You need to run redis for this app to work \
 a. Check https://redis.io to download  it \
 b. Run command ```redis-server```
3. Build app \
 a. ```go build app.go``` \
 b. ```./app```


## Docker

```bash docker-compose up --build ```

# Test it

Feel free to import the postman collection in the directory ./postman

Or you can use curl request like this one \
  ```curl --location --request GET 'localhost:9080/api/countries'```
