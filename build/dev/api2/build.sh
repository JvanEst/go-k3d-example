#!/bin/bash
docker build -f ./build/dev/api2/Dockerfile -t localhost:5050/api2 .
docker push localhost:5050/api2:latest