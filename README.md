# devops-with-kubernetes

https://jakousa.github.io

# Notes

## Setup local docker registry

https://k3d.io/usage/guides/registries/#using-a-local-registry

```
docker volume create local_registry
docker container run -d --name registry.localhost -v local_registry:/var/lib/registry --restart always -p 5000:5000 registry:2
```

Update `/etc/hosts`, add

```
127.0.0.1 registry.localhost
```

## Create app

```
npm init
Write some code and a Dockerfile
(...)
```

## Build app

Build and push dockerfile to local registry
```
docker build -t registry.localhost:5000/docker-with-kubernetes .
docker push registry.localhost:5000/docker-with-kubernetes
```

Verify app locally
```
$ docker run --rm registry.localhost:5000/docker-with-kubernetes

> devops-with-kubernetes@1.0.0 start
> node src/app.js

2020-10-26T23:03:24.361Z: e0bbade2-9bd5-4f30-b38b-6045099e4c16
2020-10-26T23:03:29.369Z: e0bbade2-9bd5-4f30-b38b-6045099e4c16
(...)
```
`docker stop` to kill container

## Deploy to k3d cluster

Create k3d cluster with registry
```
k3d cluster create mycluster -a 2 --volume $(pwd)/registries.yaml:/etc/rancher/k3s/registries.yaml
docker network connect k3d-mycluster registry.localhost
```

Deploy
```
$ kubectl create deployment hashgenerator-dep --image=registry.localhost:5000/docker-with-kubernetes
deployment.apps/hashgenerator-dep created

$ kubectl get pods
NAME                                 READY   STATUS        RESTARTS   AGE
hashgenerator-dep-7948967794-gxtml   1/1     Running       0          17s

$ kubectl logs hashgenerator-dep-7948967794-gxtml

> devops-with-kubernetes@1.0.0 start
> node src/app.js

2020-10-26T23:04:37.628Z: 7a1c5190-b4f8-40da-8c33-e2ea3d8b8bcc
2020-10-26T23:04:42.635Z: 7a1c5190-b4f8-40da-8c33-e2ea3d8b8bcc
2020-10-26T23:04:47.635Z: 7a1c5190-b4f8-40da-8c33-e2ea3d8b8bcc
(...)
```

Remove deployment:
```
$ kubectl delete deployments.apps hashgenerator-dep
deployment.apps "hashgenerator-dep" deleted
```
