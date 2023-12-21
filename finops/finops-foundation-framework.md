# FinOps Foundation Framework

The FinOps Framework provides the operating model for how to establish and excel in the practice of FinOps.

- [x] FinOps Domains
- [x] FinOps Capabilities
- [x] FinOps Maturity Model
- [x] FinOps Personas
- [x] FinOps Phases
- [x] FinOps Principles

## FinOps Domains
FinOps domains represents a sphere of activity or knowledge. Each domain consists of FinOps capabilities. Capabilities outline functional activities that can be performed as part of that domain.

- A set of capabilities.
- Provide a high level overview of what functional activities are needed to run a FinOps practice.

The domains are:
- Understanding cloud usage and cost
- Performance traking & benchmarking
- Real-time decision making
- Cloud rate optimization
- Cloud usage optimization
- Organizational alignment


### Understanding cloud usage and cost
Within this domain, the organization will work to gather all required information about its cloud usage and cost, normalize it, make it available for review.

This domain enables an organization to understand what cloud services it’s using, what is driving spend, and who owns that spending.

*This domain answers the question: What are we spending on cloud and what are we using?*

The capabilities within this domain focus on understanding and allocating the cost and usage within the organization, and not on actioning it for a particular purpose:
- Measuring unit costs
- Managing shared costs
- Managing anomalies
- Forecasting
- Data ingestion & normalization
- Cost allocation (meatadata and hierarchy)
- Data analysis and showback


### Performance traking & benchmarking
Within this domain, the organization sets and maps its usage and cost to budgets, uses historical information to forecast, and establishes and measures KPIs and other performance indicators, including benchmarking.

*This domain answers the question: Does what we’re using/spending allow us to achieve our strategic and organizational objectives?*

This domain entails the capabilities that look at the past and current spend, set baselines and budgets, and then forecast:
- Resource utilization & efficiency
- Measuring unit costs
- Managing commitment based discounts
- Managing anomalies
- Forecasting
- Budget management


### Real-time decision making
When we understand what we are spending, and understand how we are performing relative to expectations and standards, we can then use that information to make real-time decisions as we receive new information.

The goal here is not to specifically strive for real-time data in and of itself, but to leverage timely access to consistent cost & usage data for making continuous adjustments.

*This domain answers the question: What actions can I take now to allow me to better meet my organizational goals and objectives?*

The capabilities for this domain are:
- Measuring unit costs
- Managing anomalies
- Establishing a FinOps decision and accountability structure
- Data analysis and showback


### Cloud rate optimization
Within this domain, the organization works to define its pricing model goals, uses historical data to make pricing model adjustments by buying commitment based discounts, and works to manage the pricing aspects of services it is using in the cloud.

*This domain answers the question: How can we change what we’re paying and how we’re buying what we use in the cloud to achieve better price performance?*

This domain will contain the specialized capabilities to improve the manner in which we purchase cloud services, and ensure our pricing models, purchase options, and committed use are consistent with our goals:
- Intersection of cloud FinOps and sustainability
- Managing commitment based discounts
- Data analysis and showback


### Cloud usage optimization
Within this domain, the organization identifies and takes action to match running cloud resources to the actual demand of the workloads running at any given time (rightsizing,turn-off, delete resources). 

*This domain answers the question: How can we change what we’re using and when we’re using it in the cloud to better meet our organizational goals?*

This domain contains the set of capabilities to use the right resources, in the right size, only when we need them to produce business value:
- Intersection of cloud FinOps and sustainability
- Workload management and automation
- Resource utilization and efficiency
- Onboarding workloads
- Data analysis and showback


### Organizational alignment
Within this domain, the organization acts and automates to manage cloud use within the context of other IT Finance activities, and integrates FinOps capabilities with existing organizational processes, organizational units, and technology.

*This domain answers the question: What changes can I make within my organization to use cloud more effectively?*

This domain contains the capabilities we’ll use to continuously improve, to change and align our organization itself – its people, processes and technology:
- Cloud policy and governance
- Managing shared costs
- Managing commitment based discounts
- Establising FinOps culture
- FinOps education and enablement
- Establising a FinOps decision and accountability structure
- Chargeback and finance integration
- Budget management
- FinOps and intersecting frameworks (ITSM, ITFM/TBM, ITAM/SAM, Security, Sustainability) 


## FinOps Capabilities
FinOps capabilities represent functional areas of activity in support of their corresponding domain.
- Tasks or processes that support the FinOps practice through the FinOps phases


### 1. Cost allocation (Metadata & Hierarchy)

Cost Allocation is the set of practices to divide up a consolidated invoice or bill among those who are responsible for its various component parts.

Cost Allocation is done through a combination of functional activities primarily focused around the use of a consistent hierarchy of accounts, projects, subscriptions, resource groups and other logical groupings of resources; along with resource-level metadata – tags or labels – applied within the cloud service provider.

Cloud provider invoices can be analyzed, using a combination of hierarchy and metadata to allocate costs to various organizational, functional, or financial targets within the user company. 
- Any resources or hierarchy groupings which have no metadata attached can be addressed by the FinOps team as a compliance issue.
- Resources and hierarchy groupings which can be directly attributed to an internal target can be mapped to that target.
- Resources and hierarchy groups which are shared, can be identified as shared cost, and then various models of distributing those costs can be applied.

An important prerequisite to effective cost allocation is the metadata strategy, which should include all of the details governing the cost allocation process.


#### As someone in a FinOps Practitioner role, I will…
- develop the Naming Standards for all required and optional layers of hierarchical groupings
- develop compliance standards for various groups


### 2. Data analysis and showback

Data analysis and showback is the ability to leverage data, along with metadata on cloud resources and resource hierarchies, to create a near “real time” reporting mechanism for stakeholders.

In the context of FinOps, this work will typically focus on the cloud cost and usage data. This Capability relies heavily upon adequate data ingestion and data normalization capabilities.

It is within the scope of this Capability to create the data repository of normalized, queryable data from which reporting, analysis, and visualization of cloud cost and usage will occur.

#### As a FinOps Practitioner, I will…
- spend a significant amount of effort in collaboration with Engineering/Finance/Procurement/Product Management to build a cost reporting ecosystem which is aimed at helping consumers understand the important aspects of their spend, as well as, opportunities to optimize their spend.


### 3. Managing anomalies
Anomaly Management is the ability to detect, identify, clarify, alert and manage unexpected or unforecasted cloud cost events in a timely manner, in order to minimize detrimental impact to the business, cost or otherwise.

In the context of Cloud FinOps, anomalies are levels of spending that are different from the normal or expected spend.

Resolving Anomalies typically involves some level of investigation and then either a change to adjust the environment, or to adjust the expectation of the cost of the affected scope. Another resolution may be to simply acknowledge the anomaly.

#### As someone in a FinOps Practitioner role, I will…
- establish requirements for anomaly detection tool selection that is appropriate for cost monitoring and capable of defining, refining, detecting and alerting unexpected cloud spending events (cost anomalies).
- work with stakeholder teams to establish anomalous detection thresholds and reporting/notifications frequency
- document and communicate anomaly detection mechanism and thresholds to all stakeholders
- ensure that anomaly detection is tied appropriately to cost allocation metadata
- generate reports that surface all and/or alerted anomalous spending


### 4. Managing shared cost
“Everyone takes ownership for their cloud usage”. The true key to understanding total cost of ownership is built upon transparency and accuracy, but unallocated shared costs hinders both of these. Without appropriately splitting costs that are shared, engineers and product managers lack a complete picture of how much their products are really costing.

The goal of cost sharing can be:
- full allocation; 
- however it can also be the adoption of an informed ignore approach. 

Almost every organization has cloud costs that need to be segmented and allocated to business departments; the cost of running networking services, support costs which benefit all Engineering teams or simply the cost of service resources from a cloud provider for which they do not (yet) support tagging/labelling.

#### As someone in a Finance/FinOps role, I will…
- understand the basics of how cloud computing works, know the key services offered by cloud providers being used, including their common use cases, and have a basic understanding of billing and pricing models.
- work with Finance to ensure required allocations are implemented to support financial reporting so shared costs can be assign to specific business owners.
- build processes and define rules to enable full allocation of shared costs.
- allocate the shared costs based on the defined rules and create stakeholder persona reporting
- understand available cost sharing models (fixed, proportional, even-split) and determine which models should be applied based on the use case.


### 5. Forecasting
Forecasting is the practice of predicting future spending, usually based on a combination of historical spending and an evaluation of future plans, understanding how future cloud infrastructure and application lifecycle changes may impact current budgets and influence budget planning and future cloud investment decisions.

This capability also involves collaboration between stakeholder teams like Finance, Engineering, and Executives to build agreed upon forecast models and KPIs from which to establish budgets that align with business goals.

Cloud spend is variable which is inherently difficult to predict. Specifically engineers can start workloads at any time typically without having to go through a procurement process.

Accurate financial forecasting depends on the provision of accurate data as input.

#### Forescasting methodologies
- Trend-Based	
    - Uses historic trends to forecast future spend. 
    - Ideally this takes seasonality into consideration. Seasonality can include annual peaks during holidays but also daily peaks    when more people are using a service during specific hours of the day.
- Driver-Based	
    - Uses Key Performance Indicators (KPIs) to forecast the effect on business results. 
    - KPIs can be things like active accounts, widgets sold, ad impressions and so forth.
- Rolling Forecast	
    - To predict next month, quarter, and year. It allows companies to adjust their plans based on any shifts in the business such as economic changes, COVID as an example. 
- Static Forecast	
    - Predict for the fiscal year only with no adjustments.
- Special Projects	
    - Are planned cloud workloads that currently do not yet exist in the cloud. 
    - Their cost needs to be estimated by engineers and layered into trend or driver based forecasting to get a complete picture of future cloud spend. 
    - Special projects can also be costs that will not materialize on the cloud bill like licensing fees, professional services, or small workloads running on other cloud providers where automation isn’t feasible.

#### As someone in a Finance/FinOps role, I will…
- establish requirements for when forecasting is due and how frequently forecast updates are needed
- generate granular forecasts with reasonable accuracy
- help to establish forecasting KPIs that are in alignment with business goals
- explore optimization opportunities with teams that are forecast to overspend
- provide forecast data for cloud costs to enable stakeholders to create budgets
- provide granular reporting to teams on forecasted spend by different business-centric dimensions
- provide reporting on budgets vs actuals vs forecast to establish trends and compare against variance KPIs


### 6. Budget management
Budgeting for Cloud is a process of collecting estimated expenses for a specific period of time. Decisions on how to operate as a business, what to invest in and other strategic decisions are made based on budgets.

- The term ‘favorable to budget’ means that there are less expenses than as planned in the budget.
- The term ‘unfavorable to budget’ means that there are more expenses than as planned in the budget.

It may sound like being favorable to budget is ideal, and while it is not necessarily a bad thing, it is far from ideal. The more accurate the budget is, the better it is.. 

Budget Management contains the acts of:
- Accumulating the estimated expenses for a specific time period.
- Tracking actuals and forecasting remaining spend and comparing to the budgeted amount to identify material variances to budget that are favorable or unfavorable.
- Investigating the causes of variances to budget.

### As someone in a Finance/FinOps role, I will…
- Work with budget leaders in budget planning
- Help monitor cloud spend vs budget
- Provide reporting on budget vs actual vs forecast
- Alert teams who are projected to materially overspend or underspend
- Work with teams who are forecast to overspend on possible optimisations
- Communicate to appropriate leadership when material underspend to budget is expected so that they can make appropriate decisions


### 7. Workload management & automation
Workload management and Automation focuses on running resources only when they are needed, and creating the mechanisms to automatically adjust what resources are running at any given time.

This Capability is intended to give FinOps teams the ability to match supply to demand most efficiently, and effectively optimize cloud usage through measurement of workload demand and provisioning capacity dynamically.

#### As a FinOps Practitioner, I will:
- work with application owner to assign the missing tags to the resources
- automate the communication of the statistics of the affected resources by the automation, non-tags resources


### 8. Managing commitment based discounts
Cloud services have different approaches that leverage spend commitment to offer discounts on services. These vary from customized commercially negotiated discounts, to spend-based commitment discounts like AWS Savings Plans, resource-based commitment discounts like Google CUDs and others.

Spend-based commitment discounts and resource-based commitment discounts are the most popular rate optimizations that cloud service providers offer. This is partially because CSP native tooling and FinOps platforms enable you to plan, manage, and benefit from these types of discount constructs.

Altogether the implementation of these strategies drives an organizations Effective Savings Rate (ESR). It is important to note that under utilization of a commitment based discount would also negatively impact ESR as would significant usage not covered by discounts.


### 9. Resource utilization & efficiency
Resource utilization is about ensuring there is sufficient business value for the cloud costs associated with each class or type of resource being consumed. It is necessary to observe a resource’s utilization over time to understand if the performance, availability or other quality metrics are of value for the expense incurred.

For compute resources, there may be times when it is deemed that for performance or availability gains, average utilization may need to decrease and the extra expense incurred is worth the value creation the resource provides. Or the opposite may be true and performance expectations can be lowered to improve cost. For these decisions to be made, resource utilization, efficiency and cost must be looked at together.

By comparison, for storage resources, it is necessary to estimate the latent inefficiency in the stored data, and by extension the potential gross savings that can be realized by removing, or rightsizing, that inefficiency. Keep in mind that different data sets have unique latent inefficiencies and require tailored approaches.

#### As someone in a Finance/FinOps role, I will…
- highlight any opportunities to increase utilization and efficiency and work with the teams to review feasibility of alternative options;
- help create the reporting to track and report on the impact on value of underutilization and inefficiencies;
- partner with the Engineering organization to establish budgetary & efficiency targets.


### 10. Measuring unit costs
This capability is about developing metrics that reveal the business value of your cloud spend.

When practitioners address measuring unit costs, it’s often in the context of Cloud Unit Economics. Our practitioners define Cloud Unit Economics as a system of profit maximization based on objective measurements of how well your organization is performing against not only its FinOps goals, but as a business overall. 

Cloud Unit Economics achieves these goals by leveraging the measurement of marginal cost (a.k.a., unit cost metrics) specific to the development and delivery of cloud-based software and marginal revenue (a.k.a., unit revenue metrics).

By calculating the difference between marginal cost and marginal revenue, practitioners can determine where cloud operations break even and begin to generate a profit.


### 11. Data ingestion & normalization
Data ingestion and normalization in the context of FinOps represents the set of functional activities involved with processing/transforming data sets to create a queryable common repository for your cloud cost management needs.

 In this context, data ingestion and normalization occurs when bringing together cloud billing data, cloud usage data, cloud utilization and performance data, on-premises CMDB or ITAM data, business-specific data, and other data points from a variety of cloud providers and IT data repositories to create a collection of cost and usage information which can be queried to support and enable the FinOps cpabilities.

An effective data ingestion and normalization strategy should strive to provide the FinOps team with the right combination of data from:
- the right source systems
- at the right timeliness to support the cadence of decision making
- at the right level of granularity to support aggregate reporting and drill down investigation
- with the appropriate standardization, augmentation, normalization, etc.

The strategy for data ingestion will be driven largely by the needs of the reporting, cost allocation and optimization reporting needs. Data required to make decisions at a business unit level will not be in need of as much detailed or granular information.

FinOps teams which manage or allocate costs at a resource level may require multiple sources of data to gather resource information for some cloud providers which don’t natively provide it.

#### As someone in the FinOps team role, I will…
- determine the list of data sources required to fulfill my current reporting and operational needs;
- determine the level of granularity required in each data source;
- establish a data model for normalization, mapping fields from various sources to one another;
- regularly and proactively validate data source content, and clearly understand when changes occur, react to them, adjust and re-document accordingly, and notify all those affected;
- ensure that the data sources and resulting repository of cost and usage information is kept accurately, is appropriately sized, backed up, and managed throughout its useful lifecycle;
- provide and ensure everyone with a need to access information can do so;
- work with every group to determine the right metrics, measures and metadata that should be included in “official” output;
- develop reporting output expectations document (update over time as maturity grows).


### 12. Chargeback & finance integration
Chargeback andfFinance integration is about pushing spend accountability to the edges of the organization that are responsible for creating the expense.

Chargeback is the focus in this capability, but Showback is a foundational part of any FinOps practice. The difference is that Chargeback sends expenses to a product or department P&L and Showback shows the charges by product or department but keeps the expenses in a centralized budget.

A tagging and account strategy is vital as ways to identify costs. By either assigning a tag to a resource, or by designating a cost center that pays for resources in a certain account, practitioners can identify who is accountable for the expense incurred.

The goal is to make it easy for users to be accountable for their expenses. The best way to do this is by integrating with the company’s finance tools and processes.

#### As someone in a Finance/FinOps role, I will…
- understand how cloud expense is generated;
- ensure there is appropriate documentation on chargeback/showback policy and that operations are auditable according to company policies;
- help teams reconcile their portion of expense that is allocated to them each month;
- help teach teams the tagging and account policies and the importance of expense accountability.


### 13. Onboarding workloads (NEW)
This capability is about establishing a cloud front door process to onboard brownfield and greenfield applications through financial viability and technical feasibility assessment criteria.

#### As someone in a Finance/FinOps role, I will…
- assess the financial viability of the workloads moving to the cloud;
- capture and track cloud value benefits realization;


### 14. Establishing FinOps culture (NEW)
This capability is about creating a movement to establish cultures of accountability so that your organization understands the practice of cloud cost management is really about leveraging FinOps to accelerate the creation of business value.

An organization’s FinOps culture embraces a long term roadmap of transformation and continuous improvement across all the FinOps domains at each stage of FinOps maturity.

It requires finding allies in the face of opposition, winning over detractors during times of change, defining a common language, creating metrics to prove value, building enablement strategies to elevate stakeholder teams, and developing communication programs to inspire the FinOps personas in your organization.

#### As someone in a FinOps Practitioner role, I will…
- establish cloud cost management best practices;
- establish benchmarks for stakeholder teams to use;
- centralize cloud cost management in single cloud or multi-cloud environment;
- create visibility and transparency to cloud cost;
- align accountability to cloud users;
- identify unallocated spend;
- create and contribute to cloud budgets and forecasts;
- create and contribute to showback to increase financial accountability;
- advance communication and socialize FinOps throughout the organization.


### 15. FinOps & intersecting frameworks(NEW)
This capability examines the intersection between FinOps with other standards and frameworks used within your organization. 
- FinOps & ITAM/SAM
- FinOps & ITFM/TBM
- FinOps & Sustainability
- FinOps & ITSM
- FinOps & Security


### 16. Cloud policy & governance
Policy and Governance can be thought of as a set of statements of intent, with associated assurances of adherence.

A “Cloud Policy” is a clear statement of intent, describing the execution of specific cloud-related activities in accordance with a standard model designed to deliver some improvement of business value.

“Cloud Governance” is a set of processes, tooling or other guardrail solution that aims to control the activity as described by the Cloud Policy to promote the desired behaviour and outcomes.

Combining good Policy and Governance provides us with a mechanism to orchestrate and direct our Cloud FinOps Activity.

Policy and Governance is how we establish and sustain a FinOps culture. In fact, it is the way in which all culture is established and sustained. 


### 17. FinOps education & enablement
FinOps education & enablement allows all those participating in FinOps practices to increase the business value of cloud by accelerating FinOps adoption.

- Internal communications, events and learning experiences. These may focus on specific technical, financial or business topics or a combination of all three.
- Training – either function-based, technology-based, or focused on FinOps processes themselves
- Initiatives aimed at improving the business value of cloud that give participants the opportunity to use the knowledge they have gained.


### 18. Establishing a FinOps decision & accountability structure
Establishing a FinOps decision & accountability structure is about capturing an organization’s FinOps-related roles, responsibilities and activities to bridge operational cloud cost management gaps between teams.

These decision-making and accountability structures help cross-functional teams work out the processes and decision trees they’ll need to use to tackle challenges and resolve conflicts, in addition to having them be proactively available when they need to take action ahead of time.

This structure should include clear lines of authority and guidelines, enabling issues to be escalated and resolved, and giving senior decision makers the opportunity to make informed decisions quickly and in a consistent manner.


## FinOps Maturity Model
A “Crawl, Walk, Run” approach to performing FinOps enables organizations to start small, and grow in scale, scope, and complexity as business value warrants maturing a functional activity.

In assessing the state of an organization’s FinOps capability or domain, we can also use these maturity designations to identify where we are currently operating, and to identify areas we’d like to move from a Crawl to a Walk, or from a Walk to a Run maturity.

In other words, establishing a Walk stage in a particular capability is not necessarily a good or bad thing, FinOps practitioners should focus less on maturing each Capability to “Run” for everything and more on achieving the outcomes the FinOps capabilities aim to provide.

Prioritize maturing the capabilities that provide your organization the highest business value.

Every capability and functional activity can be at a different level of maturity.

### Crawl
- Very little reporting and tooling
- Measurements only provide insight into the benefits of maturing the capability
- Basic KPIs set for the measurement of success
- Basic processes and policies are defined around the capability
- Capability is understood but not followed by all the major teams within the organization
- Plans to address “low hanging fruit”

### Walk
- Capability is understood and followed within the organization
- Difficult edge cases are identified but decision to not address them is adopted
- Automation and/or processes cover most of the Capability requirements
- Most difficult edge cases (ones that threaten the financial well-being of the organization) are identified and effort to resolve has been estimated
- Medium to high goals/KPIs set on the measurement of success

### Run
- Capability is understood and followed by all teams within the organization
- Difficult edge cases are being addressed
- Very high goals/KPIs set on the measurement of success
- Automation is the preferred approach

## FinOps Personas
- FinOps Practitioner
- Executives
- Business/Product Owner
- Engineering and Operations
- Finance
- Procurement

### FinOps Practitioner
-Primary Goal: Drive best practices into the organization through education, standardization, and cultural growth and support.

### Executives

#### CEO
Primary Goal: Assurance that cloud investments are aligned with business objectives.

#### CTO / CIO
Primary Goal: Leverage technology to give the business a market and competitive advantage.

#### CFO
Primary Goal: Manage the cost of cloud utilization (among other costs across the org) and ensuring that money is wisely spent.

### Business/Product Owner

#### Product Owner
Primary Goal: Quickly bring new products and features to market with an accurate price point.

### Engineering and Operations

#### Engineering Lead
Primary Goal: Deliver faster and high quality services to the organisation, whilst maintaining business as usual.

### Finance

#### IT Finance Manager
Primary Goal: Accurately budget, forecast, and report cloud costs.

### Procurement
Primary Goal: Cloud platform relationship management.

### ITAM Leader / Practitioner (NEW)
Primary Goal: Ensure that all IT assets are being used to their fullest potential and that the organization is getting the best value for its investments in IT. Also ensuring that the organization is fully compliant with all licensing and regulatory requirements.

### Sustainability Practitioner within IT (NEW)
Primary Goal: To integrate sustainability principles into enterprise IT operations, reducing environmental impact, and driving cost and carbon efficiencies.


## FinOps Phases
The FinOps journey consists of three iterative phases — Inform, Optimize and Operate.

### Inform
- This is the first phase in the FinOps journey, empowering organizations and teams with visibility, allocation, benchmarking, budgeting, and forecasting. 
- The on-demand and elastic nature of cloud, along with customized pricing and discounts, makes it necessary for accurate and timely visibility for intelligent decisions.
- Accurate allocation of cloud spend based on tags, accounts, or business mappings enable accurate chargeback and showback.
- Business and financial stakeholders also want to ensure they are driving ROI while staying within budget and accurately forecasting spend, avoiding surprises. 
- Benchmarking as a cohort and against teams provides organizations with the necessary metrics to develop a high performing team.

### Optimize
- Once organizations and teams are empowered, they need to optimize their cloud footprint.
- Cloud providers offer multiple levers to optimize. 
    - On-demand capacity is the most expensive. 
    - To encourage advanced reservation planning and increased commitment, cloud providers offer discounts for commitments which typically involves complex calculations for making reservations - Reserved Instances (RI) / Committed Use Discounts (CUD – Google Cloud). 
- In addition, teams and organizations can optimize the environment by rightsizing and automating turning off any wasteful use of resources.

### Operate
- Organizations start to continuously evaluate business objectives and the metrics they are tracking against those objectives, and how they are trending. 
- They measure business alignment based on speed, quality, and cost. 
- Any organizational success is only possible if the organization builds a culture of FinOps which involves a Cloud Cost Center of Excellence built around business, financial, and operational stakeholders who also define the appropriate governance policies and models.


## FinOps Principles
FinOps Principles act as a north star that guide the activities of our FinOps practice. 

### Teams need to collaborate
- Finance, technology, product, and business teams work together in near real time as the cloud operates on a per-resource, per-second basis.
- Teams work together to continuously improve for efficiency and innovation.

### Decisions are driven by business value of cloud
- Unit economic and value-based metrics demonstrate business impact better than aggregate spend.
- Make conscious trade-off decisions among cost, quality, and speed.
- Think of cloud as a driver of innovation.

### Everyone takes ownership for their cloud usage
- Accountability of usage and cost is pushed to the edge, with engineers taking ownership of costs from architecture design to ongoing operations.
- Individual feature and product teams are empowered to manage their own usage of cloud against their budget.
- Decentralize the decision making around cost-effective architecture, resource usage, and optimization.
- Technical teams must begin to consider cost as a new efficiency metric from the beginning of the software development lifecycle.

### FinOps data should be accessible and timely
- Process and share cost data as soon as it becomes available.
- Real-time visibility autonomously drives better cloud utilization.
- Fast feedback loops result in more efficient behavior.
- Consistent visibility into cloud spend is provided to all levels of the organization.
- Create, monitor, and improve real-time financial forecasting and planning.
- Trending and variance analysis helps explain why costs increased.
- Internal team benchmarking drives best practices and celebrates wins.
- Industry peer-level benchmarking assesses your company’s performance.

### A centralized team drives FinOps
- The central team encourages, evangelizes, and enables best practices in a shared accountability model, much like security, which has a central team yet everyone remains responsible for their portion.
- Executive buy-in for FinOps and its practices and processes is required.
- Rate, commitment, and discount optimization are centralized to take advantage of economies of scale.
- Remove the need for engineers and operations teams to think about rate negotiations, allowing them to stay focused on usage optimization of their own environments.

### Take advantage of the variable cost model of the cloud.
- The variable cost model of the cloud should be viewed as an opportunity to deliver more value, not as a risk.
- Embrace just-in-time prediction, planning, and purchasing of capacity.
- Agile iterative planning is preferred over static long-term plans.
- Embrace proactive system design with continuous adjustments in cloud optimization over infrequent reactive cleanups.