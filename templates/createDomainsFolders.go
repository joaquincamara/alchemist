package templates

import (
	"os"
)

type AlchemistYAML struct {
	Internal map[string]map[string]map[string]string `yaml:"internal"`
	AppName  string                                  `yaml:"appName"`
}

func CreateDomainsFolders(structure AlchemistYAML) {
	for k, v := range structure.Internal {
		os.Mkdir("internal/"+k, 0755)

		CreateModelFile(k, v["model"])
		CreateServiceFile(k)
		CreateRespositoryFile(k)
		CreateApiHandlers(k, structure.AppName)

	}
}
