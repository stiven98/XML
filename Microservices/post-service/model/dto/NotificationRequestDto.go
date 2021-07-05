package dto

type NotificationRequestDto struct {
	USERID string `json:"userId"`
	NOTIFYUSERID string `json:"notify_user_id"`
	TYPEOFNOTIFY string `json:"type_of_notify"`
	NOTIFYID string `json:"notify_id"`
}
