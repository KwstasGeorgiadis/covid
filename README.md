# covid
Stats for covid-19

# Quick set up

For macos and linux users

1.Set up log files  <br />
<pre> a. mkdir /var/log/covid  <br />
<pre> b. touch /var/log/covid/app.ndjson  <br />
 c. chmod -R 0777 /var/log/covid/app.ndjson  <br />
 Reminder that you can change the log's path by updating the config file  <br />
2. You need to run redis for this app to work  <br />
 a. Check https://redis.io to download  <br />
 b. redis-server  <br />
3.Build app  <br />
 a. go build app.go  <br />
 b. ./app  <br />

Feel free to import the postman collection in the directory ./postman  <br />
