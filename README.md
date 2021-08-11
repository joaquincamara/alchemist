  <img  src="https://user-images.githubusercontent.com/26718123/94979651-96461b00-04e9-11eb-94d6-660af9663975.png">

Alchemist.go it's a command cli powered by Go, it's commands allow you to transmute an Alchemist web service with Go.

## Requirements

- Go 1.15

## Installation

In order to install alchemist cli, run in your terminal at the root:

`go get github.com/joaquincamara/alchemist@latest`

## Commands

To invoque any of the ancestral spells of our book of Alchemy we need to use the special word "alchemist" at the terminal, This gonna open the door to the world of alchemy and bring all the resources to transmute a new alchemist's project.

- `$ alchemist transmute my-app` : this command is going to start the transmutation of a new alchemist's web service with go.

## Alchemist Router

We know that sometimes you do not need all our transmutations to achieve your projects goals, so if you want to create a router without our base project, you can use just the Silver.go router library for you transmutations.

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

**we know, we know, still baby steps but fraid not, we are working in new transmutatios for the server library**


### Future plans:

- Add a cli command to create models, services, controllers and handlers.
- Add cli command to create a basic database connection.

### Dependencies:

**Our transmutations are possible by:**

- Silver.go
- cobra.go

======================

**Alchemist.go version: 0.1.3**

======================
