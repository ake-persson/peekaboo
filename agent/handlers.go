package main

import (
	"context"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/system"
	"github.com/peekaboo-labs/peekaboo/pkg/user"
)

func (s *server) GetSystem(ctx context.Context, in *services.GetSystemRequest) (*resources.System, error) {
	return system.GetSystem()
}

func (s *server) ListUsers(ctx context.Context, in *services.ListUsersRequest) (*services.ListUsersResponse, error) {
	return user.ListUsers()
}
