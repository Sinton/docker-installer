package main

import "github.com/desertbit/grumble"

func init() {
	App.AddCommand(&grumble.Command{
		Name: "validate",
		Help: "validate env",
		Run: func(c *grumble.Context) error {
			validate()
			return nil
		},
	})
}

func validate() {
	validateOs()
}

func validateOs() error {
	getValidateEnvCmd("osType")
	return nil
}
