package restserver

import "net/http"

// general usage constants for demo
const (
	installationURL   = "/installation"
	specificInstallation = installationURL + "/install_pkg"
	argumentsURL    = "/arguments"
	argumentsByName   = argumentsURL + "/{argumentname}"
	allcommandsURL = "/allcommands"
)

// declare Route structure containing Methods, Patterns and http Handlers
type Route struct {
	Method string

	Pattern string

	HandlerFunc http.HandlerFunc
}

// declaring Routes as type []Route array
type Routes []Route

// declaring our RESTful routes with their Methods, Patterns and Handlers
// Handlers will be used to do specific actions based on patterns and methods
var restroutes = Routes {
	Route{
		Method:       "GET",
		Pattern:      "/",
		HandlerFunc:  Index,
	},
	Route{
		Method:       "GET",
		Pattern:      allcommandsURL,
		HandlerFunc:  GetAllCommands,
	},
	Route{
		Method:       "POST",
		Pattern:      installationURL,
		HandlerFunc:  InstallDemoProgram,
	},
	Route{
		Method:       "POST",
		Pattern:      specificInstallation,
		HandlerFunc:  SpecificSinglePkgInstallation,
	},
}