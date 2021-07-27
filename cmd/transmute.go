package cmd

import (
	"fmt"
	"os"

	"github.com/joaquincamara/alchemist/server"
	"github.com/spf13/cobra"
)

var transmuteCmd = &cobra.Command{
	Use:   "transmute",
	Short: "Init the setup configuration for a alchemist project web service",
	Long:  `Init the setup configuration for a alchemist project with React.js for the client and Golang for the backend`,
	Run: func(cmd *cobra.Command, args []string) {

		switch len(args) {
		case 0:
			fmt.Println("")
			fmt.Println("Uups looks like something get wrong. Use any of the next commands:")
			fmt.Println("")
			fmt.Println("Insert the name of your App after the alchemist sub-command: 'alchemist transmute  my-app'")
		case 1:
			createStone(args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(transmuteCmd)
}

// Creates the folders for a alchemist web service
func createStone(appName string) {
	os.Mkdir(appName, 0755)
	server.CreateAlchemistServer(appName + "/server")
}
