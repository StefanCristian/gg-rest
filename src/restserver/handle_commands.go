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

// InstallDemoProgram function implements reading from a standard file.json
// and decoding the json for providing the exact style and model
func InstallDemoProgram(w http.ResponseWriter, r *http.Request) {
	ArgumentSecond := "-v"
	ArgumentZero := "argent-skel"

	openJsonFile, err := os.Open("src/imports/godemoexec.json")
	if err != nil {
		fmt.Printf("Failed to open file, dying.\n")
		panic(err)
	}

	decoder := json.NewDecoder(openJsonFile)
	pkgconfig := &datamodel.GentooCommands{}
	decoder.Decode(&pkgconfig)

	cmd := exec.Command(pkgconfig.Commands,  ArgumentZero, ArgumentSecond)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err = cmd.Run()
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
func SpecificInstallation(w http.ResponseWriter, r *http.Request) {
	var specificPkg datamodel.GentooCommands
	bytesPkg, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytesPkg, &specificPkg)
	if err != nil {
		fmt.Fprintf(w, "Failed to unmarshal: %s\n", err)
	} else {
		cmd := exec.Command("/usr/lib/python-exec/python2.7/emerge", string(specificPkg.Commands), "-v")
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err := cmd.Run()
		if err != nil {
			os.Stderr.WriteString(err.Error())
			WriteJson(w, "Installation failed, internal server error.", string(specificPkg.Commands))
		} else {
			WriteJson(w, "Package installed successfully!", string(specificPkg.Commands) )
			fmt.Print(string(cmdOutput.Bytes()))
		}
		WriteJson(w, "Program installed successfully", string(specificPkg.Commands))
	}
}

// to be used and refactored in the future
func (a *GCommands) specificListAppend(updatedPkg datamodel.GentooCommands) (datamodel.GentooCommands, error) {
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