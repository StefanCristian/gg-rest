package restserver

import (
	"net/http"
	"os"
	"gitlab.com/stefancristian/mux"
)

const RESTAPI_PORT_NAME = "API_PORT"
const RESTAPI_PORT_VALUE = "8003"

// Starting the REST server function
func StartRestServer() {
	listRouter := mux.NewRouter()
	for _, serverRoutes := range restroutes {
		listRouter.HandleFunc(serverRoutes.Pattern, serverRoutes.HandlerFunc).Methods(serverRoutes.Method)
	}
	var port = getPort()
	// Out of personal observations, this ":" must never be changed
	if err := http.ListenAndServe(":" + port, listRouter); err != nil {
		panic(err)
	}
}

// Getting the port through the os function Getenv
// if the port is null, it returns to the port passing
// os.Getenv will not get the value from anything else or nowhere else
// than declaring its constant at line 10
func getPort() string {
	var port string = os.Getenv(RESTAPI_PORT_NAME)
	if port != "" {
		return port
	} else {
		return RESTAPI_PORT_VALUE
	}
}
