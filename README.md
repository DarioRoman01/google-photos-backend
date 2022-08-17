# Google Photos Backend

## What is this?
this a backend based on microservices architecture for Google Photos like service.
it uses s3 as a storage for images, kubernetes as a microservice for processing images and rds to store metadata.

## How it works?

### Storage
images are stored in s3 bucket.

### Processing
images are processed by kubernetes pods.

## microservices 
its is compose of the following microservices:
* upload service:
    * uses grpc to recibe the data and store it in s3

* mail service:
    * uses grpc to recibe the data and send emails to the user

* command service:
    * a rest services that handles all write actions related to the images and users

* query service:
    * a rest services that handles all read actions related to the images and users

* nginx:
    * a reverse proxy that forwards the requests to the microservices


## Features

* upload images
* delete images
* update images
* oder images by folder

## How to run it?
the only requisit to run it is to have a mailgun account and a s3 bucket with the right permissions.
check the .env.example file for setting up the env variables.

to try it locally you can run the following commands:

```bash
docker-compose up -d --build
```

then you can access the service at http://localhost:8000

to run it on kubernetes check the services folder there you will find the kubernetes deployment and service files.
so you just have to run the following command:

**note you need to have at least 2 nodes in your cluster**


```bash
kubectl apply -f service/deployment.yaml
kubectl apply -f service/service.yaml
```