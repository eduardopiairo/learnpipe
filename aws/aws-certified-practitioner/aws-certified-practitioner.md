# AWS Certified Practitioner

## What's Cloud Computing

- On-demoand delivery of compute power, data storage, applications, and other IT resources.
- Through a cloud services platform with pay-as-you-go pricing.

### Deploy Models

- Private Cloud
- Pulic Cloud
- Hybrid Cloud

### Five charecteristics of Cloud Computing

- On-demand self service;
- Broad network acess;
- Multi-tenancy and resource pooling;
- Rapid elasticity and scalability;
- Measured service

### Six advantages of Cloud Computing

- Trade capital expense (CAPEX) for operational expense (OPEX);
- Benefit of massive economies of scale;
- Stop guessing capacity;
- Increase speed and agility;
- Stop spending money running and maintaining data centers;
- Go global in minutes.

### Problems solved by the Cloud

- Flexibility
- Cost-effectiveness
- Scalability
- Elasticity
- High-availability
- Agility

### Types of Cloud Computing

- IaaS: Infrastructure as a Service
- PaaS: Platform as a Service
- SaaS: Software as a Service


### Picing of the Cloud

- Compute
    - Pay for compute time
- Storage
    - Pay for data stored in cloud
- Network (data transfered OUT of the cloud)
    - Data transfer IN is free

### AWS Global Infrastructure

- AWS Regions
    - Specific geografic location
    - A cluster of data centers
- AWS Availability Zones
    - Each region has mutiple availability zones (at least 3, max is 6)
- AWS Data Centers
- AWS Edge Locations / Points of Presence


#### How to choose a AWS region?

- Compliance;
- Proximity;
- Available services;
- Pricing;

## IAM: Users and Groups

- IAM: Identity and Access Management (Global Service)
- Users: are people within your organization, and can be grouped. (don't have to belong a group, or can belong to multiple groups)
- Groups: only contain user
- Permissions: users or groups can be assigned policies (json documents)
    - Least privilege principles

### Policies
- Version: policy language version
- Id: identified of the policy (optional)
- Statement: one or more statments (required)
    - Sid: identifier of the statement (optional)
    - Effect: allow or deny
    - Principal: account / user /role
    - Action: list of actions this policy allows or denies
    - Resource: list of resources to which the acctions applied to
    - Condition: conditions for when this policy is in effect (optional)



