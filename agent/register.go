package main

import (
	"context"
	"fmt"
	"log"
	"time"

	//	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/system"
)

func registerSystem(conf *config) error {
	opts := []grpc.DialOption{grpc.WithBlock()}
	if conf.NoTLS {
		opts = append(opts, grpc.WithInsecure())
	} else {
		// Load CA certificate file.
		creds, err := credentials.NewClientTLSFromFile(conf.CAFile, "")
		if err != nil {
			return fmt.Errorf("failed to load ca certificate: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// Connect to gRPC server.
	conn, err := grpc.Dial(conf.CatalogAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Initialize new client.
	client := services.NewCatalogServiceClient(conn)

	// Create context for gRPC calls.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Register system.
	sys, err := system.GetSystem()
	if err != nil {
		return err
	}

	if _, err := client.RegisterSystem(ctx, sys); err != nil {
		return fmt.Errorf("register system: %v", err)
	}

	// System keep alive.
	for {
		log.Printf("keep alive")

		// Create context for gRPC calls.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if _, err := client.SystemKeepAlive(ctx, &services.SystemKeepAliveRequest{Hostname: sys.Hostname}); err != nil {
			return fmt.Errorf("keep alive: %v", err)
		}

		time.Sleep(time.Duration(conf.KeepAlive) * time.Second)
	}

	return nil
}
