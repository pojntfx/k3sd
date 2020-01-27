package cmd

const (
	keyPrefix         = "k3sagent."
	configFileDefault = ""
	serverHostPortKey = keyPrefix + "serverHostPort"
	configFileKey     = keyPrefix + "configFile"
	networkDeviceKey  = keyPrefix + "networkDevice"
	tokenKey          = keyPrefix + "token"
	serverUrlKey      = keyPrefix + "serverUrl"
)

var (
	serverHostPortFlag string
	configFileFlag     string
)
