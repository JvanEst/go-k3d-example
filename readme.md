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

------------------------

# Kafka

kubectl apply --namespace=kafka -R -f k3d/kafka
kubectl wait kafka/my-cluster --for=condition=Ready --timeout=300s -n kafka

Test kafka:

Produce
```
$ kubectl -n kafka run kafka-producer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-producer.sh --broker-list my-cluster-kafka-bootstrap:9092 --topic my-topic
```

Consume
```
kubectl -n kafka run kafka-consumer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server my-cluster-kafka-bootstrap:9092 --topic my-topic --from-beginning
```

# Cassandra

kubectl apply --namespace=cassandra -R -f k3d/cassandra
kubectl rollout status --watch --timeout=600s statefulset/cassandra -n cassandra


$ kubectl run kafka-producer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-producer.sh --broker-list my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9092 --topic my-topic

kubectl run kafka-consumer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9092 --topic my-topic --from-beginning

kubectl run cassandra-cql -ti --image=cassandra:latest --rm=true --restart=Never -- cqlsh cassandra.cassandra.svc.cluster.local

TODO: New cassandra version!