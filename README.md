### using docker commands

## to build the image

```shell
docker build -t agency-management .
docker run --name postgres  \
    -e POSTGRES_PASSWORD=password \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v /data:/var/lib/postgresql/data \
    -p 5432:5432 -d postgres
```


### using docker composer 
or better you can use this command that will create all for you 

```shell
docker compose up -d
```

### 


```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
```
