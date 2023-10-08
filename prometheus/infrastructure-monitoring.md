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

Google four folden signs:
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