package restserver

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os/exec"
	"bytes"
	"os"
)

// the WriteJson function will be used by the server to output
// useful messages to the REST client using GET, PUT, POST, etc.
// it parses events through json.Marshal-ing a data interface
// received
func WriteJson(w http.ResponseWriter, data interface{}, s string) {
	var jsonBytes, err = json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Failed to marshal data: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonBytes), string(s))
}

func Index(w http.ResponseWriter, r *http.Request) {
	initialDescription := struct {
		Message        string `json:"message"`
	}{
		Message:        "General Gentoo REST Api Service",
	}
	WriteJson(w, initialDescription, "")
}

func ParseJson(jsonByteArray []byte, jsonStruct interface{}) (error) {
	err := json.Unmarshal(jsonByteArray, &jsonStruct);
	if err != nil {
		fmt.Println("An error ocurred: ", err)
	}

	return err
}

// command log output function is not used (yet)
// it will be used as long as there will be a better method
func OutputCommandLog(a *exec.Cmd) {
	cmdOutput := &bytes.Buffer{}
	a.Stdout = cmdOutput
	err := a.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		fmt.Println("command failed\n ", err)
	} else {
		fmt.Println("command ran successfully\n", a)
		fmt.Print(string(cmdOutput.Bytes()))
	}
}