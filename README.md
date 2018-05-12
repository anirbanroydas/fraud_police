# fraud_police

A sample microservice written in Golang which takes request(order transaction details) via its exposed http api and returns if a transaction is fraudulent or not. This is used along with the dummy_orders and alertman service.


## Details

* **Author :**  Anirban Roy Das        
* **Email  :**  anirban.nick@gmail.com 


## Features

* Golang microservice
* Http api
* Sample app used by [dummy_orders](htpps://github.com/anirbanroydas/dummy_orders) along with [alertman](https://github.com/anirbanroydas/alertman)
* A flavour of Clean architecture by Robert Martin aka Uncle Bob


## Overview

* **Some Code Design Choice Specific Note:**

  I have tried to structure the code using [**Clean Architecture**](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) proposed by [**Robert Martin**](ttps://en.wikipedia.org/wiki/Robert_C._Martin) famously known as **Uncle Bob**.

  **Clean Architecture** is some better or flavour of other known architectures like [Porst & Adapters](https://spin.atomicobject.com/2013/02/23/ports-adapters-software-architecture/), 
  [The Onion Architecture](http://jeffreypalermo.com/blog/the-onion-architecture-part-1/), etc.

  Clean architecture mainly focusses on following one principle strictly that is **Dependency Inversion**. Keeping all dependencies flow in a uni direction 
  makes it quite quite powerful. Infact, I have finally realized the value of proper **dependency injections** while implementing clean architecture.

  **NOTE :** This is not the best architecture for all usecases and of course a little more verbose and more boilerplate than some other design patterns, but it 
  does help you keep you codebase fully maintainable for the long run. You may not agree with Clean architecture's philosophy sometimes. But I am just using it to understand it more.

  Btw, [Robert Martin](ttps://en.wikipedia.org/wiki/Robert_C._Martin) is also known for the [**S.O.L.I.D**](https://medium.com/@cramirez92/s-o-l-i-d-the-first-5-priciples-of-object-oriented-design-with-javascript-790f6ac9b9fa) principles which have shaped real greate design choices when it comes to writing good maintainable **Object Oriented Code**. But, ***SOLID*** is talked about even in non object oriented langauges like **Golang**.

  >**P.S :** Just for motivational purposes, I am implementing few projects in ***Clean Architecture*** to understand it better and challenge the norm or my own design choices.

  > This is a good talk by the same guy [**link to talk**](https://www.youtube.com/watch?v=o_TH-Y78tt4)


* **Service Details**

  fraud_police is a very simple http server based service which exposes one POST api which accepts
  some transaction related information and processes the transaction.

  It then returns isFraud true or false dependending on the transaction. But for simplicity and demo purposes there is no actual fraud processing that happens. This processing is dummied by a very random sleep and return true or false again randomly (may be 1 out of 10 times false).

  The main idea is to help the [dummy_orders](https://github.com/anirbanroydas/dummy_orders) service to hit this **fraud_police** service and mimic some kind of *microservice to microservice* communication.

  This service although has some proper interfaces which can be implemented by concrete glang structures. This shows the power of Go interfaces and also the Clean Architecture which helps you
  write a generic code which adhers to the SOLID principles (majorly the Open/Closed principle).
  The code is extensible and by added the real fraud processing logic or any real implementation will not change other parts or other domain or business logic.

  To make the application complete and work properly, I have added few dummy structures and very easy implementations which you will know if you read the code. I have documented them in the code properly. Thats it. No other usecase for thise service.

  **NOTE :** Although *Golang's* powerful concept of [**Interfaces**](https://www.youtube.com/watch?v=F4wUrj6pmSI&t=2489s) which not only emphasize more on the [Composition over Inheritance](https://www.youtube.com/watch?v=wfMtDGfHWpA) principle but also helps in writing *Clean architecture* code more straight forward. But since it also is more boilerplate so sometimes the code looks a little verbose but it actually isn't. Its is great for maintainability and extensibility.

## Technical Specs

* **Go 1.10 :** Golang programming language
* **gin :** HTTP framework for rest apis
* **Docker :** A containerization tool for better devops


## Deployment

There are two ways to deploy:

* using [Docker](https://www.docker.com/)
* via direct compilation and running the binary


#### Prerequisite 

* **Required**

  Copy (not move) the ``env`` file in the root project directory to ``.env`` and add/edit 
  the configurations accordingly.

  This needs to be done because the server, or docker deployment, or some script may want some pre configurations like ports, 
  hostnames, etc before it can start the service, or deploy the service or may be to run some scripts.

* **Optional**

  To safegurad secret and confidential data leakage via your git commits to public 
  github repo, check ``git-secrets``.

  This `git secrets <https://github.com/awslabs/git-secrets>`_ project helps in 
  preventing secrete leakage by mistake.


#### Direct compilation and running the binary

* Just go to the `cmd/fraud_police_server` directory and compile the `main.go`
    
        $ cd cmd/fraud_police_server
        $ go install

* Then run the binary to start the server.

        $ fraud_police_server



#### Using Docker

* **Step 1:**
    
  Install **docker** and **make** command if you don't have it already.

  * Install Docker
    
    Follow my another github project, where everything related to DevOps and scripts are 
    mentioned along with setting up a development environemt to use Docker is mentioned.

    * Project: https://github.com/anirbanroydas/DevOps

    * Go to setup directory and follow the setup instructions for your own platform, linux/macos

  * Install Make
            
        # (Mac Os)
        $ brew install automake

        # (Ubuntu)
        $ sudo apt-get update
        $ sudo apt-get install make

* **Step 2:**

  There is a ``Makefile`` present in the root project directory which actually hides
  away all the docker commands and other complex commands. So you don't have to actually 
  know the **Docker** commands to run the service via docker. **Make** commands will do the
  job for you.

  * Make sure the ``env`` file has been copied to ``.env`` and necessary configuration changes done.
  * There are only two values that need to be taken care of in the ``Makefile``

    * BRANCH: Change this to whatever branch you are in if making changes and creating the docker images again.
    * COMMIT = Change this to a 6 char hash of the commit value so that the new docker images can be tracked.

  * Run the command to start building new docker image and push it to docker hub.
        
    * There is a script called ``build_tag_push.sh`` which actually does all the job of building the image, tagging the image ans finally pushing it to the repository.
    * Make sure you are logged into to your docker hub acount. 
    * Currently the ``build_tag_push.sh`` scripts pushes the images to ``hub.docker.com/aroyd`` acount. Change the settings in that file if you need to send it to some other place.
    * The script tags the new built docker image with the branch, commit and datetime value.
    * To know more, you can read the ``Dockerfile`` to get idea about the image that gets built on runing this make command.

        
            $ make build-tag-push

* **Step 3:**

  Pull the image or run the image separately or you can run it along with other services, docker containers etc.
  To know about this, check the sample [dummy_orders](htpps://github.com/anirbanroydas/dummy_orders) service which makes use of this **fraud_police** servic.
    
  That service has a well defined ``docker-compose.yml`` file which explains the whole setup process to make the **fraud_police** service work/communicate with other services.


## Usage

Check the above **Step 3** which will direct you to a place on how to use it. There is not API as such to know what and how request are processed, for now just go through the code. Docs may be added later for detail description.


