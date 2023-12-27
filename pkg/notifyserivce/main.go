package notifyserivce

import (
	"github.com/Jimbo8702/notification_service/client"
	"github.com/Jimbo8702/notification_service/types"
)

type SendNotificationRequest types.SendNotificationRequest

// client interface for the notification service
type Client client.Client

// grpc client for notification service
type GRPCClient client.GRPCClient

// makes a new grpc notification service
func NewGRPCClient(endpoint string) (*client.GRPCClient, error) {
	return client.NewGRPCClient(endpoint)
}
