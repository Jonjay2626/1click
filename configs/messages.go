package configs

// All the strings that are needed for debugging and info logging, and constant strings.
const (
	CheckingDependencies            = "Checking dependencies: %s"
	DependenciesPending             = "pending dependencies: %s"
	DependenciesOK                  = "All dependencies are installed on host machine"
	GeneratingDockerComposeScript   = "Generating docker-compose script for current selection of clients"
	GeneratingEnvFile               = "Generating environment file for current selection of clients"
	InstallNotSupported             = "Install support for %s is not available for %s. Please install it and try again"
	Exiting                         = "Exiting..."
	InstructionsFor                 = "Instructions for %s"
	OSNotSupported                  = "installation not supported for %s"
	ProvideClients                  = "Please provide both execution client and consensus client"
	CreatedFile                     = "Created file %s"
	DefaultDockerComposeScriptsPath = "./docker-compose-scripts"
	OnPremiseExecutionURL           = "http://localhost:8545"
	OnPremiseConsensusURL           = "http://localhost:4000"
	ClientNotSupported              = "client %s is not supported. Please use 'clients' command to see the list of supported clients"
	PrintingFile                    = "File %s :"
	SupportedClients                = "Supported clients of type %s: %v"
	ConfigClients                   = "Provided clients of type %s in configuration file: %v"
	RunningDockerCompose            = "Running docker-compose script"
	Component                       = "component"
	RunningCommand                  = "Running command: %s"
	ConfigFileName                  = ".1Click"
	UnableToProceed                 = "Unable to proceed. Please check the logs for more details"
	DefaultDockerComposeScriptName  = "docker-compose.yml"
	CheckingDockerEngine            = "Checking if docker engine is on"
	DepositCLIDockerImageName       = "deposit-cli:local"
	GeneratingKeystore              = "Generating keystore folder"
	KeysFoundAt                     = "If everything went well, your keys can be found at: %s"
)
