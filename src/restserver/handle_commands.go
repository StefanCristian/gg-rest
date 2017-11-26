package restserver

import (
	"net/http"
	"datamodel"
	"imports"
	"fmt"
	"os"
	"os/exec"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type GCommands []datamodel.GentooCommands
type PkgLists []datamodel.PackagesInputs

var commands GCommands = imports.ImportCommands()

func GetAllCommands(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, commands, "")
}

func (a *GCommands) get(firstCommand string) (datamodel.GentooCommands, error) {
	err := fmt.Errorf("could not find commands %s\n", firstCommand)
	for _, commandz := range *a {
		if commandz.Commands == firstCommand {
			return commandz, nil
		}
	}
	var commands datamodel.GentooCommands
	return commands, err
}

// Implement a most simple get portage path
func GetGentooPortagePath() string {

	openJsonFile, err := os.Open("src/imports/GentooPortage.json")
	if err != nil {
		fmt.Printf("Failed to open file, dying.\n")
		panic(err)
	}

	decoder := json.NewDecoder(openJsonFile)
	gentooPortage := &datamodel.GentooCommands{}
	decoder.Decode(&gentooPortage)

	return gentooPortage.Commands
}


// InstallDemoProgram function implements reading from a standard file.json
// and decoding the json for providing the exact style and model
// Alternative method to the bellow suggested "SpecificInstallation"
// I think this soon will overthrow the other
func InstallDemoProgram(w http.ResponseWriter, r *http.Request) {
	ArgumentSecond := "-v"
	ArgumentZero := "argent-skel"

	cmd := exec.Command(GetGentooPortagePath(),  ArgumentZero, ArgumentSecond)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		WriteJson(w, "Installation failed: : \n", ArgumentZero)
	} else {
		WriteJson(w, "Successfully installed the program: \n", ArgumentZero)
		fmt.Print(string(cmdOutput.Bytes()))
	}

}

// Installs a specific program you name while using POST method on
// http://<IP>:8003/installation/install_pkg with a body json-formatted
// { "emerge" : "<any-program-name>"
// The method does NOT search through available Argent packages
// Until 'search' method is implemented, take your chances with this function
// without any fuss.
// Does not take a list [as specific in the name of the function]
// Does not understand [SPACE] between json elements
func SpecificSinglePkgInstallation(w http.ResponseWriter, r *http.Request) {
	var specificPkg datamodel.PackagesInputs

	bytesPkg, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytesPkg, &specificPkg)
	if err != nil {
		fmt.Fprintf(w, "Failed to unmarshal: %s\n", err)
	} else {
		cmd := exec.Command(GetGentooPortagePath(), string(specificPkg.Packages), "-v")
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err := cmd.Run()
		if err != nil {
			os.Stderr.WriteString(err.Error())
			WriteJson(w, "Package not found or body corrupted. Try without [space] between json elements.", string(specificPkg.Packages))
		} else {
			fmt.Print(string(cmdOutput.Bytes()))
			WriteJson(w, "Program installed successfully:", string(specificPkg.Packages))
		}
	}
}

// to be used and refactored in the future
func (a *GCommands) CommandListUpdate(updatedPkg datamodel.GentooCommands) (datamodel.GentooCommands, error) {
	var err error = fmt.Errorf("Could not find commands \n")
	var newPkg GCommands
	for _, packages := range *a {
		if packages.Commands != updatedPkg.Commands {
				newPkg = append(newPkg, updatedPkg)
				err = nil
		} else {
			newPkg = append(newPkg, packages)
		}
	}
	if err == nil {
		*a = newPkg
	}

	return updatedPkg, err
}