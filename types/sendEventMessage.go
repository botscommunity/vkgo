package types

import "encoding/json"

type SendEventMessage struct {
	EventID string
	ChatID  int
	UserID  int
	Action  string
}

type EventAction struct {
	Type string `json:"type,omitempty"`

	Text string `json:"text,omitempty"` //  show_snackbar
	Link string `json:"link,omitempty"` // open_link

	// open_app
	AppID         int    `json:"app_id,omitempty"`
	ApplicationID int    `json:"-"`
	OwnerID       int    `json:"owner_id,omitempty"`
	Hash          string `json:"hash,omitempty"`
}

func (action EventAction) JSON() string {
	var data []byte

	if action.Type == "open_app" && action.ApplicationID != 0 {
		action.AppID = action.ApplicationID
	}

	if marshal, err := json.Marshal(action); err == nil {
		data = marshal
	}

	return string(data)
}
