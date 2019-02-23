# restapi
basic REST API

## Build and run
```
(clone repo)
docker build --tag restapi .
docker run --name restapi -p 8080:8080 -d restapi
docker stop restapi
docker rm restapi
```

## Test
```
curl localhost:8080/fruit 
[{"type":"apple"}, {"type":"orange"}, {"type":"pear"}]
```

## Cleanup
```
(clone repo)
docker stop restapi
docker rm restapi
```
