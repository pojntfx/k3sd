package cmd

const (
	keyPrefix         = "k3sserver."
	configFileDefault = ""
	serverHostPortKey = keyPrefix + "serverHostPort"
	configFileKey     = keyPrefix + "configFile"
	networkDeviceKey  = keyPrefix + "networkDevice"
	tlsSanKey         = keyPrefix + "tlsSan"
)

var (
	serverHostPortFlag string
	configFileFlag     string
)
