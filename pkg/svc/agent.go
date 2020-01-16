package svc

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"
//go:generate sh -c "rm -rf k3s statik; mkdir -p k3s/dist/artifacts; curl -Lo k3s/dist/artifacts/k3s https://github.com/rancher/k3s/releases/latest/download/k3s; statik -src k3s/dist/artifacts"
// In case you want to compile `k3s` yourself, uncomment the line below
////go:generate sh -c "rm -rf k3s statik; git clone --depth 1 https://github.com/rancher/k3s.git; cd k3s; make; mkdir -p dist; cd ../; statik -src k3s/dist/artifacts"

import (
	"context"
	k3sd "github.com/pojntfx/k3sd/pkg/proto/generated"
	"github.com/pojntfx/k3sd/pkg/workers"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// K3SAgentManager manages k3s agents.
type K3SAgentManager struct {
	ExtractService
	k3sd.UnimplementedK3SAgentManagerServer
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

// Stop stops the k3s agent.
func (a *K3SAgentManager) Stop(_ context.Context, _ *k3sd.K3SAgentStopArgs) (*k3sd.K3SAgentState, error) {
	if a.K3SManaged == nil {
		msg := "k3s agent hasn't been started yet"

		log.Error(msg)

		return nil, status.Errorf(codes.NotFound, msg)
	}

	log.Info("Stopping k3s agent")

	if err := a.K3SManaged.DisableAutoRestart(); err != nil { // Manually disable auto restart; disables crash recovery even if process is not running
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if a.K3SManaged.IsRunning() {
		if err := a.K3SManaged.Stop(); err != nil { // Stop is sync, so no need to `.Wait()`
			log.Error(err.Error())

			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	a.K3SManaged = nil

	return &k3sd.K3SAgentState{
		Running: false,
	}, nil
}
