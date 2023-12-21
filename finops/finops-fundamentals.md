# FinOps

## Cloud Computing Fundamentals 

### Five essential characteristics, three service modesl, four deployment models

Characteristics:

- On demand self-service
- Broad network access
- Resource pooling
- Rapid elasticity
- Measured service

Service Models:
- Software as a Service (SaaS) - for users (Salesforce)
- Platform as a Servie (PaaS) - for developers (GCP App Engine)
- Infrastructure as a Service (IaaS) - for operations (EC2)

Deployment Models (how we use cloud):
- Private
- Hybrid
- Community
- Public

### On-prem vs Cloud

- Localization of the cost management (from central management to developers - more waste)
- Material spend: not possible to predict the cost of infrastructure
- Variable price: spend what you use. ("Every minute counts") 

### Cloud Challenges

- Security
- Performance
- Connectivity
- Billing

## FinOps Fundamentals 

FinOps is an evolving cloud financial management discipline and cultural practice that enables organizations to get maximum business value by helping engineering, finance, technology and business teams to collaborate on data-driven spending decisions.

- An efficient way to teams manage their cloud costs
- Ownership driven and supported by a central group
- Work towards to financial and operational control
- A way prioritize cloud costs optimization

FinOps is about ROI (profit)


Why FinOps:

- Control spending (control & optimize)
- Set best practices


1. Tagging
2. Optimize
3. Forecast consuption

FinOps appeared because people did not learnead to mamage cloud costs. So FinOps is a try of managing costs has on-prem resources style with budgeting approval.


### FinOps Lifecycle
1. Inform (Visibility and allocation)
2. Optimize (Utilization)
3. Operate
4. (Repeat)


## CAPEX and OPEX

- Capital Expenditures (long term investioment) (ownership of assets)
- Operational Expenditures (day to day costs) (paying for services)


## ROI and TOC

ROI: Return On Investiment (profit)

TCO: Total Cost Onwership - total cost of cloud technology


## Chargeback

- Is a policy of allocating resource usage costs to actual business units.
- Is a common method to avoid overspending.
- Provides visibility and accountability for resource spending (what was spent and who spent).
- Accounting model.
- Tagging (and other metadata) implementation


## FinOps Team

Should be responsible for:
- Defining enterprise cloud usage strategy;
- Definig and adjusting enterprise cloud budgets
- Setting up the cloud usage practices and guidelines
- Reviewing the data and determining corrective actions as needed.


Cloud costs should be visible to everyone => Showback

Establishing a team by identifying specific members from:
- Cloud Center of Excellence
- Sponsor(s)
- Engineering
- Finance
- Application owners
- DevOps
- And other business units

### Team Roles (= Personas)
- FinOps Practitioner
- Financial Controller
- Cloud Practitioners
- Executives
- Engineers


### Centralized FinOps Team
- Increases collaboration
- Increase culture
- Removes duplication of efforts
- Streamlines planning and decisions
- Removes bias and reduce disagreements
- Alignment with business goals

The level of centralization needs to be adapted to the organization.

### Organization Changes and adoption

1. Executive buy-in
2. Obtaining funding to establish a FinOps practice (people, training and tools)
3. assigning responsible members (key personas) to drive adoption

### Motivations

- FinOps practitioners must build reporting and processes that reduce the need for both finance and engineering teams (efficiency);
- Common language enables teams to familiarize themselves with a common set of terms that we understand (remove barriers);
- Every FinOps stakeholder will have specific motivations.


### FinOps Triangle
- Quality
- Speed
- Cost


## FinOps Capabilities
- FinOps domains represents a sphere of activity or knowledge;
- Each FinOps domain consists of FinOps capabilities;
- Capabilities are "functional activities" to meet the demands of a FinOps practice/domain.

### Domains
- Understand cloud usage & cost
- Performance tracking & benchmarking
- Real-time decision making
- Cloud rate optimization
- Cloud usage optimization
- Organizational alignment

#### Understand cloud usage & cost
- Cost allocation (metadat & hierarchy)
- Data analysis and showback
- Managing shared cost
- Data ingestion & normalization
- Managing anomalies
- Forecasting
- Measuring unit costs

#### Performance tracking & benchmarking
- Measuring unit costs
- Managing commitment based discounts
- Resource utilization and efficiency
- Forecasting
- Budget management
- Managing anomalies


#### Real-time decision making
- Cost allocation (metadata & hierarchy)
- Data analysis and showback
- Managing anomalies
- Managing commitment based discounts
- Resource utilization and efficiency
- Measuring unit costs
- Establishing a FinOps decision & Accountability structure


#### Cloud rate optimization
- Data analysis and showback
- Managing commitment based discounts


#### Cloud usage optimization
- Data analysis and showback
- Onboarding workloads
- Resource utilization and efficiency
- Workload management & automation


#### Organizational aligment (Policy & Governance)
- Establishing FinOps culture
- Managing shared costs
- Chargeback & IT Finance integration
- Data analysis and showback
- Budget management
- FinOps education & enablement
- Establishing a FinOps decision & Accountability structure
- Cloud policy & governance
- IT asset management integration

## 6 Principles of FinOps

### Collaboration - Teams need to colaborate
- Communication and collaboration
- Define governace and control
- Increase cost efficiency and performance 

### Business value - Decisions are driven by the business value of Cloud
- Maximize the value created by the spend (ROI)
- Business metrics


### Ownership - Everyone takes ownership of their cloude usage
- Responsability
- Chargeback
- Edge


### Timely reports - FinOps reports should be accessible and timely
- Real-time decisions
- Alerts and monitoring
- Cost avoidance
- True cost focus on clean decisions (What data? When data is collected?)
- FinOps decisions are based on fully loaded and properly allocated costs.


### Centralized team - A centralized team drives FinOps
- Drive value thru education, standardization, and cheerleading
- Cost centralization
- Benchmarking


### Variable cost model - Take advantage of the variable cost model of the Cloud
- Foreword looking (forecasting)
- Reserved instances
- Rightsizing
- Volume discounts

## Best Practices

### Inform
- Allocate
- Benchmark
- Forecast

### Optimize
- Compare the needs of the company with discounts for time and reserved usage in order to get the better rates.

### Operate
- Continuous evaluate the performance of cloud services againts business objectives for speend, quality and cost.

### Best practices
- Define organizational grouping and tagging
- Cloud as a business value generator
- Develop a common vocabulary
- Setup Kubernetes shared cost allocation models
- Forecast for a 3 - 12 month timeframe
- Develop scorecards aimed to comparing teams (showback and chargeback)
- Shift left cost accountability


## Aligning teams to Resources
- Resources: VMS, Services, Storage
- Teams: Business Units, defined stakeholders.

## FinOps Lifecycle
Crawl, Walk, Run
- Inform
    - Understand your organization resources comsuption thru context, data, budgets and forecasting
- Optimize
    - Setting your goals and tagerts by determining yor cloud costs thru strategy, tracking, cost breakdown and cost groth boundary
- Operate
    - Align organization teams to business gols thru drive and action. Focus on getting results for the organization.

### Inform
- Forecasting
- Visibility
- Allocation
- Benchmarking
- Budgeting


### Optimize
- Definig clear SLO
- Improve Reserved Instance usage
- Right sizing
- Eliminate waste
- Automate scaling, reporting


### Operate
- Define processes (onboarding, deployments, etc)
- Automate everything
- Evaluate team performance
- Establish new goals that improve speed, quality and performance
- Automate scaling

## Cloud economics and pricing
Is the study of cloud computing's costs and benefits and the economic principles that underpin them.
    - What's the return on investiment (ROI) of migrating to the cloud or switching current cloud providers?
    - What's the total cost of ownership (TCO) of a cloud solutionn versus a traditional on-prem solution?

## Capacity Reservations and Reserved Instances
- RI provides a "predictable" cost model
- With RIs, you commit in advance for usage which in return means a lower price.
- With reservations you can choose (mode upfront payment, the bigger discount)
    - Pay with no-upfront
    - Partial upfront
    - All upfront

On AWS:
- On-demand instances
- Reserved instances
- Spot instances
- Dedicated hosts

## Cloud Billing
- Process of generating bills from the resource usage data which is done via predefined billing policies.
- A central team is better equipped to manage and CUD commitments due to thei ability to look across the entire estate and the complexity of managing a large portfolio of commitments.
- A FinOps professional needs to be able to view and understand rwa data

Reducing your cloud bill:
- Aggregate the usage data and to identify waste
- Identify rightsizing opportunities and correct instance types
- Pay upfront when appropriate (usage discounts)
- Terminate resources not used with automation

### Cost allocation methods
- Tagging and labelling
- Hierarchal account separation 

### Tagging methods
- Rsource-level tag
- Accounts / projects / subscriptions
- Post-bill data constructs

### Benchmarking performance
Benchmarking (aka comparing) is important to compare performance
    - Ability to identify precisely dimensions, elements and processes that need benchmarking

## Cloud Spend Formula?
Spend = Usage x Rate

- AWS Cost Explorer is used to visualize costs after consumed
- AWS Budgets are used to forecasts costs and prevent overspending.


## AWS Pricing Models
- On-demand instances
- Reserved Instances
- Spot Instances
- Dedicated Hosts

### AWS Cloud Spend Management Tools
- AWS Trusted Advisor
- AWS Budgets
- Cost Explorer
- AWS Cost & Usage Reports
- Biling Alerts
- Billing Preferences

## Azure Pricing Models
- Pay-as-you-go
    - Provides for instant use with no commitment
- Reserved Instances
    - Provides for cost savings with commitment
- Spot Intances
    - Available Azure capacity at a significant discount for short term

### Azure Cost Management Tools
- Azure Price Calculator
- Azure Budgets
- Azure Advisor
- Azure Exports


## GCP Pricing Model

### Billing Discounts
- Sustained use disconts
- Committed use discounts
- Discounts for preemptible VM instances (similar to spot)


## Exam Preparation
- 1 hour
- 50 questions
- 3 Tries