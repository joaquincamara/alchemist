  <img  src="https://user-images.githubusercontent.com/26718123/94979651-96461b00-04e9-11eb-94d6-660af9663975.png">

Alchemist.go it's a command cli powered by Go, it's commands allow you to transmute projects with React.js for the client app and Go for the backend api.

### Requirements

- Go 1.15
- Node.js 12.19 or higher.
- Npm or Yarn to install React.js and Nest.js dependencies.

### Installation

In order to use the alchemist command, compile it using the following command:

`go get github.com/joaquincamara/alchemist`

### Commands

To invoque any of the ancestral spells of our book of Alchemy we need to use the special word "alchemist" at the terminal, This gonna open the door to the world of alchemy and bring all the resources to transmute a new alchemist's project.

- `$ alchemist transmute stone my-app` : this command is going to start the transmutation of a new alchemist's project, with a React.js app and RestApi go server.

- `$ alchemist transmute monolit my-app` : this command is going to start the transmutation of a new monolit app with a nest.js app based on the https://github.com/nestjs/typescript-starter.git configuration for the server and a React.js app based on the create-react-app configuration for the client.

- `$ alchemist transmute react my-app` : this command is going to start the transmutation of a new stand alone React.js app based on the create-react-app configuration.

- `$ alchemist transmute nest my-app` : this command is going to start the transmutation of a new stand alone nest.js app based on the https://github.com/nestjs/typescript-starter.git project configuration.

### Server library

We know that sometimes you do not need all our transmutations to achieve your projects goals, so if you only need to create a server without our base project, you can use just the server library.

In order to use the alchemist server library, just run:

`go get github.com/joaquincamara/alchemist/server`

### Simple Alchemist server

Here is a simple implementation of simple router with the Alchemist server library:

```golang
package main

import (
"log"
"net/http"
alchemy "github.com/joaquincamara/alchemist/server"
)

func main() {
router := alchemy.NewRouter()
      router.GET("/", alchemy.HomeRoute)
      log.Fatal(http.ListenAndServe(":8080", router))
}
```

**Alchemist server methods**

- GET: `router.GET("/")`

- PUT: `router.PUT("/")`

- DELETE: `router.DELETE("/")`

- PUT: `router.PUT("/")`

**we know, we know, still baby steps but fraid not, we are working in new transmutatios for the server library**

**NOTE:** At the 0.1.0 version the transmute command is just able to setup the neccesary files to duplicate a "create react app" base project. You need to run "Yarn" or "npm install" to install the dev dependencies and run the React.js project.

### Future plans:

- Add a go RestFul server and a command to transmute it.
- Add a alchemist command to transmute a React.js and Go restful app.
- Add a alchemist command for hot reloading at all our apps.

### Dependencies:

**Our transmutations are possible by:**

- Nest.js
- React.js and create-react-app
- cobra.go
- Server code reference to: Chris Gregori

======================

**Alchemist.go version: 0.1.1**

======================
