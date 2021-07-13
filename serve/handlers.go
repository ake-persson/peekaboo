package serve

import (
	"context"

	"github.com/ake-persson/peekaboo/pkg/filesystem"
	"github.com/ake-persson/peekaboo/pkg/group"
	"github.com/ake-persson/peekaboo/pkg/pb/v1/resources"
	"github.com/ake-persson/peekaboo/pkg/pb/v1/services"
	"github.com/ake-persson/peekaboo/pkg/system"
	"github.com/ake-persson/peekaboo/pkg/user"
)

func (s *server) GetSystem(ctx context.Context, in *services.GetSystemRequest) (*resources.System, error) {
	return system.GetSystem()
}

func (s *server) ListUsers(ctx context.Context, in *services.ListUsersRequest) (*services.ListUsersResponse, error) {
	return user.ListUsers()
}

func (s *server) ListGroups(ctx context.Context, in *services.ListGroupsRequest) (*services.ListGroupsResponse, error) {
	return group.ListGroups()
}

func (s *server) ListFilesystems(ctx context.Context, in *services.ListFilesystemsRequest) (*services.ListFilesystemsResponse, error) {
	return filesystem.ListFilesystems()
}
