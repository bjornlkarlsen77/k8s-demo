# k8s-demo

This repository is used for testing and demonstrating kubernetes features

## Deploy
Need to configure ghcr-pull-secret to grant access to pull images from ghcr.
Create a personal access token in [github](https://github.com/settings/tokens)

See https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/

```bash
helm install k8s-demo 
```

## Minikube
Enable ingress to be able to access public services 
Enable metrics-server to be able use [hpa](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)

```bash
minikube addons enable metrics-server
minikube addons enable ingress
```

### Configure /etc/hosts
Get the ip for minikube and add a mapping in /etc/host to map the domain "minikube" to its ip
```bash
 minikube ip
 ```


## Load testing
Run 2000 request in 5 threads
```bash
seq 1 2000 | xargs -n1 -P5 -I {} curl --connect-timeout 3 "http://minikube/go-hello/slow"
```

Spin up a pod that continously calls an api
```bash
kubectl run -i --tty load-generator --rm --image=busybox:1.28 --restart=Never -- /bin/sh -c "while sleep 0.01; do wget -q -O- http://minikube/go-hello/slow; done"
```

## Watch auto scaling
```bash
 k get hpa -w
 ```