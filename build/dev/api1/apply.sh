#!/bin/bash
kubectl delete -f build/dev/api1/api1-deployment.yaml
kubectl delete -f build/dev/api1/api1-service.yaml
kubectl delete -f build/dev/api1/api1-ingress.yaml

kubectl apply -f build/dev/api1/api1-deployment.yaml
kubectl apply -f build/dev/api1/api1-service.yaml
kubectl apply -f build/dev/api1/api1-ingress.yaml