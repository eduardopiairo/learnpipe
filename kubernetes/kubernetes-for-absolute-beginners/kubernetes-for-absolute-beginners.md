# Kubernetes For Absolute Beginners

## K8s Overview

### K8s Architecture

Node: is a phisical or a virtual machine. Master and Worker nodes

Cluster: a set of nodes

Components:
- API Server (acts as the forntend)
- etcd (key-value store)
- kubelet (agent that runs in each node)
- Container runtime (software to run containers)
- Controller (brain behind orchestration)
- Scheduler (distribute work or containers)


#### Master vs Worker Nodes

Worker node - were the containers run
    - kubelet
Master node - were the API lives
    - etcd
    - controller
    - scheduler

#### Kubectl

Is the K8s CLI.


### Docker vs ContainerD

CRI - Container Runtime Interface
Allow any vendor to work as container runtime for Lubernetes as long as they adhere to OCI (Open Container Initiative). 

OCI - Open Container Initiative
- imagespec
- runtimespec

#### nerdctl
nerdctl - is a Docker-like CLI for caontainerD

docker --> nerdctl

#### crictl

crictl - provide a CLI for CRI compatible container runtimes. Only used for debuging purposes. 


## K8s Concepts

### Pods
Kubernetes does not deploy containers directly to nodes. Containers are deployed in pods.

Pod - single instance of an application. Pod and application should a relation of 1-1.

