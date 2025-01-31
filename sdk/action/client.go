package action

import (
	"context"
	"fmt"

	pb "github.com/pastelnetwork/gonode/gen/action"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.ActionServiceClient
}

type ActionSDK interface {
	CheckHealth(ctx context.Context) (GetHealthCheckResponse, error)
}

func NewClient(serverAddr string) (*Client, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return &Client{
		conn:   conn,
		client: pb.NewActionServiceClient(conn),
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
