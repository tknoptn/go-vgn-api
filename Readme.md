 **Biodata user Rest API in Go using Mux, Postgres Docker and Docker Compose**

 **importing** 
 
 database/sql   >   As a connector to the Postgres db /n
 encoding/json  >   work with objects in json format 
 log to         >   To log errors 
 net/http       >   It will handle http requests 
 os             >   Handle environment variables

The struct defined is for an User with an Id (autoincremented by the db), a name and an email.
 
**The  application is dockerized to run**

**What actions it will perform** :

> We set an environment variable before connecting to the Postgres database.

> If a table in the database is missing, we build it.

> Define and manage the 5 endpoints used Mux.

> used port 8000, to monitor the server.

> The  function jsonMiddleware adds an application/json header to all responses.It will have the responses correctly prepared 
   and ready for usage from a future frontend.
   
> 5 controller to Creat, Read, Update and Delete users

**DOCKERIZE the App**

**create DockerFile** :

FROM sets the base image to use. In this case we are using the golang:1.16.3-alpine3.13 image, a lightweight version

WORKDIR establishes the image's working directory.

All files in the current directory are copied to the working directory with the command COPY.

RUNNING go get -d -v./... Is there a command to install the dependencies before creating the image?

COMPLETE go build -o api. create the Go application within the Image filesystem.

EXPOSE 8000 makes port 8000 accessible.

The command is set to run when the container starts by CMD ["./api"].

**Docker COMPOSE yaml file**

 we have defined 2 services, go-app and go_db  -->  go-app is the Go application we just Dockerized writing the Dockerfile , go_db is a Postgres container, to store the data. We will use the official Postgres image

1) version is the version of the docker-compose file. We are using the verwion 3.9

2) services is the list of services (containers) we want to run. In this case, we have 2 services: "go-app" and "go_db"

3) container_name is the name of the container. It's not mandatory, it's a good practice to have a name for the container. Containers find each other by their name, so it's important to have a name for the containers we want to communicate with.

4) image is the name of the image we want to use. replace "dockerhub-" with YOUR Dockerhub account.

5) build is the path to the Dockerfile. In this case, it's the current directory, so we are using .

6) environment is to define the environment variables. for the go-app, we will have a database url to configure the configuration. 
  For the go_db container, we will have the environment variables we have to define when we want to use the Postgres container (we can't change the keys here, because we are using the Postgres image, defined by the Postgres team).

7) ports is the list of ports we want to expose. In this case, we are exposing the port 8000 of the go-app container, and the port 5432 of the go_db container. The format is "host_port:container_port"

8) depends_on is the list of services we want to start before this one. In this case, we want to start the Postgres container before the app container.

9) A named volume defined in the go db will be used for persistency. As containers are by definition last for a short time, we require this additional  capability to ensure that our data persists after the container is removed.(a container is just a process).

volumes at the end of the file is the list of volumes we want to create. In this case, we are creating a volume called pgdata. The format is volume_name: {}


As dockerFile and .yaml Created Now

**Run the postgres container**

**docker compose up -d go_db**

**docker ps -a    // to see all containers**

CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS                      PORTS                                       NAMES

a9e023da5412   postgres:12                 "docker-entrypoint.s…"   45 seconds ago   Up 34 seconds               0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   go_db

connect to db localhost:5432 now

**BUILD and RUN**

// BUILD the go-app image, with the name defined in the "image" value

**docker compose build**  
**docker images**

REPOSITORY                   TAG       IMAGE ID       CREATED          SIZE
avgsoccers/go-app            1.0.0     2a15ca7c298d   23 seconds ago   312MB

**RUNNING the service**

// run a container based on the image we just built.

**docker compose up go-app**

