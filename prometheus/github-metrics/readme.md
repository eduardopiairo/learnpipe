# GitHub Metrics

## Get Helm Repository Info (kube-prometheus-stack)

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

```
helm repo update
```

## Create namespace

kubectl create namespace monitoring

kubectl create namespace demo


## Install stack

helm upgrade --install prom prometheus-community/kube-prometheus-stack -n monitoring --values values.yaml


kubectl port-forward service/prometheus-operated -n monitoring 9090:9090

kubectl port-forward service/prom-grafana -n monitoring 3000:80


kubectl apply -f deployment.yaml -n demo

kubectl apply -f service.yaml -n demo
kubectl apply -f servicemonitor.yaml -n demo
