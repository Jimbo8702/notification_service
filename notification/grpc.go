package main

import (
	"context"

	"github.com/Jimbo8702/notification_service/types"
)

type GRPCNotificationServer struct {
	types.UnimplementedNotificationServiceServer
	svc NotificationService
}

func NewGRPCNotificationServer(svc NotificationService) *GRPCNotificationServer {
	return &GRPCNotificationServer{
		svc: svc,
	}
}

func (s *GRPCNotificationServer) SendUserNotification(ctx context.Context, req *types.SendNotificationRequest) (*types.None, error) {
	notification := &types.Notification{
		Title: req.Title,
		Message: req.Message,
		ToScreen: req.ToScreen,
	}
	return  &types.None{}, s.svc.SendUserNotification(ctx, req.UserID, notification)
}