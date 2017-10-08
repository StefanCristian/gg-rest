package restserver

import (
	"datamodel"
	"imports"
	"net/http"
)

type GArguments []datamodel.GentooArguments

var arguments GArguments = imports.ImportArguments()

func GetAllArguments(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, arguments, "")
}