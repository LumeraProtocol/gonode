package lumera

import (
	"context"
	"fmt"
	"time"

	pb "github.com/pastelnetwork/gonode/gen/lumera"
)

type ActionState int32

const (
	UNKNOWN   ActionState = 0
	PENDING   ActionState = 1
	COMPLETED ActionState = 2
	FAILED    ActionState = 3
)

type ActionResponse struct {
	BlockHeight    int64
	Creator        string
	ActionID       string
	Price          float64
	ExpirationTime int64
	State          ActionState
	Metadata       Metadata
}

type Cascade struct {
	FileName string
	RqIC     int32
	RqMax    int32
}

type Sense struct {
	DDAndFingerprintsIC  int32
	DDAndFingerprintsMax int32
	CollectionID         int32
	GroupID              int32
}

type Metadata struct {
	DataHash string
	Cascade  Cascade
	Sense    Sense
}

func (c *Client) GetAction(ctx context.Context, actionID string) (ActionResponse, error) {
	req := &pb.GetActionRequest{ActionId: actionID}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.actionClient.GetAction(ctx, req)
	if err != nil {
		return ActionResponse{}, fmt.Errorf("failed to fetch lumera: %w", err)
	}

	return toActionResponse(resp), nil
}

func toActionResponse(resp *pb.GetActionResponse) ActionResponse {
	return ActionResponse{
		BlockHeight:    resp.BlockHeight,
		Creator:        resp.Creator,
		ActionID:       resp.ActionId,
		Price:          resp.Price,
		ExpirationTime: resp.ExpirationTime,
		State:          ActionState(resp.State),
		Metadata: Metadata{
			DataHash: resp.Metadata.Datahash,
			Cascade: Cascade{
				FileName: resp.Metadata.Cascade.FileName,
				RqIC:     resp.Metadata.Cascade.RqIc,
				RqMax:    resp.Metadata.Cascade.RqMax,
			},
			Sense: Sense{
				DDAndFingerprintsIC:  resp.Metadata.Sense.DdAndFingerprintsIc,
				DDAndFingerprintsMax: resp.Metadata.Sense.DdAndFingerprintsMax,
				CollectionID:         resp.Metadata.Sense.CollectionId,
				GroupID:              resp.Metadata.Sense.GroupId,
			},
		},
	}
}
