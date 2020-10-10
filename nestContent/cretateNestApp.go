package nestContent

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateNestApp(appName string) {

	os.Mkdir(appName, 0755)
	os.Mkdir(appName+"/src", 0755)
	os.Mkdir(appName+"/test", 0755)

	// Project base files creation

	eslintrcErr := ioutil.WriteFile(appName+"/.eslintrc.js", Eslintrc(), 0644)
	if eslintrcErr != nil {
		log.Fatal(eslintrcErr)
	}

	gitIgnoreErr := ioutil.WriteFile(appName+"/.gitignore", GitIgnore(), 0644)
	if gitIgnoreErr != nil {
		log.Fatal(gitIgnoreErr)
	}

	prettierrcErr := ioutil.WriteFile(appName+"/.prettierrc", Prettierrc(), 0644)
	if prettierrcErr != nil {
		log.Fatal(prettierrcErr)
	}

	nestCliJsonErr := ioutil.WriteFile(appName+"/nest-cli.json", NestCliJson(), 0644)
	if nestCliJsonErr != nil {
		log.Fatal(nestCliJsonErr)
	}

	packageLockJsonErr := ioutil.WriteFile(appName+"/package-lock.json", PackageLockJson(), 0644)
	if packageLockJsonErr != nil {
		log.Fatal(packageLockJsonErr)
	}

	packageJsonErr := ioutil.WriteFile(appName+"/package.json", PackageJson(), 0644)
	if packageJsonErr != nil {
		log.Fatal(packageJsonErr)
	}

	readmeErr := ioutil.WriteFile(appName+"/README.md", Readme(), 0644)
	if readmeErr != nil {
		log.Fatal(readmeErr)
	}

	tsConfigBuildJsonErr := ioutil.WriteFile(appName+"/tsconfig.build.json", TsConfigBuildJson(), 0644)
	if tsConfigBuildJsonErr != nil {
		log.Fatal(tsConfigBuildJsonErr)
	}

	tsConfigErr := ioutil.WriteFile(appName+"/tsconfig.json", TsConfig(), 0644)
	if tsConfigErr != nil {
		log.Fatal(tsConfigErr)
	}

	// SRC directory files creation

	appControllerSpecErr := ioutil.WriteFile(appName+"/src/app.controller.spec.ts", AppControllerSpec(), 0644)
	if appControllerSpecErr != nil {
		log.Fatal(appControllerSpecErr)
	}

	appControllerErr := ioutil.WriteFile(appName+"/src/app.controller.ts", AppController(), 0644)
	if appControllerErr != nil {
		log.Fatal(appControllerErr)
	}

	appModuleErr := ioutil.WriteFile(appName+"/src/app.module.ts", AppModule(), 0644)
	if appModuleErr != nil {
		log.Fatal(appModuleErr)
	}

	appServiceErr := ioutil.WriteFile(appName+"/src/app.service.ts", AppService(), 0644)
	if appServiceErr != nil {
		log.Fatal(appServiceErr)
	}

	// TEST directory files creation

	appE2eSpecErr := ioutil.WriteFile(appName+"/test/app.e2e-spec.ts", AppE2eSpec(), 0644)
	if appE2eSpecErr != nil {
		log.Fatal(appE2eSpecErr)
	}

	jestE2eErr := ioutil.WriteFile(appName+"/test/jest-e2e.json", JestE2e(), 0644)
	if jestE2eErr != nil {
		log.Fatal(jestE2eErr)
	}
}
