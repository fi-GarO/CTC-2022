
# Kubernetes 
# Commands:
## Install 
sudo snap install kubectl --classic

## Get node
kubectl get node

response:
NAME                   STATUS   ROLES                  AGE     VERSION
lima-rancher-desktop   Ready    control-plane,master   3h24m   v1.23.6+k3s1

## apply configuration
kubectl apply -f ./k8s.yml

response: 
statefulset.apps/etcd created
service/etcd created
deployment.apps/ctcgrpc created
service/ctcgrpc created

## Get pod
kubectl get pod

NAME                       READY   STATUS    RESTARTS        AGE
etcd-0                     1/1     Running   0               2m45s
ctcgrpc-6575bfcbf5-59g2l   1/1     Running   3 (2m12s ago)   2m44s
ctcgrpc-6575bfcbf5-bnmc7   1/1     Running   3 (2m11s ago)   2m45s

### Assignment

Deploy GRPC server and Etcd from assignment 5 into Kubernetes

1. Deploy etcd as stateful set with `local-path` PVC
2. Deploy GRPC server as Deployment
3. Connect to GRPC server using GRPC client using Rancher port forwarding or NodePort
