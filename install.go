package main

import (
	"fmt"
	"github.com/desertbit/grumble"
	"gopkg.in/AlecAivazis/survey.v1"
	"strings"
)

func init() {
	App.AddCommand(&grumble.Command{
		Name: "install",
		Help: "install docker env in the host",
		Run: func(c *grumble.Context) error {
			install()
			return nil
		},
	})
}

func install() {
	selectAvailableMirror()
	rpmPackageDownloadOnly()
	selectAvailableDockerVersion()
}

func selectAvailableMirror() {
	availableMirrors := struct {
		SelectedMirror string `survey:"Mirror"`
	}{}

	var questions = []*survey.Question{
		{
			Name: "Mirror",
			Prompt: &survey.Select{
				Message: "Choose a mirror:",
				Options: []string{"阿里云", "官方"},
				Default: "阿里云",
			},
		},
	}
	err := survey.Ask(questions, &availableMirrors)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func rpmPackageDownloadOnly() {
	downloadOnly := struct {
		SelectedMirror string `survey:"Download Only"`
	}{}

	var questions = []*survey.Question{
		{
			Name: "Download Only",
			Prompt: &survey.Select{
				Message: "Only download docker dependence rpm file?",
				Options: []string{"Yes", "No"},
			},
		},
	}
	err := survey.Ask(questions, &downloadOnly)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func selectAvailableDockerVersion() {
	cmd := getInstallDockerCmd("chooseSortedVersion")
	versions := strings.Split(executeCommand(cmd), "\n")
	versions = versions[:len(versions)-1]

	availableVersions := struct {
		SelectedVersion string `survey:"Version"`
	}{}

	var questions = []*survey.Question{
		{
			Name: "Version",
			Prompt: &survey.Select{
				Message:  "Choose a version:",
				Options:  versions,
				PageSize: 10,
			},
		},
	}
	err := survey.Ask(questions, &availableVersions)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
