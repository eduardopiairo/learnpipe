# Observability with Grafana, Prometheus,Loki, Alloy and Tempo

## Foundations of Observability

Why are CI/CD and microservices relevant to observability?

1. Many services to monitor.
2. Intra-service communications can fail.
3. More vulnerable to security threats.
4. More changes are deployed.

### Monitoring

Is collecting and visualising data about systems regurly so the system health can be viewed and tracked.

Three questions:
1. Is the service on?
2. Is the service functioning as expected?
3. Is the service performing well?

The data collected for monitoring is called telemetry data.

Monitoring methods:
- Core Web Vitals (UI Layer)
- RED (service Layer) (Rate, Errors, Duration)
- USE (Infrastructure Layer) (Utilization, Saturation, Errors)
- Four golden signals ( Service and Infrastructure Layers) (Latency, Traffic, Errors, Saturation)

### Observability

What is Observability?

- Monitoring is part of observability.
- Observability is gathering actionable data in a way that gives an holistic view of the entire system, abd tells us where and why an issue occurs.


### Types of Telemetry Data

MELT: 
- M = Metrics
- E = Event
- L = Log
- T = Traces

## Methods of collecting metrics

Push: applications and services send metrics to an endpoint, via TCP, UDP or HTTP.

Scrape: applications and services provide APIs for the times series database to read the metrics.

Push vs Scrape
- Think about the type os systems and applications;
- Think about scalability
- Think about complexity

## Install Prometheus

Install
```
brew install prometheus
```

Run the Prometheus service
```
brew services start prometheus
```

### Data Collection

- Exporter
- Push Gateway

### Node Exporter

Node = Every UNIX-based kernel