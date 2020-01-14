package svc

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"
//go:generate sh -c "rm -rf k3s statik; mkdir -p k3s/dist/artifacts; curl -Lo k3s/dist/artifacts/k3s https://github.com/rancher/k3s/releases/latest/download/k3s; statik -src k3s/dist/artifacts"
// In case you want to compile `k3s` yourself, uncomment the line below
////go:generate sh -c "rm -rf k3s statik; git clone --depth 1 https://github.com/rancher/k3s.git; cd k3s; make; mkdir -p dist; cd ../; statik -src k3s/dist/artifacts"

import (
	"context"
	k3sd "github.com/pojntfx/k3sd/pkg/proto/generated"
	_ "github.com/pojntfx/k3sd/pkg/svc/statik" // Embedded k3s binary
	"github.com/pojntfx/k3sd/pkg/workers"
	"github.com/rakyll/statik/fs"
	"gitlab.com/bloom42/libs/rz-go/log"
	"os"
)

// K3SAgentManager manages k3s agents.
type K3SAgentManager struct {
	k3sd.UnimplementedK3SAgentManagerServer
	BinaryDir  string
	K3SManaged *workers.K3SAgent
}

// Start starts the k3s agent.
func (a *K3SAgentManager) Start(_ context.Context, args *k3sd.K3SAgent) (*k3sd.K3SAgentState, error) {
	k3s := workers.K3SAgent{
		BinaryDir:     a.BinaryDir,
		NetworkDevice: args.GetNetworkDevice(),
		Token:         args.GetToken(),
		ServerURL:     args.GetServerURL(),
	}

	if err := k3s.Start(); err != nil {
		log.Error(err.Error())
	}

	go func(k3s *workers.K3SAgent) {
		log.Info("Starting k3s agent")

		_ = k3s.Wait()

		// Keep the k3s agent running
		for {
			if !k3s.IsScheduledForDeletion() {
				log.Info("Restarting k3s agent")

				if err := k3s.Start(); err != nil {
					log.Error(err.Error())
				}

				_ = k3s.Wait()
			} else {
				break
			}
		}
	}(&k3s)

	a.K3SManaged = &k3s

	return &k3sd.K3SAgentState{
		Running: true,
	}, nil
}

// Extract extracts the k3s binary.
func (a *K3SAgentManager) Extract() error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	data, err := fs.ReadFile(statikFS, "/k3s")
	if err != nil {
		return err
	}

	binaryFile, err := os.Create(a.BinaryDir)
	if err != nil {
		return err
	}

	if _, err = binaryFile.Write(data); err != nil {
		return err
	}
	defer binaryFile.Close()

	if err := os.Chmod(a.BinaryDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Cleanup deletes the extracted k3s binary.
func (a *K3SAgentManager) Cleanup() error {
	return os.Remove(a.BinaryDir)
}
