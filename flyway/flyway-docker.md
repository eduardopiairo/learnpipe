# Flyway Docker

## Run
```
docker run --rm --net="host" flyway/flyway -user=sa -password=Strong@Passw0rd  -url="jdbc:sqlserver://localhost;databaseName=TestDB;encrypt=true;trustServerCertificate=true" info
```

```
docker run --rm flyway/flyway -user=sa -password=Strong@Passw0rd -url="jdbc:sqlserver://0.0.0.0;databaseName=TestDB" info
```