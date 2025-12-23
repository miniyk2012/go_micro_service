## 准备工作:

1. 创建本地的k8s: kind create cluster --config=kind/cluster.yaml
2. 部署consul: consul-k8s consul-k8s install -config-file=helm/values-v1.yaml
3. 宿主机8500直接转到service的80端口: kubectl port-forward svc/consul-ui --namespace consul 8500:80