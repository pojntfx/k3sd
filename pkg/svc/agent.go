package svc

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	k3sd "github.com/pojntfx/k3sd/pkg/proto/generated"
)

// K3SAgentManager manages k3s agents.
type K3SAgentManager struct {
	k3sd.UnimplementedK3SAgentManagerServer
}
