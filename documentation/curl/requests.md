# covid
CURL requests for covid api

# Quick set up

* ```curl --location --request GET 'localhost:9080/api/csse/USA' --header 'Content-Type: application/json'``` for endpoint /api/continent``` for endpoint /api/csse/USA
* ```curl --location --request GET 'localhost:9080/api/countries/all' --header 'Content-Type: application/json'``` for endpoint /api/continent``` for endpoint /api/countries/all
* ```curl --location --request POST 'localhost:9080/api/compare/all'  --header 'Content-Type: application/javascript'  --data-raw '{ "countryOne" : "Spain", "countryTwo" : "Italy"}'''``` for endpoint /api/countries/all
* ```curl --location --request GET 'localhost:9080/api/total' --header 'Content-Type: application/json'``` for endpoint /api/continent``` for endpoint /api/total
* ```curl --location --request POST 'localhost:9080/api/sort' --header 'Content-Type: application/json' --data-raw '{"type" : "deaths"}'``` for endpoint /api/sort
* ```curl --location --request GET 'localhost:9080/api/countries' --header 'Content-Type: application/json'``` for endpoint /api/continent``` for endpoint /api/countries
* ```curl --location --request POST 'localhost:9080/api/country' --header 'Content-Type: application/json' --data-raw '{ "country" : "USA"}''``` for endpoint /api/country
* ```curl --location --request GET 'localhost:9080/api/news/all' --header 'Content-Type: application/json'``` for endpoint /api/continent``` for endpoint /api/news/all
* ```curl --location --request GET 'localhost:9080/api/hotspot/12' --header 'Content-Type: application/json'``` for endpoint /api/continent``` for endpoint /api/hotspot
* ```curl --location --request GET 'localhost:9080/api/continent' --header 'Content-Type: application/json'``` for endpoint /api/continent
* ```curl --location --request GET 'localhost:9080/api/world'``` for endpoint /api/world
