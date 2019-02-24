# restapi
Follow the steps below to use Docker to build and run a basic REST API service written in golang.

1. clone this repo

2. build the docker image
```
docker build --tag restapi-image .
```

3. run the container
```
docker run --name restapi -p 8080:8080 -d restapi-image
```

4. list the containers to verify the container is running (should see an entry with the image restapi-image and name restapi)
```
docker ps
```

5. test the service, list the fruit
```
curl localhost:8080/api/fruit 
```

returns:
```
[{"id":"0","type":"apple"},{"id":"1","type":"orange"},{"id":"2","type":"pear"}]
```
6. add a fruit
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"type":"banana"}' \
  http://localhost:8080/api/fruit
curl localhost:8080/api/fruit 
```
returns:
```
[{"id":"0","type":"apple"},{"id":"1","type":"orange"},{"id":"2","type":"pear"},{"id":"3","type":"banana"}]
```

7. delete a fruit
```
curl --header "Content-Type: application/json" \
  --request DELETE \
  http://localhost:8080/api/fruit/3
curl localhost:8080/api/fruit 
```
returns:
```
[{"id":"0","type":"apple"},{"id":"1","type":"orange"},{"id":"2","type":"pear"}]
```

8. cleanup the container and image
```
docker stop restapi
docker rm restapi
docker rmi restapi-image
```
