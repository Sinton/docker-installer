package main

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

func init() {
	App.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println(" ____                _                 ___              _          _  _             ")
		a.Println("|  _ \\   ___    ___ | | __ ___  _ __  |_ _| _ __   ___ | |_  __ _ | || |  ___  _ __ ")
		a.Println("| | | | / _ \\  / __|| |/ // _ \\| '__|  | | | '_ \\ / __|| __|/ _` || || | / _ \\| '__|")
		a.Println("| |_| || (_) || (__ |   <|  __/| |     | | | | | |\\__ \\| |_| (_| || || ||  __/| |   ")
		a.Println("|____/  \\___/  \\___||_|\\_\\\\___||_|    |___||_| |_||___/ \\__|\\__,_||_||_| \\___||_|   ")
		a.Println("")
	})
}

var App = grumble.New(&grumble.Config{
	Name:                  "installer",
	Description:           "Docker auto install tool",
	Prompt:                "exec » ",
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})
