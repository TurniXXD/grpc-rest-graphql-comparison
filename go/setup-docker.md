# Setup docker

## Dev

- Build `docker build --target dev . -t go`
- Run and execute terminal `docker run -it -p 80:80 -v ${PWD}:/app go sh`

## Build

- Build `docker build --target build . -t server`
- Run docker container
```bash
docker run -it -p 80:80 --net redis -e REDIS_SENTINELS="sentinel-0:5000,sentinel-1:5000,sentinel-2:5000" -e REDIS_MASTER_NAME="mymaster" -e REDIS_PASSWORD="1234" server
```
