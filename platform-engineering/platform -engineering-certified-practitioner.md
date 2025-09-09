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
- Imgae Registry
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