package datamodel

type GentooCommands struct {
	Commands string `json:"emerge"`
}

type PackagesInputs struct {
	Packages string `json:"packages"`
}

// arguments struct are definitely going to change
type GentooArguments struct {
	Install string `json:"install"`
	Remove  string `json:"remove"`
}