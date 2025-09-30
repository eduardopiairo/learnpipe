# Platform Engineering Certified Practitioner


## Module 1: Intro to Platform Engineering

### What is platform engineering?
Platform engineering is the discipline of designing and building toolchains and workflows that enable self-service capabilities for software engineering organizations in the cloud-native era. 

Platform engineers provide an integrated product most often referred to as an “Internal Developer Platform” covering the operational necessities of the entire lifecycle of an application.


### What is an Internal Developer Platform?

An Internal Developer Platform (IDP) is built by a platform team to build golden paths and enable developer self-service. 
An IDP consists of many different techs and tools, glued together in a way that lowers cognitive load on developers without abstracting away context and underlying technologies.

Following best practices, platform teams treat their platform as a product and build it based on user research, maintain and continuously improve it.

### Why platform engineering has emerged, and its impact on the industry

- Complexity
- Significant cognitive load on developers
- Need of self-serve their infrastructure needs through "golden paths"

A big number of organizations will adopt platform engineering initiatives to improve developer experience, accelerate time to market, enhance standardization and security, and ultimately run their engineering organizations more efficiently.

### Essential skills to learn to become a Platform Engineer
- Define and build golden paths and abstraction layers
- Knowledge of both backend and frontend designs for platform concepts
- Be proficient in automation and standardization
- Adopt a product mindset, where developers are viewed as customers 
- Understand stakeholder management and identify and addess pain points in developer and operations workflows.
- Pragmatic apporach to tooling


### Homework 

Platform Engineering Maturity Model
- Investment (How are staff and funds allocated to platform capabilities?) - Voluntary or temporary
- Adoption (Why and how do users discover and use internal platforms and platformcapabilities?) - Erratic
- Interfaces (How do users interact with and consume platform capabilities?) - Custom processes
- Operations (How are platforms and their capabilities planned, prioritized, developed and maintained?) - By request
- Measurment (What is the process for gathering and incorporating feedback and learning?) - Ad hoc / nonexistent


## Module 2: How to build an Internal Developer Platform

### Platform as a Product 

"A digital platform is a foundation of self-service APIs, tools, services, knowledge and support which are arranged as a compellling internal product".

Key pillars:
- Ownsership: the platform is a product owned by the platform team that follows the product life cycle;
- New perspective: developers are threaded as internal customers;
- Prduct management basics: user research, user centric design, product roadmap, user feedback, product marketing, etc. 

### Platform as a Product: 10 best practices
1. Do not reinvent the wheel
2. Apply well established architecture patterns
3. Everything as code
4. Build golden paths over cages
5. Be Pareto efficient (80-20) 
6. Leave pltaform interface choice to the developer
7. Security by design
8. Measure from the beginning
9. Gain stakeholder buy-in
10. Think about adoption from day 1


### Platform architecture choices

Backend design
- Pipeline-based
- Graph- based (often leveraging a Platform Orchestrator) 

- Architecture reference
- 3 Tier platform architecture
    - Application Choreography "Code, ship, run"
    - Platform Orchestration "Design enable, optimize"
    - Infrastrcture Orchestration "Plan, build, maintain"

### Front-end and back-end designs

The __front-end__ of an IDP serves as the primary interface through which developers interact with the underlying platform capabilities.

Key apects:
- Ergonomics and design
- Presentation
- Handling multiple aspects
- User Experience (UX) and User Interface (UI) standards
- Variety of interfaces

It's crucial to understand that the front-end (like a portal) is not the platform itself but merely an interface to the underlying infrastructure and automation. 

The __back-end__ of an IDP is where the core logic for standardisation, security enforcement, automation, and the overall orchestration of platform capabilities resides. 

Key aspects:
- Automation and standardization
    - The back-end is primarily concerned with automating tasks and enforcing organisational standards across infrastructure and application lifecycles. 
- Lifecycle management
    - Back-end abstractions define how resources and applications are created, updated, and eventually decommissioned.
- Handling one aspect of many things
    - The back-end often handles a single aspect (e.g., deployment, security policy enforcement) across a large number of applications or infrastructure components.
- Enforcement by design
    - A well-designed back-end enforces best practices, compliance, security, and governance at a fundamental level, often making them automatic and less reliant on manual checks.
- Orchestration
    - The back-end often includes an orchestration layer that manages the complex workflows and dependencies involved in provisioning and managing resources. 


### Platform tooling landscape (The 5 planes)

#### 1 - Developer Control Plane

This is the primary interface through which both application developers and platform engineers interact with the platform's capabilities.

- IDE
- Developer Portal
- Version Control
    - Application Source Code
    - Platform Source Code

#### 2 - Integration & Delivery Plane

This is where your CI/CD pipelines and other automation logic reside. It acts as the back-end of your platform, containing the logic that can be orchestrated to deliver platform capabilities. This includes tools for building and storing images, deployment tools, and potentially a Platform Orchestrator.

- CI Pipeline
- Image Registry
- Platform Orchestrator
- CD Pipeline
- Infrastructure Control Plane

#### 3 - Resource Plane

This represents the underlying infrastructure – your physical infrastructure or cloud resources.

- Compute
- Data
- Networking
- Services

#### 4 - Monitoring & Logging Plane

This layer address observability aspects.

- Observability
- Analytics

#### 5 - Security Plane

This layer adress security aspects.

- Secrets Management
- Security


### Where do we start?

Start with the __backend__. Here's why:

- Focus on core logic
    - The back-end is where the core logic for automation, standardization, and policy enforcement resides. Establishing this solid foundation first is crucial.
- Avoid shoehorning logic
    - Starting with the front-end (a "portal-first" approach) can lead to prematurely trying to fit complex business logic into the presentation layer, which violates the single responsibility principle and can hinder scalability.
- Enable future flexibility
    - A well-defined back-end allows for more flexibility in the future when you decide to build various front-end interfaces (UI portals, CLIs, APIs).


Graph-based backends -> Platform Orchestrators:
- Developer
    - Writes a workload specification defining the workload and the the resources it needs
- Platform Engineer
    - Configures how resources are created and where the workloads are deployed

Orchestration allows to build Golden Paths (Standardization and Automation)


Examples: 
- CNOE Setup
- Pocket IDP

### Homework
Create your own platform reference architecture.


## Module 3: Platform tooling 101

> If you follow Golden Paths everybodt wins

### What tools platform teams are using

#### Developer Control Plane
This layer focuses on providing developers with self-service capabilities and a streamlined, user-friendly experience.

- Version Control
- Developer Portals
- Infrastructure as Code (IaC)

#### Integration & Delivery Plane
This plane encompasses the tools and processes involved in building, testing, and deploying software efficiently and reliably.

- CI/CD
- Container Registries
- Platform Orchestrators

#### Resource Plane
The Resource Plane represents the underlying infrastructure that supports the applications and services. This can encompass cloud resources, on-premises data centres, or a hybrid environment.

- Infrasctructure provisioning 
- Container Orchestration
- Resource Management

#### Monitoring & Logging Plane
The Monitoring & Logging Plane is crucial for gaining insights into the performance, health, and behaviour of applications and the underlying infrastructure.

- Metrics collection
- Data visualization
- Logs aggregation
- Alerting

#### Security Plane
The Security Plane focuses on ensuring the security and compliance of the platform and the applications running on it.

### Hands on lab

https://github.com/InternalDeveloperPlatform/pe-course-labs

#### CNOE
CNOE is presented as a framework designed for scaffolding an Internal Developer Platform (IDP), employing a pipeline-based backend. This implementation primarily utilises cloud-native open-source tools such as Argo Workflows and ArgoCD for Continuous Integration and Continuous Delivery, Git for version control and as a container registry, and optionally integrates with Backstage as a developer portal.

- CNOE is a framework to builf cloud-native INternal Developer Platform mainly based on CNCF open-source tools.
- CNOE is a platform builder, not your platform. 
- It's a tool for scaffolfing an IDP using open-source tools. 
- Pipeline based backend

https://github.com/InternalDeveloperPlatform/pe-course-labs/blob/main/cnoe.md

#### PocketIDP
Conversely, the PocketIDP demonstrates an IDP architecture centred around a Platform Orchestrator (Humanitec), representing a vendor and graph-based backend. This approach focuses on creating a unified workflow that connects developer application code with resource definitions. The PocketIDP implementation incorporates Git for version control, Backstage as a developer interface, Humanitec Pipelines for Continuous Delivery, and Score for defining application workload specifications.

- IDP with a Platform Orchestrator.
- Graph-based backend 
- The purpose of a Platform Orchestrator is to connect the developer's application source code to the platform source code laid out by platform/infra teams.

https://github.com/InternalDeveloperPlatform/pe-course-labs/blob/main/pocketidp.md


## Module 4: The art of building golden paths

### What's a Golden Path?

Also known as Happy Path!

- Platform user performs a task
- Minimal cognitive load
- Nothing can go wrong
- Standardize and automated
- Diff to "golden cage" (you can also go off path)

A golden path is a preconfigured, paved road that offers developers an end-to-end workflow for common tasks. A true golden path allows for flexibility; developers should be able to go off-path when needed.

In contrast, a golden cage arises when developers are forced to follow potentially non-ideal patterns with little to no flexibility

### How to identify Golden Paths

#### Value Stream Mapping
- Visualization of steps of the value stream
- Lean methodology to identify waste
- "Waste" in our context:
    - Waiting times
    - Manual touchpoints/handoffs
    - Context switching 
    - Operational steps without adding value

Key metrics:
- Lead time
- Process time (work time)
- % of CA (complete and accurate)

IDP can also act as a discovery mechanism, allowing developers to request new functionalities, which can then inform the creation of new golden paths over time.

### Common Golden Paths
- Scaffolfing, create a new service

### Pain points in day 2 operations
- Operations that developers do every day

Add a postgress to an existing workload and deploy all the way to production
- Developer
- Ops/infra
- Security

Frequency makes the diffference:
- Scaffolding: 3x per year
- Deployment. 3 x per day

### Golden Paths - Sorted by Impact
- Daily deployments including changes in configuration (happening every day)
- Spinning up ephemeral environments (needed with every review)
- Adding resources (happening sometimes in the project lifetime)
- Scaffolding new workloads and create environments (happening once per project / product)

## Homework
- Mapp out at least 2 end-to-end developer workflows
    - How are they corrently performed?
    - Whtat's the golden path version?


## Module 5: Find the right abstractions

### Why do we need abdstractions?

"Let's introduce an abstraction layer to shield away complexity..."

### Frontend vs Backend

#### Frontend
- Ergonomics & design
- How to use someting
- How to display something
- How to handle a lot of aspects of one thing at the same time

We use abstractions as a term in general to describe details being hidden behind a simpler interface
- Visual
- Usage
- Collaborative

##### Portals
- Backstage
- Cortex
- Port

Pros
- Visualization
- Overview in terms of onwership
- Managment loves them

Cons:
- Risk of low developer adoption
- Context switching
- ClickOps
- Risk incomplete visualization, lack of integrations

Portals can function as an abstraction layer to visualize existing information (meta dat) but they do not solve underlying problems. 

##### Declarative configurations (workload spec)
- Score
    - developer-centric
    - platform-agnostic
    - declarative (workload specification)
- Radius
- KubeVela


#### Backend
- Automation
- How to use something
- How to lifecycle something
- How to handle one aspect of a lot of things at once at the same time


## Module 6: Infrastructure Platform Engineering (IPE)

### Infrastructure PLatform Engineering (IPE) definition
IPE is the discipline of building internal infrastructure products that present infrastructure to users or other parts of the platform like portals in an easely consumable was as a service.

Infrastructure platforms are self-service tools that allow non-expert users to deploy and manage infrastructure themselves while I&O retains governance, security and compliance. 

IPE is about service delivery and can be defined by the following core principles:
- Productizing infrastructure
- Standardization
- Automation
- Self-service


### Rules for designing APIs for platforms
- Ubiquitious
- Interface agnostic
- Be prepared for complex backend logic

### What makes good infrastructure "products"?
- Interfaces
- Better Processs
- Composability

### A way to enhance composability: Implement cross-cutting-concerns
- Cross-cuting
- Usability
- Convevience in Commodity

Examples: observability, cost, security

### Everything as Code

- Repos: the usual status quo
- Scaffolding: first step forward

Do it right: clear separation of concerns
- App source code (app devs)
    - workload sources code
    - workload spec (score)
    - Docker file
    Pipeline yaml
- Platform source code (ops/infra)
    - Resource definitions
    - Infrastructure as code
    - Automation/compliance

Why?
- 95% less config files to main (no longer 1-1 between app and resources config)
- high degrees of standardization
- security baked in
- allows for specialization
- allows for self-service
- easy to update


## Module 7: How to build your Minimum Viable Platform (MVP)

### MVP Why, What, How?

- Do not get beyond planning phase
- Try to build something really large that would take them 2-5 years
- Fail to prove value and get further sponsorship
- Fail to gain stakeholder buy-in
- Fail to generate adoption for the platform

Risk of failure: low value, low adoption, no funding

A MVP (Minimum Viable Platform) is foundational version of a platform designed to quickly demonstrate value and support scalability without requiring extensive initial development.  

### Make Platform Engineering Predictable

MVP Program
- Duration: 8 weeks
- Outcome: Demoable first platform running
Production Readiness Program (PRP)
- Duration: 8 weeks
- Outcome: First teams using platform daily
Adoption Program
- Outcome: Large scale adoption

### Desired outcomes by persona
- Infra & Ops teams: develop technical conviction that platform is vending machine layer for their infrastructure
- App devs: convinced that platform will enable developer selservice at low cognitive load
- Security teams: convinced that the design i secure and future-proof
- Executives: belive that ROI justifies investment in production-grade platform  

### Overall goals of an MVP
- Show value within weeks, not years
- Generate stakeholder buy in ("what's in for me?")
- Ensure long term funding
- Set the stage for adoption

### From design perspective, MVPs should be
1. Representative 
2. Repeatable
3. Iterative
4. Innovative

### MVP Framework
Phase 1: Discovery
- Track 1: Tech - design reference architecture and define scope
    - MVP technical Discovery Workshop
        - Which tools are already in use?
        - What is missing?
        - Define tartget reference architecture
        - Define golden paths you want to start with
- Track 2: Business - measure status quo
    - MVP Objetive + Goals Planning Workshop
        - Business objectives
        - Technical objectives
        - KPIs and scucess criteria
- Track 3: Security - understand security posture
    - Security workshop
        - Document security requirements
Phase 2: Integration
- Track 1: Tech - integrate platform tooling
    - Admin/RBAC
    - Cloud provider integration
    - Cluster connection
    - Deploy sample application
    - Envrionment creation for non-prod envs
    - resource creation/definitions
- Track 2: Business - develop business case
- Track 3: Security - gather requirements
Phase 3: Deployment
- Track 1: Tech - Onboard 1st app
    - Application integration
    - Resources matching
    - Values + Secrets
    - CI/CD pipeline setup
    - Portal integration, if required
    - Workload spec setup (Score)
    - E2E deployment test 
- Track 2: Business - involve finance team
- Track 3: Security - engage with security team
Phase 4: Adoption
- Track 1: Tech - onboard 1st team
    - Success testing
    - First dev team onboarding
    - Demo prep + training
    - Platform roadmap planning
    - Additional team demos + onboarding planning
- Track 2: Business - measure MVP, build ROI 
- Track 3: Security - green light to production
    - Full-scale security requirements confirmation

#### MVP Success Metrics
- Complexity index
- Onboarding time
- Service create time
- Survey.... (development satisfaction score)


## Module 8: How to sell your MVP to key stakeholders

### Who you are selling to
- App developers
- Infra teams
- Executives
- Security teams
- Architects and CCOE

### Understand your customers and stakeholders's problems first
- Which personas do you sell to?
- What are their pain points/problems?
- What really matters to them?
- What are the symptoms of their problems?
- What's the impact on the persona?
- What's the impact on the business?
- How your platform solves their pain points?
- What are proof points/metrics?

### App Developers
Main pain points:
- Waiting times
- Cognitive load (complexity)
- Fear of screwing up
- Shadow ops

#### How to sell to application developers?
- Build tight customer feedback loops
- Provide internal documentation
- Brand your platform, give it a name
- Onboarding and training workshops
- Identify the right first team
- Convert the first into advocates

### Infra teams
Main pain points:
- Lack of standardization
- Ticket Ops
- Need to compete against third party like cloud provider interfaces
- Under constant pressure, treated as a cost center 

#### How to sell to Infra teams?
- A commmon layer that equilizes consumption experience of infra and cloud, boosts efficiency, improves security, and strengthens compiance across all providers 
- Be able to reduce costs that are ever-growing
- Gain control and overshight

### Security teams
Main pain points:
- Developers do not care about security
- Time consuming manual security checks
- Fear of overseeing security violations

#### How to sell to Security teams?
- Security by design by implementing security best practices into developer's golden paths
- Enforce gardrails and policies by default
- Automated security checks

### Executives
Main pain points:
- Costs
- Long time to market, slow innovation cycles
- Security, compliance
- Talent retention

#### How to sell to Executives?
- Create business case
- Make a ROI calculation the core of it
