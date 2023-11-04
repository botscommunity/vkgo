package update

import "encoding/json"

type Updates struct {
	TS      any      `json:"ts,omitempty"`
	Updates []Update `json:"updates,omitempty"`
	Failed  int      `json:"failed,omitempty"`
}

type Update struct {
	GroupID int             `json:"group_id,omitempty"`
	Type    string          `json:"type,omitempty"`
	EventID string          `json:"event_id,omitempty"`
	Version string          `json:"v,omitempty"`
	Object  json.RawMessage `json:"object,omitempty"`
}
