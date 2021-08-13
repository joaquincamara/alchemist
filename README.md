  <img  src="https://user-images.githubusercontent.com/26718123/94979651-96461b00-04e9-11eb-94d6-660af9663975.png">

Alchemist.go it's a command cli powered by Go, it's command allow you to transmute an Alchemist web service with Go, based on the hexagonal architecture. Our transmutation command create for you all the handlers, services, models and repositories.

## Requirements

- Go 1.15 or higher

## Installation

In order to install alchemist cli, run this command at your terminal in the root:

`go get github.com/joaquincamara/alchemist@latest`

## Commands

To invoque any of the ancestral spells of our book of Alchemy we need first to prepare our transmutation circle:

1. mkdir myApp && cd myApp
2. Create a alchemist.yaml file with the next structure:

```yaml
name: myApp
internal: 
  domainName: 
    model:
      Property: int `json:"Property,omitempty"`
      Property: string `json:"Property"`
      Property: string `json:"Property"`
projectVersion: v0.1.0
```

Inside the internal section you need to specify the name of all the domains that your application is going to have and also the name of the properties of each domain.

**Example:**

```yaml

name: curriculum
internal: 
  aboutMe: 
    model:
      Id: int `json:"Id,omitempty"`
      Info: string `json:"Title"`
      Title: string `json:"Info"`
  coolFeatures: 
    model: 
      Id: int `json:"Id,omitempty"`
      Name: string `json:"Name"`
projectVersion: v0.1.0
```
3. Use the special word "alchemist" at the terminal in you working directory, This gonna open the door to the world of alchemy and bring all the resources to transmute a new alchemist's project:

- `$ alchemist transmute`

**Alchemist project example:**

```
curriculum
│   README.md
│   alchemist.yaml    
│   .gitignore
|
└───api
│   │   aboutMeHandler.go
│   │   coolFeaturesHandler.go
│   
└───cmd
|	│   main.go
│
└───internal
    │   
	└───aboutMe
    │	│   model.go
	│   │	repository.go
	│   │	service.go
	│
	└───cmd
    	│   model.go
	    │	repository.go
	    │	service.go
```

## Alchemist Router

We know that sometimes you do not need all our transmutations to achieve your projects goals, so if you want to create a router without our base project, you can use just the Silver.go router library for your transmutations.

In order to use the alchemist server library, just run in you project folder:

`go get github.com/joaquincamara/silver`

### Simple Silver server

```golang
package main

import (
"log"
"net/http"
alchemy "github.com/joaquincamara/silver"
)

func main() {
	router := silver.NewRouter()

	router.Use(middlewares.Recovery)

	router.GET("/", silver.AlchemyDoor)
	silver.Start("8080", router)
}
```

Silver.go website: [Silver.go](https://github.com/joaquincamara/silver).

**we know, we know, we are still in baby steps but fraid not, we are working in new transmutatios**


### Future plans:

- Add cli command to create a basic database connection.


### Dependencies:

**Our transmutations are possible by:**

- Silver.go
- cobra.go

======================

**My eternal gratitude to my mentor Eduardo Villaseñor** https://github.com/eduvim

======================

**Alchemist.go version: 0.1.3**

======================
