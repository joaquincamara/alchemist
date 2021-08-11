package templates

import (
	"io/ioutil"
	"log"
)

var gitIgnoreTemplate = []byte(`
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with: go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

*.env
`)

func CreateAlchemistGitignore() {
	err := ioutil.WriteFile(".gitignore", gitIgnoreTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
