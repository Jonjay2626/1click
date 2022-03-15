package utils

import "text/template"

// DistroInfo : Struct Contains name, architecture and version of the linux distribution of the host machine
type DistroInfo struct {
	Name         string
	Architecture string
	Version      string
}

// Script : Struct Represents a script to be executed
type Script struct {
	// Tmp : script template
	Tmp *template.Template
	// getOutput : True to get output of the script
	GetOutput bool
	// Data: template data object
	Data interface{}
}
