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


### docker un link containers

--link connects two containers

```
docker run -d --name=vote -p 5000:80 --link redis:redis voting-app 
```

connects the voting app and the redis ccontainers.
