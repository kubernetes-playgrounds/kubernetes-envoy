#!/bin/bash
for F in ./envoy/*.json; do
    BASE=`basename ${F}`
    echo $BASE
    kubectl create configmap --dry-run "${BASE%.*}" --from-file=envoy.json=${F} -o yaml |
        kubectl apply --namespace=default --context=minikube -f -
done