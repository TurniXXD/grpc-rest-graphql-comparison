#!/bin/bash
docker pull kindest/node:v1.24.2
kind create cluster --name redis --image kindest/node:v1.24.2
kubectl create ns redis
kubectl apply -n redis -f ./redis-configmap.yaml
kubectl apply -n redis -f ./redis-statefulset.yaml