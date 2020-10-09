package cmd

import (
	"alchemist/reactContent"
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
				fmt.Println("Insert the name of your React-app after the 'React' command: 'alchemist transmute react my-app'")
			}
			break
		}
	},
}

func init() {
	rootCmd.AddCommand(transmuteCmd)
}
