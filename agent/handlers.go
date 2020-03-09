package main

import (
	"context"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/system"
)

func (s *server) GetSystem(ctx context.Context, in *services.GetSystemRequest) (*resources.System, error) {
	return system.GetSystem()
}
