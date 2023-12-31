package main

import (
	"context"
	"errors"

	"firebase.google.com/go/v4/messaging"
	"github.com/Jimbo8702/notification_service/types"
)

type NotificationService interface {
	SendUserNotification(ctx context.Context, userID string, data *types.Notification) error
	SendNotificationCampaign(ctx context.Context, users []string, data *types.Notification) error
	//add in send email
}

type FCMNotificationService struct {
	fBClient *messaging.Client
	store    NotificationTokenStore
}

func NewFCMNotificationService(client *messaging.Client, store NotificationTokenStore) NotificationService {
	return &FCMNotificationService{
		fBClient: client,
		store: store,
	}
}

func (s *FCMNotificationService) SendUserNotification(ctx context.Context, userID string, data *types.Notification) error {
	tokens, err := s.store.GetTokensByProfileID(ctx, userID)
	if err != nil {
		return err
	}
	if len(tokens) > 0 {
		_, err = s.fBClient.Send(ctx, &messaging.Message{
			Notification: &messaging.Notification{
				Title: data.Title,
				Body: data.Message,
			},
			Token: tokens[0].DeviceID,
			Data: map[string]string{
				"screen": data.ToScreen,
			},
		})
		if err != nil {
			return err
		}
	} else {
		return errors.New("no tokens found for user")
	}
	return nil
}

func (s *FCMNotificationService) SendNotificationCampaign(ctx context.Context, users []string, data *types.Notification) error {
	return nil
}