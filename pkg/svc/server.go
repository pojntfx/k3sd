package svc

import (
	"context"
	k3sd "github.com/pojntfx/k3sd/pkg/proto/generated"
	"github.com/pojntfx/k3sd/pkg/workers"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// K3SServerManager manages k3s servers.
type K3SServerManager struct {
	ExtractService
	k3sd.UnimplementedK3SServerManagerServer
	K3SManaged *workers.K3SServer
}

// Start starts the k3s server.
func (a *K3SServerManager) Start(_ context.Context, args *k3sd.K3SServer) (*k3sd.K3SServerState, error) {
	k3s := workers.K3SServer{
		BinaryDir:     a.BinaryDir,
		NetworkDevice: args.GetNetworkDevice(),
		TLSSan:        args.GetTLSSan(),
	}

	if err := k3s.Start(); err != nil {
		log.Error(err.Error())
	}

	go func(k3s *workers.K3SServer) {
		log.Info("Starting k3s server")

		_ = k3s.Wait()

		// Keep the k3s server running
		for {
			if !k3s.IsScheduledForDeletion() {
				log.Info("Restarting k3s server")

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

	return &k3sd.K3SServerState{
		Running: true,
	}, nil
}

// Stop stops the k3s server.
func (a *K3SServerManager) Stop(_ context.Context, _ *k3sd.K3SServerStopArgs) (*k3sd.K3SServerState, error) {
	if a.K3SManaged == nil {
		msg := "k3s server hasn't been started yet"

		log.Error(msg)

		return nil, status.Errorf(codes.NotFound, msg)
	}

	log.Info("Stopping k3s server")

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

	return &k3sd.K3SServerState{
		Running: false,
	}, nil
}
