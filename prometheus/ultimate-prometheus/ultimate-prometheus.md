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

### Arithmetic Operators
- Addition (+)
- Subtraction (-)
- Multiplication (*)
- Division (/)
- Modulo (%)
- Power (^)

### Binary Operators
- And
- Or
- Unless

### Aggregation Operators
- max(M)
- min(M)
- sum(M)
- count(M)
- avg(M)
- quantile (0.5, M)
- count_values("abc", M)
- topk(2, M)
- bottomk(2, M)

### Query Functions
- abs(v instant-vector)
    - Coverts values to their absolute form
- absent(v instant-vector)
    - Returns 1 if the given vector doesn't exist
- absent_over_time(v range-vector)
    - Same as absent but for a range
- ceil(v instant-vector)
    - Rounds sample values up
- floor(v instant-vector)
    - Rounds sample values down
- clamp_max(v instant-vector, max scalar)
    - Return sample values with an upper limit
- clamp_min(v instant-vector, max scalar)
    - Return sample values with a lower limit
- changes(v range-vector)
    - The number of times a vector changed in the given time
- delta(v range-vector)
    - The difference between the first and the last sample values
- idelta(v range-vector)
    - The difference between the last two sample values
- increase(v range-vector)
    - The difference between the first and the last sample values
    - To be used with counters because it can handle counter resets
- rate(v range-vector)
    - increase(v) / number of seconds in time range
- irate(v range-vector)
    - idelta(v) / time delta between last two samples
- label_join
    - Same as joining a string from a string array, but with labels
- label_replace
    - Used to replace a label's value or create a new label
- resets(v range-vector)
    - The number of times a counter was reset
- sort(v instant-vector)
    - Sorts vector by element's sample value
- sort_desc(v instant-vector)
    - Sorts vector by element's sample value in descending order
- timestamp(v instant-vector)
    - Returns the samples time for each time series in the vector
- scalar(v instant-vector)
    - Returns the value for a single elemnt vector
- vector(s scalar)
    - Wraps a scalar in a single element instant-vector
- histogram_quantile(q float, v instant-vector)
    - Calculates the q quantile on a histogram


Date and time functions:
- day_of_month(M)
- day_of_week(M)
- days_in_month(M)
- hour(M)
- minute(M)
- month(M)
- year(M)


Agregation over time
- max_over_time(M[30s])
- min_over_time(M[30s])
- sum_over_time(M[30s])
- count_over_time(M[30s])
- avg_over_time(M[30s])

## AlertManager