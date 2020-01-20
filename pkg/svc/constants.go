package svc

import "path/filepath"

var (
	DataPath       = filepath.Join("/var", "lib", "rancher", "k3s") // DataPath is the path to where the k3s data is being stored.
	ConfigPath     = filepath.Join("/etc", "rancher", "k3s")        // ConfigPath is the path in which the k3s configuration resides.
	KubeconfigPath = filepath.Join(ConfigPath, "k3s.yaml")          // KubeconfigPath is the path to the kubeconfig file.
	TokenPath      = filepath.Join(DataPath, "server", "token")     // TokenPath is the path to the k3s token.
)
