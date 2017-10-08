package datamodel

type GentooCommands struct {
	Commands string `json:"emerge"`
}

// arguments struct are definitely going to change
type GentooArguments struct {
	Install string `json:"install"`
	Remove  string `json:"remove"`
}

var Commands = []GentooCommands{}

var Arguments = []GentooArguments{}