#!/bin/bash
kubectl delete -f build/dev/api2/api2-deployment.yaml
kubectl delete -f build/dev/api2/api2-service.yaml
kubectl delete -f build/dev/api2/api2-ingress.yaml

kubectl apply -f build/dev/api2/api2-deployment.yaml
kubectl apply -f build/dev/api2/api2-service.yaml
kubectl apply -f build/dev/api2/api2-ingress.yaml