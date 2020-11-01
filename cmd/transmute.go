package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joaquincamara/alchemist/nestContent"
	"github.com/joaquincamara/alchemist/reactContent"
	"github.com/joaquincamara/alchemist/server"
	"github.com/joaquincamara/alchemist/utils"
	"github.com/spf13/cobra"
)

var transmuteCmd = &cobra.Command{
	Use:   "transmute",
	Short: "Init the setup configuration for a alchemist project",
	Long:  `Init the setup configuration for a alchemist project with React.js for the client and Golang for the backend`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("")
			fmt.Println("Uups looks like something get wrong. Use any of the next commands:")
			fmt.Println("")
			fmt.Println("> alchemist transmute stone my-app: create a alchemist server and React.js app")
			fmt.Println("")
			fmt.Println("> alchemist transmute monolit my-app: create a monolite Nest.js and React.js app")
			fmt.Println("")
			fmt.Println("> alchemist transmute react my-app: create a stand alone React.js app")
			fmt.Println("")
			fmt.Println("> alchemist transmute nest my-app: create a stand alone Nest.js app")
			fmt.Println("")
		}

		if len(args) != 2 && len(args) != 0 {
			fmt.Println("Insert the name of your App after the alchemist sub-command: 'alchemist transmute [sub-command] my-app'")
		} else {
			switch {
			case "react" == args[0]:
				reactContent.CreateReactApp(args[1])
				break
			case "nest" == args[0]:
				nestContent.CreateNestApp(args[1])
				break
			case "monolit" == args[0]:
				createMonolit(args[1])
				break
			case "stone" == args[0]:
				createStone(args[1])
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(transmuteCmd)
}

// Creates the folders and files for a React and Go ApiRest application
func createStone(appName string) {
	os.Mkdir(appName, 0755)
	server.CreateAlchemistServer(appName + "/server")
	reactContent.CreateReactApp(appName + "/client")

	alchemistJsonErr := ioutil.WriteFile(appName+"/alchemist.json", utils.AlchemistJson(appName, "monolit"), 0644)
	if alchemistJsonErr != nil {
		log.Fatal(alchemistJsonErr)
	}
}

// Creates the folders and files for a Nest and React application
func createMonolit(appName string) {
	os.Mkdir(appName, 0755)
	nestContent.CreateNestApp(appName + "/server")
	reactContent.CreateReactApp(appName + "/client")

	alchemistJsonErr := ioutil.WriteFile(appName+"/alchemist.json", utils.AlchemistJson(appName, "monolit"), 0644)
	if alchemistJsonErr != nil {
		log.Fatal(alchemistJsonErr)
	}
}
