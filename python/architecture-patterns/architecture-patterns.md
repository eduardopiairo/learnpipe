# Architecture Patterns With Python

## Encapsulation and Abstractions

The term *encapsulation* covers two closely related ideas: simplifying behavior and hiding data.

## Layering

- Presentation layer
- Business logic
- Database layer

## Dependency Inversion Principle

Is the D in SOLID.

1. High-level modules should not depend on low-level modules. Both should depend on abstractions.
2. Abstractions should not depend on details. Instead, details should depend on abstractions. 

High-level modules
- Are the functions, classes, and packages that deal with our **real-word concepts**.
- Should be ease to change in response to business needs.

Low-level modules
- Are the code that your organization doesn't care about.

Abstractions: are simplified interfaces that encapsulate behavior. 


## Domain Modeling

