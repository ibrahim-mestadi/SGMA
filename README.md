## What I'm building here : 
A front end web app that connects to 5 microservices :

* [ ] <span style="font-weight:bold">Brocker</span> : Optional single point of entry to microservices, nothing just recieving requests from front, firing off requests & consumer responses ! 
* [ ] <span style="font-weight:bold">Authentication :  </span> Postgress
* [ ] <span style="font-weight:bold">Logger</span> : MongoDB
* [ ] <span style="font-weight:bold">Mail</span> : send email with a specific template 
* [ ] <span style="font-weight:bold">Listneer </span> : consumes messages in in RabbitMQ & initiates a process

## Communication :

* [ ] I will start by Connecting those microservices using Rest API with JSON as transport 
* [ ] Then, Play with RPC, gRPC and testing performences 
* [ ] instantiate and responding to events using AMQP  

## Wirframe : 

![](assets/WireframesSGMA.png)