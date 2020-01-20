package workers

import (
	"errors"
	dhcpUtils "github.com/pojntfx/go-isc-dhcp/pkg/utils"
	"github.com/pojntfx/k3sd/pkg/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

// K3SServer is the k3s server.
type K3SServer struct {
	utils.DirCleanupWorker
	dhcpUtils.ProcessWorker
	BinaryDir      string
	NetworkDevice  string
	TLSSan         string
	KubeconfigPath string
	TokenPath      string
}

// Start starts the the k3s server.
func (s *K3SServer) Start() error {
	s.ScheduledForDeletion = false

	command := exec.Command(
		s.BinaryDir,
		"server",
		"--flannel-iface",
		s.NetworkDevice,
		"--tls-san",
		s.TLSSan)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	s.Instance = command

	return command.Start()
}

// GetKubeconfig returns the Kubeconfig for the server.
func (s *K3SServer) GetKubeconfig() (string, error) {
	if _, err := os.Stat(s.KubeconfigPath); os.IsNotExist(err) {
		return "", errors.New("could not find Kubeconfig")
	}

	data, err := ioutil.ReadFile(s.KubeconfigPath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// GetToken returns the token for the server.
func (s *K3SServer) GetToken() (string, error) {
	if _, err := os.Stat(s.KubeconfigPath); os.IsNotExist(err) {
		return "", errors.New("could not find token")
	}

	data, err := ioutil.ReadFile(s.TokenPath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
