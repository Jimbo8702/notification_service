package client

import (
	"context"

	"github.com/Jimbo8702/notification_service/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	client   types.NotificationServiceClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := types.NewNotificationServiceClient(conn)
	return &GRPCClient{
		Endpoint: endpoint,
		client: c,
	}, nil
}

func (c *GRPCClient) SendUserNotification(ctx context.Context, req *types.SendNotificationRequest) error {
	_, err := c.client.SendUserNotification(ctx, req)
	return err
}