package workers

import (
	dhcpUtils "github.com/pojntfx/go-isc-dhcp/pkg/utils"
	"github.com/pojntfx/k3sd/pkg/utils"
	"os"
	"os/exec"
	"syscall"
)

// K3SAgent is the k3s agent.
type K3SAgent struct {
	dhcpUtils.ProcessWorker
	utils.DirCleanupWorker
	BinaryDir     string
	NetworkDevice string
	Token         string
	ServerURL     string
}

// Start starts the the k3s agent.
func (a *K3SAgent) Start() error {
	a.ScheduledForDeletion = false

	command := exec.Command(
		a.BinaryDir,
		"agent",
		"--flannel-iface",
		a.NetworkDevice,
		"--token",
		a.Token,
		"--server",
		a.ServerURL)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	a.Instance = command

	return command.Start()
}
