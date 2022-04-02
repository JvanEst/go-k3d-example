#!/bin/bash
docker build -f build/dev/api1/Dockerfile -t localhost:5050/api1 .
docker push localhost:5050/api1:latest