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
		templates.CreateAlchemistMain()
		templates.CreateAlchemistGitignore()
		templates.CreateAlchemistReadme()

		ct.Foreground(ct.Green, false)
		fmt.Fprintln(writer, "Transmutation Complete")
		writer.Print()
		ct.ResetColor()
		time.Sleep(time.Second)
		templates.CreateDomainsFolders(projectStructure)
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

func CreateGoModFile() {
	cmnd := "go mod init"

	//value := "example.com/joaquincamara/" + appName

	cmd := exec.Command(cmnd)
	//cmd.Path = "./usr/local/go/bin"
	//cmd.Dir = appName

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(stdout))
}

/*	cmnd := "go mod init"

	//value := "example.com/joaquincamara/" + appName

	cmd := exec.Command(cmnd)
	cmd.Path = "./usr/local/go/bin"
	cmd.Dir = appName

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(stdout))*/
