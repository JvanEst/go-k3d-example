# Go K3D example

Download K3D:
https://k3d.io/v5.4.1/

## Cluster:

Create a registry for the local docker images:
```
k3d registry create rtk-dev-registry --port 5050
```

Create the cluster:
```
k3d cluster create rtk0001-dev --registry-use rtk-dev-registry:5050 --registry-config k3d/registries.yaml  -p "8081:80@loadbalancer"
```

Tag and push the docker images:
```
docker tag api1:latest localhost:5050/api1:latest
docker push localhost:5050/api1:latest
```

Or build and tag them correctly directly:
```
docker build -f build/dev/api1/Dockerfile -t localhost:5050/api1 .
```

## Applying config

```
kubectl apply -f <file>
```