package lumera

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/pastelnetwork/gonode/gen/lumera"
)

type Client struct {
	conn            *grpc.ClientConn
	actionClient    pb.ActionServiceClient
	supernodeClient pb.SupernodeServiceClient
}

type LumeraService interface {
	GetAction(ctx context.Context, actionID string) (ActionResponse, error)
	GetTopSNsByBlockHeight(ctx context.Context, blockHeight int64) (GetTopSupernodesForBlockResponse, error)
}

func NewClient(serverAddr string) (LumeraService, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return &Client{
		conn:            conn,
		actionClient:    pb.NewActionServiceClient(conn),
		supernodeClient: pb.NewSupernodeServiceClient(conn),
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
