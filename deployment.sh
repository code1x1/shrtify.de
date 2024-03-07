#!/usr/bin/env zsh

docker build -t localhost:5000/frontend-shrtify-de -f Dockerfile.frontend .
docker build -t localhost:5000/backend-shrtify-de -f Dockerfile.backend .

kubectl port-forward --namespace kube-system kube-registry-v0-pltfq 5000:5000 &

sleep 2

docker push localhost:5000/backend-shrtify-de
docker push localhost:5000/frontend-shrtify-de
