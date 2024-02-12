package grpc

import (
	"context"

	umeeOracleTypes "bharvest.io/oraclemon/client/grpc/protobuf/umee-oracle"
	"google.golang.org/grpc"
)

type Client interface {
	Connect(ctx context.Context) error
	Terminate(_ context.Context) error
}

type (
	Umee struct {
		host string
		conn *grpc.ClientConn
		oracleClient umeeOracleTypes.QueryClient
	}
)

type (
	UmeeOracleParams struct {
		AcceptList map[string]bool
		SlashWindow uint64
		MinUptime float64
	}
)
