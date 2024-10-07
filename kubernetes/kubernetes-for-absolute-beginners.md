# Kubernetes For Absolute Beginners

## K8s Architecture

Node: is a phisical or a virtual machine. Master and Worker nodes

Cluster: a set of nodes

Components:
- API Server (acts as the forntend)
- etcd (key-value store)
- kubelet (agent that runs in each node)
- Container runtime (software to run containers)
- Controller (brain behind orchestration)
- Scheduler (distribute work or containers)


### Master vs Worker Nodes

Worker node - were the containers run
    - kubelet
Master node - were the API lives
    - etcd
    - controller
    - scheduler

### Kubectl

Is the K8s CLI.