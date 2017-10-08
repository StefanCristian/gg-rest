package imports

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"datamodel"
)


// Naming convention is subject to change soon enough
// This method implements importing of gentoo_commands json
// The commands json will have the desired gentoo-based commands that will be used
// in the future to install packages in the same manner as a standard gentoo-based distro
// For the moment it's just in the ALPHA, it may be subject to modularization also.
func ImportCommands() []datamodel.GentooCommands {
	importedcommands, err := ioutil.ReadFile("src/imports/gentoo_commands.json")
	if err != nil {
		fmt.Printf("Failed to read gentoo_commands.json file: %s\n", err)
		os.Exit(1)
	}
	var commands []datamodel.GentooCommands
	err = json.Unmarshal(importedcommands, &commands)
	if err != nil {
		fmt.Printf("Failed to unmarshal gentoo commands: %s\n", err)
		os.Exit(1)
	}
	return commands
}

// Imported Arguments are subject to change soon enough
// They implement future functions related to the commands sections
// Not used for the moment
func ImportArguments() []datamodel.GentooArguments {
	importedArguments, err := ioutil.ReadFile("src/imports/gentoo_arguments.json")
	if err != nil {
		fmt.Printf("Failed to read gentoo_arguments.json file: %s\n", err)
		os.Exit(1)
	}
	var arguments []datamodel.GentooArguments
	err = json.Unmarshal(importedArguments, &arguments)
	if err != nil {
		fmt.Printf("Failed to unmarshal gentoo arguments: %s\n", err)
		os.Exit(1)
	}
	return arguments
}
