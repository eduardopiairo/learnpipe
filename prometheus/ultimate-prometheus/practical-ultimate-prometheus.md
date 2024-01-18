# Practical side of Ultimate Prometheus

## Prometheus

- Installation directory: /Users/eduardopiairo/playground/prometheus/prometheus-2.47.2

Run prometheus

```
./prometheus --config.file=prometheus.yml 
```

Go to localhost:9090


## Go Application
```
go run . 
```

## AlertManager
```
./alertmanager --config.file=alertmanager.yml 
```

### Configure AlertManager and Prometheus

1. Create new alert file

```
cd prometheus

touch alert-rules.yml
```

