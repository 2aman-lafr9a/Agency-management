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

# to create a container

```shell
docker run -p 8080:8080 -it -d --name agency-management-container agency-management
```

```shell

```

### using docker composer 
or better you can use this command that will create all for you 

```shell
docker compose up -d
```

### 