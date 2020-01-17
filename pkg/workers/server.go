package workers

import (
	"github.com/pojntfx/go-isc-dhcp/pkg/utils"
	"os"
	"os/exec"
	"syscall"
)

// K3SServer is the k3s server.
type K3SServer struct {
	utils.ProcessWorker
	BinaryDir     string
	NetworkDevice string
	TLSSan        string
}

// Start starts the the k3s server.
func (a *K3SServer) Start() error {
	a.ScheduledForDeletion = false

	command := exec.Command(
		a.BinaryDir,
		"server",
		"--flannel-iface",
		a.NetworkDevice,
		"--tls-san",
		a.TLSSan)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	a.Instance = command

	return command.Start()
}
