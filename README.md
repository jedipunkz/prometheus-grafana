# Grafana & Prometheus docker-compose env

## Pre-requirement

- docker
- docker-compose

## Boot up docker-compose

```shell
$ docker-compose build
$ docker-compose up -d
```

## Access Grafana via Browser

```
http://localhost:3000
```

## Access Prometheus via Browser

```
http://localhost:9090
```

## Clean Up

```shell
$ docker-compose stop
$ docker-compose rm
```

## Migration Data

old host

```shell
sudo docker run --rm -v prometheus_data:/data -v $(pwd):/backup busybox tar czvf /backup/prometheus_data.tar.gz -C /data .
scp prometheus_data.tar.gz <new_host>:<path>/
```

new host

```shell
sudo docker volume create prometheus-grafana_prometheus_data
sudo docker run --rm -v prometheus-grafana_prometheus_data:/data -v /tmp:/backup busybox tar xzvf /backup/prometheus_data.tar.gz -C /data
```
