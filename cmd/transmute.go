package cmd

import (
	"alchemist/nestContent"
	"alchemist/reactContent"
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

var transmuteCmd = &cobra.Command{
	Use:   "transmute",
	Short: "Init the setup configuration for a alchemist project",
	Long:  `Init the setup configuration for a alchemist project with React.js for the client and Golang for the backend`,
	Run: func(cmd *cobra.Command, args []string) {

		switch {
		case len(args) == 0:
			fmt.Println("")
			fmt.Println("Still working on this spell. Use any of the next commands:")
			fmt.Println("")
			fmt.Println("> alchemist transmute react my-app: create a stand alone React.js app")
			fmt.Println("")
			break
		case "react" == args[0]:
			if len(args) == 2 {
				reactContent.CreateReactApp(args[1])
				break
			} else {
				fmt.Println("Insert the name of your React-app after the 'react' command: 'alchemist transmute react my-app'")
			}
			break
		case "nest" == args[0]:
			if len(args) == 2 {
				nestContent.CreateNestApp(args[1])
				break
			} else {
				fmt.Println("Insert the name of your Nest-app after the 'nest' command: 'alchemist transmute nest my-app'")
			}
			break
		case "monolit" == args[0]:
			if len(args) == 2 {
				createMonolit(args[1])
				break
			} else {
				fmt.Println("Insert the name of your App after the 'monolit' command: 'alchemist transmute monolit my-app'")
			}
			break
		}
	},
}

func init() {
	rootCmd.AddCommand(transmuteCmd)
}

func createMonolit(appName string) {
	os.Mkdir(appName, 0755)
	nestContent.CreateNestApp(appName + "/client")
	reactContent.CreateReactApp(appName + "/server")
}
