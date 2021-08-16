package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/apoorvam/goterminal"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/joaquincamara/alchemist/templates"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var transmuteCmd = &cobra.Command{
	Use:   "transmute",
	Short: "Create an alchemist project web service",
	Run: func(cmd *cobra.Command, args []string) {

		CreateAlchemistProject()
	},
}

func init() {
	rootCmd.AddCommand(transmuteCmd)
}

func CreateAlchemistProject() {

	writer := goterminal.New(os.Stdout)
	ct.Foreground(ct.Red, false)
	fmt.Fprintln(writer, "Opening the Alchemy Door")
	writer.Print()
	ct.ResetColor()
	time.Sleep(time.Second)

	ct.Foreground(ct.Cyan, false)
	fmt.Fprintln(writer, "Transmuting your Alchemist project")
	writer.Print()
	ct.ResetColor()
	time.Sleep(time.Second)

	projectStructure, err := readAlchemistYaml()
	if err != nil {
		fmt.Println(err)
	} else {

		os.Mkdir("api", 0755)
		os.Mkdir("cmd", 0755)
		os.Mkdir("internal", 0755)
		os.Mkdir("public", 0755)
		CreateGoModFile(projectStructure.AppName)
		runGoGetCommand(projectStructure.AppName)
		templates.CreateAlchemistMain()
		templates.CreateAlchemistGitignore()
		templates.CreateAlchemistReadme()
		templates.CreateIndexHTML()
		templates.CreateDomainsFolders(projectStructure)

		ct.Foreground(ct.Green, false)
		fmt.Fprintln(writer, "Transmutation Complete")
		writer.Print()
		ct.ResetColor()
		time.Sleep(time.Second)
	}

}

func readAlchemistYaml() (templates.AlchemistYAML, error) {

	f, err := os.Open("alchemist.yaml")
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)

	var yamlFile templates.AlchemistYAML
	err = dec.Decode(&yamlFile)
	if err != nil {
		return yamlFile, err
	}

	return yamlFile, nil
}

// add err management
func CreateGoModFile(appName string) {

	args := []string{
		"mod",
		"init",
	}
	args = append(args, appName)
	e := exec.Command("go", args...)
	e.Output()

}

// add err management
func runGoGetCommand(appName string) {
	args := []string{
		"get",
		"github.com/joaquincamara/silver",
	}
	e := exec.Command("go", args...)
	e.Output()
}
