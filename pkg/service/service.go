package service

import (
	"fmt"

	"github.com/google/wire"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/node"
	"github.com/livekit/livekit-server/proto"
)

var ServiceSet = wire.NewSet(
	NewRoomService,
	NewRTCService,
)

func NewRoomService(conf *config.Config, localNode *node.Node) (proto.RoomService, error) {
	if conf.MultiNode {
		return nil, fmt.Errorf("multinode is not supported")
	} else {
		return NewSimpleRoomService(localNode)
	}
}