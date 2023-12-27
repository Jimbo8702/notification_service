package main

import (
	"context"
	"time"

	"github.com/Jimbo8702/notification_service/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next NotificationService
}

func NewLogMiddleware(next NotificationService) NotificationService {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) SendUserNotification(ctx context.Context, userID string, data *types.Notification) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err": err,
			"user_id": userID,
		}).Info("sending user push notification")
	}(time.Now())
	return l.next.SendUserNotification(ctx, userID, data)
}

func (l *LogMiddleware) SendNotificationCampaign(ctx context.Context, users []string, data *types.Notification) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err": err,
			"user_count": len(users),
		}).Info("sending push notification campaign")
	}(time.Now())
	return l.next.SendNotificationCampaign(ctx, users, data)
}

