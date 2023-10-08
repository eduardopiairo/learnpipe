# PostgreSQL


## Get the image

```
docker pull postgres:15.1
```

## Run Postgres instance

```
docker run --name postgres -e POSTGRES_PASSWORD=supersecretpass -d postgres:15.1
```


docker run --name myPostgresDb -p 5455:5432 -e POSTGRES_PASSWORD=postgresPW -d postgres:15.1