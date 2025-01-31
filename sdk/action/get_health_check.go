package action

import (
	"context"
	"fmt"
	"time"

	pb "github.com/pastelnetwork/gonode/gen/action"
)

type GetHealthCheckResponse struct {
	Status string
}

func (c *Client) CheckHealth(ctx context.Context) (GetHealthCheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.client.CheckHealth(ctx, &pb.GetHealthCheckRequest{})
	if err != nil {
		return GetHealthCheckResponse{}, fmt.Errorf("failed to fetch lumera: %w", err)
	}

	return GetHealthCheckResponse{
		Status: resp.Status,
	}, nil
}
