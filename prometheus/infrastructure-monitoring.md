# Infrastructure Monitoring with Prometheus

## The value of monitoring
- With the growing complexity of infrastructures, it has become critical to attain a global view of all different components of and infrastructure.
- You can look at monitoring as a source of information for maintaining healthy systems, production-wise and business-wise.

## Monitoring Components
- Metrics
    - Exposes certain system resource, application action, or business characteristic as a specific point in time value;
- Logging
    - Containing much more data than a metric, this manisfests itself as an event from a system or application, containing all the infromation that's produced by such event;
- Tracing
    - This is a special case of logging where a request is given a unique identifier so that can be tracked during its entire lifecycle across every system;
- Alerting
    - This is the continuous threshold validation of metrics or logs, and fires an action or notification in the case of a transgression of said threshold;
- Visalization
    - This is a graphical representation of metrics, logs, or traces.

## Monitoring Categories
- Whitebox
    - The system under observation surfaces data about its internal state and the performance of critical sections.
    - The telemetry data is usually handled in the following ways:
        - Exported through loggin
        - Emmited as structured events
        - Maintained in memory as aggregates
- Blackbox (external probing)
    - The application or hosts is observed from the outside and, consequently, this approach can be fairly limited.
    - You last line of defense - you can rely on blackbox monitoring to assess availability.

## Metrics Collection
Two apporaches:
- Push
- Pull (Prometheus)

## What to measure
Google four golden signs:
- Latency
    - The time required to serve a request
- Traffic
    - The number of requestes being made
- Errors
    - The rate of failing requests
- Saturation
    - The amount of work not being processed, which is usually queued.

RED method:
- Rate
    - Translated as requests per second
- Errors
    - The amount of failing requests per second
- Duration
    - The time taken by those requests


## Prometheus Ecosystem
- Prometheus is a time series-based , open-source monitoring system.
- It collects data by sending HTTP requests for hosts and services on metric endpoints, which it then makes available for analysis and alerting using a powerful query language.

- Prometheus server
- Alertmanager
- Exporters
- Pushgateway
- Visualization 


### Exporter
An exporter is nothing more than a piece of software that collects data from a service or application and exposes it via HTTP in the Prometheus format.


## Metrics Foundamentals
- Metrics are the core resources

### Time Series Data
Time series data can usualy be defined as a sequence of numerical data points that are indexed chronologically from the same source.
 
A time series database store the follwoing components:

- A timestamp
- A value
- Some context about the value, encoded in a metric name or in associated key/value pairs.

Example:

timestamp=123456, Company=ACME, location=headquarters, beverage=coffee, value=4012

### Samples
Samples are the collected data pints, and they represent the numerical value of time series data. 

### Cardinality
The term is often used to mean the number of unique time series that are produced by a combination of metric names and their associated label names/values.

### Core Metric Types
- Counters
- Gauges
- Histograms
- Summaries


#### Counter
This is a strictly cumulative metric whose value can only increase.

#### Gauge
Is a metric that snapshots a given measurement at the time of collection, which can increase or decrease (temperature, disk space, memory usage)

#### Histogram
Histograms allow you to retain some granularity by counting events into buckets that are configurable on the client side, and also by providing a sum of all obeserved values.


#### Summaries
Summaies are similar to histograms in some ways, but present different trade-offs and are generally less useful.


### Aggregations
- Longitudinal 
- Cross-sectional

Aggregation is the process that reduxes or summarizes the raw data, which is to say that it receives a set of data points as input and produces a smaller set as output.


## Run a Prometheus server
There are two main types of configuration on a Prometheus server:
- command-line flags
- configuration files (runtime configuration)