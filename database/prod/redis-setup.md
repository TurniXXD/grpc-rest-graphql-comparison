# Redis setup

- `yay -S kind`
- `docker pull kindest/node:v1.24.2`
- `kind create cluster --name redis --image kindest/node:v1.24.2`
- `kubectl create ns redis`
- basic configmap for redis pods `kubectl apply -n redis -f ./redis-configmap.yaml`
- set stable settings persistent on reboot `kubectl apply -n redis -f ./redis-statefulset.yaml`
- to update `.yaml` config files run `./update-redis-config.sh`
- execute redis terminal `kubectl -n redis exec -it redis-0 sh`
  - run `redis-cli`
  - authenticate with `auth 1234`
  - show replication info `info replication`
