package client

import (
	"context"

	"github.com/Jimbo8702/notification_service/types"
)

type Client interface {
	SendUserNotification(context.Context, *types.SendNotificationRequest) error
}