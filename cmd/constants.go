package cmd

const (
	K3SDHostPortDefault        = "localhost:1340"                       // K3SDPortDefault is the default Host:port of `k3sd`.
	K3SDHostPortDefaultMessage = "Host:port of the k3sd server to use." // K3SDHostPortDefaultMessage is the documentation for the host:port flag.
)

const (
	CouldNotBindFlagsErrorMessage        = "Could not bind flags"         // CouldNotBindFlagsErrorMessage is the error message to throw if binding the flags has failed.
	CouldNotStartRootCommandErrorMessage = "Could not start root command" // CouldNotStartRootCommandErrorMessage is the error message to throw if starting the root command has failed.
)
