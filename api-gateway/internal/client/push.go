package client

import proto "api-gateway/internal/proto/push"

type PushClient struct {
	Client proto.NotificationServiceClient
}

func NewPushClient
