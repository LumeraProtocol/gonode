package lumera

import (
	"context"
	"fmt"
	pb "github.com/pastelnetwork/gonode/gen/lumera"
	"time"
)

type Supernode struct {
	NodeID string
	Height int64
}

type Supernodes []Supernode

type GetTopSupernodesForBlockResponse struct {
	Supernodes Supernodes
}

func (c *Client) GetTopSNsByBlockHeight(ctx context.Context, blockHeight int64) (GetTopSupernodesForBlockResponse, error) {
	req := &pb.GetTopSupernodesForBlockRequest{BlockHeight: blockHeight}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.supernodeClient.GetTopSupernodesForBlock(ctx, req)
	if err != nil {
		return GetTopSupernodesForBlockResponse{}, fmt.Errorf("failed to fetch lumera: %w", err)
	}

	return toGetTopSNsForBlockResponse(resp), nil
}

func toGetTopSNsForBlockResponse(response *pb.GetTopSupernodesForBlockResponse) GetTopSupernodesForBlockResponse {
	var sns Supernodes

	for _, sn := range response.Supernodes {
		sns = append(sns, Supernode{
			NodeID: sn.NodeId,
			Height: sn.BlockHeight,
		})
	}

	return GetTopSupernodesForBlockResponse{
		Supernodes: sns,
	}
}
