package saga

import "encoding/json"


const (
	ProfileChannel    string = "ProfileChannel"
	FollowerChannel    string = "FollowerChannel"
	ReplyChannel    string = "ReplyChannel"
	ServiceProfile    string = "Profile"
	ServiceFollower    string = "Follower"
	ActionStart     string = "Start"
	ActionDone      string = "DoneMsg"
	ActionError     string = "ErrorMsg"
	ActionRollback  string = "RollbackMsg"
)

type Message struct {
	Service       string         `json:"service"`
	SenderService string         `json:"sender_service"`
	Action        string         `json:"action"`
	UserId        string         `json:"user_id"`
	SaveUserId    string  	 	 `json:"save_user_id"`
	Ok            bool           `json:"ok"`
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
