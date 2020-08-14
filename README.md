# gomicrosvc-k8s

######################################################

docker build -t gomicrosvc-k8s:latest .

docker run -it -p 8080:8080 gomicrosvc-k8s:latest

######################################################

##### 1. 
curl --location --request POST 'http://localhost:8080/echo'
##### 2.
curl --location --request GET 'http://localhost:8080/api/books/'
##### 3.
curl --location --request POST 'http://localhost:8080/api/books/' \
--header 'Content-Type: application/json' \
--data-raw '{"title":"Amazing","author":"Live","isbn":"123-456-234"}'
