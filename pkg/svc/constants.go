package svc

import "path/filepath"

var (
	DataPath   = filepath.Join("/var", "lib", "rancher", "k3s") // DataPath is the path to where the k3s data is being stored.
	ConfigPath = filepath.Join("/etc", "rancher", "k3s")        // ConfigPath is the path in which the k3s configuration resides.
)
