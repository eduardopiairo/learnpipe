# Docker

## Image vs container

- Image: package, template. It's used to create containers.
- Container: instances of images that are isolated


## Commands

- run containers
    - docker run image-name
    - docker run -d image-name
- list containers 
    - docker ps
    - docker ps -a
- stop containers
    - docker stop container-name
- remove container
    - docker rm container-name
- list images
    - docker images
- remove images
    - docker rmi image-name
- download image
     - docker pull image-name
- append comand
    - docker run container-name sleep 5
- execute a command
    - docker exec container-name cat ext/hosts
- attach
    - docker attach container-id

## docker run

- tag
    - docker run redis:4.0
- STDIN
    - docker run -it container-name
- port mapping
    - docker run -p 80:8000 container-name
- valume mapping
    - docker run -v /opt/datadir:/var/lib/mysql container-name
- inspect container
    - docker inspect container-name
- logs
    - docker logs container-name

## docker build

docker build . -f Dockerfile -t myflask


## docker environment variable

docker run -e APP_COLOR=blue image-name


## Commands vs entrpoint

### CMD

CMD command param1

CMD ["command", "param1"]

Command will replace the command line arguments.

FROM Ubuntu
CMD sleep 5

docker run ubuntu-sleeper sleep 10


### ENTRYPOINT

Entrypoint will append the command line arguments.

FROM Ubuntu
ENTRYPOINT ["sleep"]

docker run ubuntu-sleeper 10


## Docker Compose

Run docker containers in a single docker host.

docker run -d --name=redis redis


### docker run link containers

--link connects two containers

```
docker run -d --name=vote -p 5000:80 --link redis:redis voting-app 
```

connects the voting app and the redis ccontainers.

```
docker-compose build
```

```
docker-compose up
```

## Docker Registry

image:  docker.io/nginx/nginx
        
        docker.io   -->     registry    (gcr.io)
        nginx       -->     user or account
        nginx       -->     image or repository

docker login private-registry.io


## Docker Storage

File Storage --> /var/lib/docker
                    - aufs
                    - containers
                    - image
                    - volumes


Layered architecture: each instruction is a layer.
    - docker is able to reuse layers since they are in cache

Layers:
    - Image layer: Read Only
    - Container layer: Read Write

```
docker volume create data_volume
```

/var/lib/docker
    - volumes
        - data_volume

```
docker run -v data_volume:/var/lib/mysql mysql
```

```
docker run -mount type=bind,source=/data/mysql,target=/var/lib/mysql mysql
```

Storage drivers (based on OS):
- AUFS
- ZFS
- BTRFS
- Device Mapper
- Overlay
- Overlay2


```
docker run -v /opt/data:/var/lib/mysql -d --name mysql-db -e MYSQL_ROOT_PASSWORD=db_pass123 mysql
```

## Docker Networking

Default networks:
- Bridge (docker run ubuntu)
- none (docker run ubuntu  --network=none)
- host (docker run ubuntu  --network=host)

User defined networks

```
docker network create --driver bridge --subnet 182.18.0.0/16 custom-isolated-network
```

Embedded network: use container name. 


## Container Orchestration

- docker swarm
- kubernetes
- mesos