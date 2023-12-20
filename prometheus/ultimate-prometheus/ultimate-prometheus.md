# Ultimate Prometheus

## The language of Prometheus

- Sample
    - Is a pair of timestamp + value (float)
- Time Series
    - List of samples ordered by time
- Time Series Database (TSDB)
    - Special type of DG used to store time series
    - Optimized toward querying by time
- Quantiles & Percentiles
    - Quantiles and percentiles are the same thing
    - Only difference is the representation (9oth = 0.9)
    - "90% of the time a request takes no more than 1s"
        - 0.9 quantile = 1
        - 90th percentile = 1
- Metric Names and Labels
    - metric_name{label1="value1", label2="value2", etc...}
- Job aand instance labels
    - Job is the service name
    - Instance is the node on which the service is running on
- Instant Vector
    - A list of samples where all samples contains the same timestamp
- Range Vector
    - A subset of samples from a time series in a given time range

## Features and Architecture
- Prometheus calls our app and not the other way around (pull)
- PromQL
    - Language to query time series im Prometheus
- Service Descovery
    - Prometheus can automatically find our services and monitor them
- Push Gateway
    - For short running jobs
    - Metrics can be pushed to the PushGateway
- Exporters
    - Export existing stats to Prometheus from 3rd party components of our system (e.g. MySQL, Kafka, Linux)
- Alert Manager
    - Prometheus can send alerts to the Alerts Manager, which in turn will send us a notification (slack, page, mail, etc)

## Type Of Metrics

- Counter
    - Type of metric that can only go up
- Gauge
    - Type of metric than can go both up and down
- Histogram
    - The purpose of a histogram is to estimate quantiles
    - It's mainly used with request latencies
- Summary
    - Serves the same goal as histogram
    - Quantiles are calculated in client side
        - It will move pressure from Prometheus to our apps
        - No ability to choose the time window

## PromQL (Prometheus Query Language)
