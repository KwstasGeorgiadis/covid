# covid
Stats for covid-19

# Quick set up

## For macos and linux users

1. Set up log files  <br />
&nbsp; a. mkdir /var/log/covid  <br />
&nbsp; b. touch /var/log/covid/app.ndjson  <br />
&nbsp; c. chmod -R 0777 /var/log/covid/app.ndjson  <br />
&nbsp; Reminder that you can change the log's path by updating the config file  <br />
2. You need to run redis for this app to work  <br />
&nbsp; a. Check https://redis.io to download  it<br />
&nbsp; b. ``bash
redis-server
```
3. Build app  <br />
&nbsp; a. ```bash
go build app.go
``` <br />
&nbsp; b. ```bash
./app
``` <br />

Feel free to import the postman collection in the directory ./postman  <br />

## Docker

```bash
docker-compose up --build
```

Or you can use curl request like this one
&nbsp; curl --location --request GET 'localhost:9080/countries'
