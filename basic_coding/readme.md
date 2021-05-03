# Microservice go app


## Run using docker

* Ensure that docker and docker-compose are installed. for [docker](https://docs.docker.com/engine/install/) and for [docker-compose](https://docs.docker.com/compose/install/)
* Type command "docker-compose up" to make images of all services and container running

```
$ docker-compose up
```
* Ensure that app running, 
golang app is on port 8013
python app is on port 8014

## Run using golang

* Navigate to service (supplier-service/user-service)
* Ensure golang installed on local machine, run this command on terminal
```
$ go run main.go
```

## Run app with .exe 

* Navigate to service (supplier-service/user-service)
* Double click main.exe

## Technical flow and design system 

* For each of the requirement will be 1 EP
* There will be 2 service user and supplier service which contain several end point 
* To make development fast, sql lite was used on this project 

