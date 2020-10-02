/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// transmuteCmd represents the transmute command
var transmuteCmd = &cobra.Command{
	Use:   "transmute",
	Short: "Init the setup configuration for a alchemist project",
	Long:  `Init the setup configuration for a alchemist project with React.js for the client and Golang for the backend`,
	Run: func(cmd *cobra.Command, args []string) {
		createReactApp()
	},
}

func init() {
	rootCmd.AddCommand(transmuteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transmuteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transmuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createReactApp() {
	fmt.Print("Enter the name of you app: ")
	var appName string
	fmt.Scanln(&appName)
	//fmt.Print(appName)
	os.Mkdir(appName, 0755)
	os.Mkdir(appName+"/src", 0755)
	os.Mkdir(appName+"/puclic", 0755)
	message := []byte(`{
		"name": "my-app",
		"version": "0.1.0",
		"private": true,
		"dependencies": {
			"@testing-library/jest-dom": "^4.2.4",
			"@testing-library/react": "^9.3.2",
			"@testing-library/user-event": "^7.1.2",
			"react": "^16.13.1",
			"react-dom": "^16.13.1",
			"react-scripts": "3.4.3"
		},
		"scripts": {
			"start": "react-scripts start",
			"build": "react-scripts build",
			"test": "react-scripts test",
			"eject": "react-scripts eject"
		},
		"eslintConfig": {
			"extends": "react-app"
		},
		"browserslist": {
			"production": [
				">0.2%",
				"not dead",
				"not op_mini all"
			],
			"development": [
				"last 1 chrome version",
				"last 1 firefox version",
				"last 1 safari version"
			]
		}
	}
	`)
	err := ioutil.WriteFile(appName+"package.json", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
