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

type SendNotificationData struct {
	ID  			string
	DeviceID 		string
	UserID          string
	Data 			Notification
}