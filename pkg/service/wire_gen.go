// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package service

import (
	"github.com/livekit/livekit-server/pkg/auth"
	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/node"
)

// Injectors from wire.go:

func InitializeServer(conf *config.Config, keyProvider auth.KeyProvider) (*LivekitServer, error) {
	nodeNode, err := node.NewLocalNode(conf)
	if err != nil {
		return nil, err
	}
	roomManager, err := newRoomManagerWithNode(conf, nodeNode)
	if err != nil {
		return nil, err
	}
	roomService, err := NewRoomService(conf, roomManager, nodeNode)
	if err != nil {
		return nil, err
	}
	rtcService := NewRTCService(conf, roomManager)
	livekitServer, err := NewLivekitServer(conf, roomService, rtcService, keyProvider)
	if err != nil {
		return nil, err
	}
	return livekitServer, nil
}