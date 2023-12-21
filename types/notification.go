package types

type NotificationToken struct {
	ID 		  string    			`json:"id"        bson:"_id,omitempty"`
	UserID    string             	`json:"userID"    bson:"user_id"`
	DeviceID  string    			`json:"deviceID"  bson:"device_id"`
}

type Notification struct {
	Message 	string `json:"message"`
	Title   	string `json:"title"`
	ToScreen 	string `json:"toScreen"`
}

type CreateNotificationTokenParams struct {
	UserID 		string	`json:"userID"`
	DeviceID 	string  `json:"deviceID"`
}

func NewNotificationToken(params *CreateNotificationTokenParams) (*NotificationToken, error) {
	return &NotificationToken{
		UserID: params.UserID,
		DeviceID: params.DeviceID,
	}, nil
}

type SendNotificationData struct {
	ID  			string
	DeviceID 		string
	UserID          string
	Data 			Notification
}