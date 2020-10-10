package reactContent

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateReactApp(appName string) {
	// Project base files
	var jsonContent = PackageJsonContentGenerator()
	var readmeContent = ReadmeContentGenerator()
	var gitIgnoreContent = GitIgnoreContentGenerator()

	//SRC files

	var appCssContent = AppCssContentGenerator()
	var appJsContent = AppJsContentGenerator()
	var appTestJsContent = AppTestJsContentGenerator()
	var indexCssContent = IndexCssContentGenerator()
	var indexJsContent = IndexJsContentGenerator()
	var logoSvgContent = LogoSvgContentGenerator()
	var serviceWorkerContent = ServiceWorkerContentGenerator()
	var setUpTestContent = SetUpTestContentGenerator()

	// PUBLIC files

	var indexHtmlContent = IndexHtmlContentGenerator()
	var manifestJsonContent = ManifestJsonContentGenerator()
	var robotsTxtContent = RobotsTxtContentGenerator()

	os.Mkdir(appName, 0755)
	os.Mkdir(appName+"/src", 0755)
	os.Mkdir(appName+"/public", 0755)

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

	// SRC directory files creation

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

	// PUBLIC directory files creation

	indexHtmlContentErr := ioutil.WriteFile(appName+"/public/index.html", indexHtmlContent, 0644)
	if indexHtmlContentErr != nil {
		log.Fatal(indexHtmlContentErr)
	}

	manifestJsonContentErr := ioutil.WriteFile(appName+"/public/manifest.json", manifestJsonContent, 0644)
	if indexHtmlContentErr != nil {
		log.Fatal(manifestJsonContentErr)
	}

	robotsTxtContentErr := ioutil.WriteFile(appName+"/public/robots.txt", robotsTxtContent, 0644)
	if indexHtmlContentErr != nil {
		log.Fatal(robotsTxtContentErr)
	}
}
