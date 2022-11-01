package client

import (
	"context"
	"github.com/pkg/errors"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/containerengine"
)

type ContainerEngineInterface interface {
	GetVirtualNode(ctx context.Context, virtualNodeId, virtualNodePoolId string) (*containerengine.VirtualNode, error)
}

func (c *client) GetVirtualNode(ctx context.Context, virtualNodeId, virtualNodePoolId string) (*containerengine.VirtualNode, error) {
	if !c.rateLimiter.Reader.TryAccept() {
		return nil, RateLimitError(false, "GetVirtualNode")
	}

	resp, err := c.containerEngine.GetVirtualNode(ctx, containerengine.GetVirtualNodeRequest{
		VirtualNodeId:     common.String(virtualNodeId),
		VirtualNodePoolId: common.String(virtualNodePoolId),
		RequestMetadata:   c.requestMetadata,
	})
	incRequestCounter(err, getVerb, virtualNodeResource)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &resp.VirtualNode, nil
}

// IsVirtualNodeInTerminalState returns true if the virtual node is in a terminal state, false otherwise.
func IsVirtualNodeInTerminalState(virtualNode *containerengine.VirtualNode) bool {
	return virtualNode.LifecycleState == containerengine.VirtualNodeLifecycleStateDeleted ||
		virtualNode.LifecycleState == containerengine.VirtualNodeLifecycleStateFailed
}
