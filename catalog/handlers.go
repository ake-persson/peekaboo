package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

func (s *server) RegisterSystem(ctx context.Context, in *resources.System) (*empty.Empty, error) {
	now := ptypes.TimestampNow()
	if _, ok := s.systems[in.Hostname]; ok {
		in.Updated = now
	} else {
		in.Created = now
	}
	in.LastSeen = now

	s.systems[in.Hostname] = in
	return &empty.Empty{}, nil
}

func (s *server) SystemKeepAlive(ctx context.Context, in *services.SystemKeepAliveRequest) (*empty.Empty, error) {
	now := ptypes.TimestampNow()
	s.logger.Debug("keep alive",
		zap.String("hostname", in.Hostname),
		zap.String("timestamp", now.String()))
	s.systems[in.Hostname].LastSeen = now
	return &empty.Empty{}, nil
}

func (s *server) ListSystems(req *services.ListSystemsRequest, stream services.CatalogService_ListSystemsServer) error {
	for _, v := range s.systems {
		stream.Send(v)
	}
	return nil
}
