#!/bin/bash
kubectl delete -n redis -f ./redis-configmap.yaml
kubectl delete -n redis -f ./redis-statefulset.yaml
kubectl create -n redis -f ./redis-configmap.yaml
kubectl create -n redis -f ./redis-statefulset.yaml