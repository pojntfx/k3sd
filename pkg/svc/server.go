package svc

import (
	"context"
	k3sd "github.com/pojntfx/k3sd/pkg/proto/generated"
	"github.com/pojntfx/k3sd/pkg/utils"
	"github.com/pojntfx/k3sd/pkg/workers"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path/filepath"
)

// K3SServerManager manages k3s servers.
type K3SServerManager struct {
	ExtractService
	k3sd.UnimplementedK3SServerManagerServer
	K3SManaged *workers.K3SServer
}

// Start starts the k3s server.
func (s *K3SServerManager) Start(_ context.Context, args *k3sd.K3SServer) (*k3sd.K3SServerState, error) {
	if s.K3SManaged != nil && s.K3SManaged.Instance != nil {
		msg := "k3s server is already running."

		log.Error(msg)

		return nil, status.Errorf(codes.AlreadyExists, msg)
	}

	dirCleanupWorker := utils.DirCleanupWorker{
		DirsToClean: []string{
			filepath.Join("/var", "lib", "rancher", "k3s"),
			filepath.Join("/etc", "rancher", "k3s")},
	}

	k3s := workers.K3SServer{
		DirCleanupWorker: dirCleanupWorker,
		BinaryDir:        s.BinaryDir,
		KubeconfigPath:   filepath.Join("/etc", "rancher", "k3s", "k3s.yaml"),
		NetworkDevice:    args.GetNetworkDevice(),
		TLSSan:           args.GetTLSSan(),
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

	s.K3SManaged = &k3s

	return &k3sd.K3SServerState{
		Running: true,
	}, nil
}

// Stop stops the k3s server.
func (s *K3SServerManager) Stop(_ context.Context, _ *k3sd.K3SServerEmptyArgs) (*k3sd.K3SServerState, error) {
	if s.K3SManaged != nil && s.K3SManaged.Instance != nil {
		log.Info("Stopping k3s server")

		if err := s.K3SManaged.DisableAutoRestart(); err != nil { // Manually disable auto restart; disables crash recovery even if process is not running
			log.Error(err.Error())

			return nil, status.Errorf(codes.Unknown, err.Error())
		}

		if s.K3SManaged.IsRunning() {
			if err := s.K3SManaged.Stop(); err != nil { // Stop is sync, so no need to `.Wait()`
				log.Error(err.Error())

				return nil, status.Errorf(codes.Unknown, err.Error())
			}
		}

		s.K3SManaged.Instance = nil

		return &k3sd.K3SServerState{
			Running: false,
		}, nil
	}

	msg := "k3s server hasn't been started yet"

	log.Error(msg)

	return nil, status.Errorf(codes.NotFound, msg)
}

// Cleanup cleans the state of the k3s server.
func (s *K3SServerManager) Cleanup(_ context.Context, _ *k3sd.K3SServerEmptyArgs) (*k3sd.K3SServerDeletionState, error) {
	if s.K3SManaged != nil && s.K3SManaged.Instance != nil {
		msg := "k3s server is running, can't clean it's state."

		log.Error(msg)

		return nil, status.Errorf(codes.Unknown, msg)
	}

	log.Info("Cleaning k3s server state")

	if err := s.K3SManaged.CleanupDirs(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	return &k3sd.K3SServerDeletionState{
		Deleted: true,
	}, nil
}

// GetKubeconfig returns the Kubeconfig for the k3s server.
func (s *K3SServerManager) GetKubeconfig(_ context.Context, _ *k3sd.K3SServerEmptyArgs) (*k3sd.K3SServerKubeconfig, error) {
	kubeconfig, err := s.K3SManaged.GetKubeconfig()
	if err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &k3sd.K3SServerKubeconfig{
		Content: kubeconfig,
	}, nil
}
