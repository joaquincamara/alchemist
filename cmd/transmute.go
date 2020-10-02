package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"alchemist/reactContent"

	"github.com/spf13/cobra"
)

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
}

func createReactApp() {
	var appName string

	// Project base files
	var jsonContent = reactContent.PackageJsonContentGenerator()
	var readmeContent = reactContent.ReadmeContentGenerator()
	var gitIgnoreContent = reactContent.GitIgnoreContentGenerator()

	//SRC files

	var appCssContent = reactContent.AppCssContentGenerator()
	var appJsContent = reactContent.AppJsContentGenerator()
	var appTestJsContent = reactContent.AppTestJsContentGenerator()
	var indexCssContent = reactContent.IndexCssContentGenerator()
	var indexJsContent = reactContent.IndexJsContentGenerator()
	var logoSvgContent = reactContent.LogoSvgContentGenerator()
	var serviceWorkerContent = reactContent.ServiceWorkerContentGenerator()
	var setUpTestContent = reactContent.SetUpTestContentGenerator()

	fmt.Print("Enter the name of you app: ")
	fmt.Scanln(&appName)

	os.Mkdir(appName, 0755)
	os.Mkdir(appName+"/src", 0755)
	os.Mkdir(appName+"/puclic", 0755)

	// Project base files creation

	jsonContentErr := ioutil.WriteFile(appName+"/package.json", jsonContent, 0644)
	if jsonContentErr != nil {
		log.Fatal(jsonContentErr)
	}

	readmeContentErr := ioutil.WriteFile(appName+"/README.md", readmeContent, 0644)
	if readmeContentErr != nil {
		log.Fatal(readmeContentErr)
	}

	gitIgnoreContentErr := ioutil.WriteFile(appName+"/.gitignore", gitIgnoreContent, 0644)
	if gitIgnoreContentErr != nil {
		log.Fatal(gitIgnoreContentErr)
	}

	// SRC files creation

	appCssContentContentErr := ioutil.WriteFile(appName+"/src/App.css", appCssContent, 0644)
	if appCssContentContentErr != nil {
		log.Fatal(appCssContentContentErr)
	}

	appJsContentErr := ioutil.WriteFile(appName+"/src/App.js", appJsContent, 0644)
	if appJsContentErr != nil {
		log.Fatal(appJsContentErr)
	}

	appTestJsContentErr := ioutil.WriteFile(appName+"/src/App.test.js", appTestJsContent, 0644)
	if appTestJsContentErr != nil {
		log.Fatal(appTestJsContentErr)
	}

	indexCssContentErr := ioutil.WriteFile(appName+"/src/index.css", indexCssContent, 0644)
	if indexCssContentErr != nil {
		log.Fatal(indexCssContentErr)
	}

	indexJsContentErr := ioutil.WriteFile(appName+"/src/index.js", indexJsContent, 0644)
	if indexCssContentErr != nil {
		log.Fatal(indexJsContentErr)
	}

	logoSvgContentErr := ioutil.WriteFile(appName+"/src/logo.svg", logoSvgContent, 0644)
	if logoSvgContentErr != nil {
		log.Fatal(logoSvgContentErr)
	}

	serviceWorkerContentErr := ioutil.WriteFile(appName+"/src/serviceWorker.js", serviceWorkerContent, 0644)
	if serviceWorkerContentErr != nil {
		log.Fatal(serviceWorkerContentErr)
	}

	setUpTestContentErr := ioutil.WriteFile(appName+"/src/setupTests.js", setUpTestContent, 0644)
	if setUpTestContentErr != nil {
		log.Fatal(setUpTestContentErr)
	}

}
