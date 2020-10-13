  <img  src="https://user-images.githubusercontent.com/26718123/94979651-96461b00-04e9-11eb-94d6-660af9663975.png">

Alchemist.go it's a command cli powered by Go, it's commands allow you to transmute projects with React.js for the client app and Go for the backend api.

### Requirements

- Go 1.15
- Node.js 12.19 or higher.
- Npm or Yarn to install React.js and Nest.js dependencies.

### Installation

- go get github.com/spf13/cobra

### Commands

To invoque any of the ancestral spells of our book of Alchemy we need to use the special word "alchemist" at the terminal, This gonna open the door to the world of alchemy and bring all the resources to transmute a new alchemist's project.

- **\$alchemist transmute** : this command is going to start the transmutation of a new alchemist's project, with a React.js app and a base go server. **NOTE: The development of the base server go is still on work by the moment run this code is going to return a "Still working on this spell." message.**

- **\$alchemist transmute monolit my-app** : this command is going to start the transmutation of a new monolit app with a nest.js app based on the https://github.com/nestjs/typescript-starter.git configuration for the server and a React.js app based on the create-react-app configuration for the client.

- **\$alchemist transmute react my-app** : this command is going to start the transmutation of a new stand alone React.js app based on the create-react-app configuration.

- **\$alchemist transmute nest my-app** : this command is going to start the transmutation of a new stand alone nest.js app based on the https://github.com/nestjs/typescript-starter.git project configuration.

**NOTE:** At the 0.1.0 version the transmute command is just able to setup the neccesary files to duplicate a "create react app" base project. You need to run "Yarn" or "npm install" to install the dev dependencies and run the React.js project.

### Future plans:

- Add a go RestFul server and a command to transmute it.
- Add a alchemist command to transmute a React.js and Go restful app.
- Add a alchemist command for hot reloading at all our apps.

### Dependencies:

**Our transmutations are made possible by Nest.js, react.js and cobra.go**

======================

**Alchemist.go version: 0.1.0**

======================
