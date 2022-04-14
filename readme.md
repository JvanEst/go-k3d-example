# K3D Go example

Example of a K3D cluster with multiple namespaces and services.

This example uses Kafka and Cassandra.

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
docker push localhost:5050/api1:latest
```

## Kafka

Apply all the files in the `k3d/kafka` directory:
```
kubectl apply --namespace=kafka -R -f k3d/kafka
kubectl wait kafka/my-cluster --for=condition=Ready --timeout=300s -n kafka
```

### Test if Kafka works

This makes a new topic called `my-topic` and will allow you to post your own message on the topic to then read.

Producer
```
kubectl run kafka-producer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-producer.sh --broker-list kafka-kafka-bootstrap.kafka.svc.cluster.local:9092 --topic my-topic
```

Consumer
```
kubectl run kafka-consumer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server kafka-kafka-bootstrap.kafka.svc.cluster.local:9092 --topic my-topic --from-beginning
```

## Cassandra

Apply all the files in the `k3d/cassandra` directory:
```
kubectl apply --namespace=cassandra -R -f k3d/cassandra
kubectl rollout status --watch --timeout=600s statefulset/cassandra -n cassandra
```

### Connect to cassandra to test
```
kubectl run cassandra-cql -ti --image=cassandra:latest --rm=true --restart=Never -- cqlsh -u cassandra -p <password> cassandra.cassandra.svc.cluster.local
```

## Removing Kafka and/or Cassandra
```
kubectl delete --namespace=cassandra -R -f k3d/cassandra

kubectl delete --namespace=kafka -R -f k3d/kafka
```