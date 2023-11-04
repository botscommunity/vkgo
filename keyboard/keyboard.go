package keyboard

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type ReplyKeyboard struct {
	Disposable bool      `json:"one_time"`
	Embedded   bool      `json:"inline"`
	Buttons    []Buttons `json:"buttons"`
}

type Buttons []Button
type Button struct {
	Action Action `json:"action,omitempty"`
	Color  string `json:"color,omitempty"`
}

type Action struct {
	Type  string `json:"type,omitempty"`
	Link  string `json:"link,omitempty"`
	Label string `json:"label,omitempty"`

	OwnerID       int    `json:"owner_id,omitempty"`
	ApplicationID int    `json:"app_id,omitempty"`
	Hash          string `json:"hash,omitempty"`

	Payload any `json:"payload,omitempty"`
}

type Label string
type Payload map[string]string
type Color string

type OwnerID int
type ApplicationID int
type AppID int
type Hash string

const (
	White     = "default"
	Default   = "default"
	Secondary = "default"

	Blue    = "primary"
	Primary = "primary"

	Green    = "positive"
	Positive = "positive"

	Red      = "negative"
	Negative = "negative"
)

func getColor(color string) (string, bool) {
	colors := map[string]string{
		"white": "default",
		"blue":  "primary",
		"green": "positive",
		"red":   "negative",
	}

	vkColor, ok := colors[color]

	return vkColor, ok
}

type Keyboard struct {
	Inline  bool
	OneTime bool
	Buttons any
}

func New(properties ...any) *ReplyKeyboard {
	keyboard := &ReplyKeyboard{
		Buttons: []Buttons{},
	}

	for _, property := range properties {
		if property, ok := property.(Keyboard); ok {
			keyboard.Disposable = property.OneTime
			keyboard.Embedded = property.Inline

			if property.Buttons == nil {
				keyboard.Buttons = make([]Buttons, 0)
			}

			switch property := property.Buttons.(type) {
			case Buttons:
				keyboard.Buttons = []Buttons{property}
			case []Buttons:
				keyboard.Buttons = property
			}
		}
	}

	return keyboard
}

func (keyboard *ReplyKeyboard) OneTime(property ...bool) *ReplyKeyboard {
	return keyboard.SetOneTime(property...)
}

func (keyboard *ReplyKeyboard) SetOneTime(property ...bool) *ReplyKeyboard {
	if len(property) > 0 {
		keyboard.Disposable = property[0]
		return keyboard
	}

	if keyboard.Embedded {
		keyboard.Embedded = false
	}

	keyboard.Disposable = true

	return keyboard
}

func (keyboard *ReplyKeyboard) Inline(property ...bool) *ReplyKeyboard {
	return keyboard.SetInline(property...)
}

func (keyboard *ReplyKeyboard) SetInline(property ...bool) *ReplyKeyboard {
	if len(property) > 0 {
		keyboard.Disposable = property[0]
		return keyboard
	}

	if keyboard.Disposable {
		keyboard.Disposable = false
	}

	keyboard.Embedded = true

	return keyboard
}

func (keyboard *ReplyKeyboard) Add(button Button) *ReplyKeyboard {
	var (
		count     = len(keyboard.Buttons)
		color, ok = getColor(button.Color)
	)

	if ok {
		button.Color = color
	}

	if count == 0 {
		keyboard.Row()
		keyboard.Buttons[0] = append(keyboard.Buttons[0], button)
	} else {
		keyboard.Buttons[count-1] = append(keyboard.Buttons[count-1], button)
	}

	return keyboard
}

func (keyboard *ReplyKeyboard) Row() *ReplyKeyboard {
	keyboard.Buttons = append(keyboard.Buttons, make(Buttons, 0))
	return keyboard
}

type Text struct {
	Text    string
	Label   string
	Payload any
	Color   any
}

type Callback Text

func (keyboard *ReplyKeyboard) Text(properties ...any) *ReplyKeyboard {
	title, payload, color := parsePrimitiveButtonProperties(properties...)

	keyboard.Add(Button{
		Action: Action{Type: "text", Label: title, Payload: payload},
		Color:  color,
	})

	return keyboard
}

func (keyboard *ReplyKeyboard) Callback(properties ...any) *ReplyKeyboard {
	title, payload, color := parsePrimitiveButtonProperties(properties...)

	keyboard.Add(Button{
		Action: Action{Type: "callback", Label: title, Payload: payload},
		Color:  color,
	})

	return keyboard
}

type Link struct {
	Label   string
	Link    string
	Payload string
}

func (keyboard *ReplyKeyboard) Link(properties ...any) *ReplyKeyboard {
	var (
		label   string
		link    string
		payload string
	)

	for _, property := range properties {
		switch property := property.(type) {
		case Link:
			if property.Label != "" {
				label = property.Label
			}

			if property.Link != "" {
				link = property.Link
			}

			if property.Payload != "" {
				payload = property.Payload
			}
		case string:
			if len(property) > 0 {
				switch {
				case label == "":
					label = property
				case link == "":
					link = property
				case (payload == "" && property[0] == '{') || (payload == "" && property != ""):
					payload = property
				}
			}
		}
	}

	keyboard.Add(Button{
		Action: Action{Type: "open_link", Label: label, Link: link, Payload: payload},
	})

	return keyboard
}

type Application struct {
	Label         string
	OwnerID       int
	ApplicationID int
	AppID         int
	Payload       string
	Hash          string
}

func (keyboard *ReplyKeyboard) Application(properties ...any) *ReplyKeyboard {
	return keyboard.App(properties...)
}

type applicationProperties struct {
	label   string
	ownerID int
	appID   int
	payload string
	hash    string
}

func (keyboard *ReplyKeyboard) App(properties ...any) *ReplyKeyboard {
	application := new(applicationProperties)

	for _, property := range properties {
		switch property := property.(type) {
		case Application:
			setApplicationStructProperty(application, property)
		case int:
			if application.ownerID == 0 {
				application.ownerID = property
			} else {
				application.appID = property
			}
		case string:
			switch {
			case application.label == "":
				application.label = property
			case (application.payload == "" && property[0] == '{') || (application.payload == "" && property != ""):
				application.payload = property
			case application.hash == "":
				application.hash = property
			}
		}
	}

	keyboard.Add(Button{
		Action: Action{
			Type:          "open_app",
			Label:         application.label,
			OwnerID:       application.ownerID,
			ApplicationID: application.appID,
			Payload:       application.payload,
			Hash:          application.hash,
		},
	})

	return keyboard
}

func setApplicationStructProperty(application *applicationProperties, property Application) {
	if property.Label != "" {
		application.label = property.Label
	}

	if property.OwnerID != 0 {
		application.ownerID = property.OwnerID
	}

	if property.ApplicationID != 0 {
		application.appID = property.ApplicationID
	} else if property.AppID != 0 {
		application.appID = property.ApplicationID
	}

	if property.Payload != "" {
		application.payload = property.Payload
	}

	if property.Hash != "" {
		application.hash = property.Hash
	}
}

type VKPay struct {
	Payload string
	Hash    VKPayHash
}
type VKPayHash struct {
	GroupID     int
	UserID      int
	Action      string
	Amount      int
	Description string
	Data        map[string]any
}

func (keyboard *ReplyKeyboard) Pay(properties ...any) *ReplyKeyboard {
	return keyboard.VKPay(properties...)
}

func (keyboard *ReplyKeyboard) VKPay(properties ...any) *ReplyKeyboard {
	var (
		payload string
		hash    string
	)

	for _, property := range properties {
		switch property := property.(type) {
		case VKPay:
			if property.Payload != "" {
				payload = property.Payload
			}

			if property.Hash.Action != "" {
				data, _ := json.Marshal(property.Hash)
				hash = string(data)
			}
		case VKPayHash:
			switch property.Action {
			case "pay-to-group":
				payToGroup(property)
			case "pay-to-user":
				payToUser(property)
			case "transfer-to-group":
				transferToGroup(property)
			case "transfer-to-user":
				transferToUser(property)
			default:
				action := property.Action
				if action == "" && action[0] == '{' {
					payload = action
				} else if hash == "" {
					hash = action
				}
			}
		}
	}

	keyboard.Add(Button{
		Action: Action{Type: "vkpay", Payload: payload, Hash: hash},
	})

	return keyboard
}

func payToGroup(property VKPayHash) string {
	hash := "action=pay-to-group"

	if property.GroupID != 0 {
		hash += fmt.Sprintf("&group_id=%d", property.GroupID)
	}

	if property.Amount != 0 {
		hash += fmt.Sprintf("&amount=%d", property.Amount)
	}

	if property.Description != "" {
		hash += fmt.Sprintf("&description=%s", url.QueryEscape(property.Description))
	}

	if len(property.Data) != 0 {
		data, _ := json.Marshal(property.Data)
		hash += fmt.Sprintf("&data=%s", url.QueryEscape(string(data)))
	}

	return hash
}

func payToUser(property VKPayHash) string {
	hash := "action=pay-to-user"

	if property.Amount != 0 {
		hash += fmt.Sprintf("&amount=%d", property.Amount)
	}

	if property.Description != "" {
		hash += fmt.Sprintf("&description=%s", url.QueryEscape(property.Description))
	}

	if property.UserID != 0 {
		hash += fmt.Sprintf("&user_id=%d", property.UserID)
	}

	return hash
}

func transferToGroup(property VKPayHash) string {
	hash := "action=transfer-to-group"

	if property.GroupID != 0 {
		hash += fmt.Sprintf("&group_id=%d", property.GroupID)
	}

	if property.Description != "" {
		hash += fmt.Sprintf("&description=%s", url.QueryEscape(property.Description))
	}

	return hash
}

func transferToUser(property VKPayHash) string {
	hash := "action=transfer-to-user"

	if property.UserID != 0 {
		hash += fmt.Sprintf("&user_id=%d", property.GroupID)
	}

	if property.Description != "" {
		hash += fmt.Sprintf("&description=%s", url.QueryEscape(property.Description))
	}

	return hash
}

type Location struct {
	Payload string
}

func (keyboard *ReplyKeyboard) Location(properties ...any) *ReplyKeyboard {
	var payload string

	for _, property := range properties {
		switch property := property.(type) {
		case Location:
			if property.Payload != "" {
				payload = property.Payload
			}
		case string:
			payload = property
		}
	}

	keyboard.Add(Button{
		Action: Action{Type: "location", Payload: payload},
	})

	return keyboard
}

// JSON 🍷🧣🍁 turns an object into a JSON format string.
func (keyboard *ReplyKeyboard) JSON() string {
	if data, err := json.Marshal(keyboard); err == nil {
		return string(data)
	}

	return ""
}

func parsePrimitiveButtonProperty(property any, title, payload, color *string) {
	switch property := property.(type) {
	case Label:
		*title = string(property)
	case Payload:
		value, _ := json.Marshal(property)
		*payload = string(value)
	case Color:
		*color = string(property)
	case Text:
		if property.Text != "" {
			*title = property.Text
		} else if property.Label != "" {
			*title = property.Label
		}

		if property.Payload != nil {
			*payload = payloadToString(property.Payload)
		}

		if property.Color != nil {
			*color = colorToString(property.Color)
		}
	case Callback:
		if property.Text != "" {
			*title = property.Text
		} else if property.Label != "" {
			*title = property.Label
		}

		if property.Payload != nil {
			*payload = payloadToString(property.Payload)
		}

		if property.Color != nil {
			*color = colorToString(property.Color)
		}
	case string:
		if *title == "" {
			*title = property
		} else if len(property) > 0 {
			if (*payload == "" && property[0] == '{') || (*payload == "" && *color != "") {
				*payload = property
			} else {
				*color = property
			}
		}
	}
}

func checkPayload(payload string) string {
	if payload == "" {
		payload = "{}"
	}

	return payload
}

func parsePrimitiveButtonProperties(properties ...any) (title, payload, color string) {
	for _, property := range properties {
		parsePrimitiveButtonProperty(property, &title, &payload, &color)
	}

	payload = checkPayload(payload)

	return title, payload, color
}

func payloadToString(property any) (payload string) {
	switch property := property.(type) {
	case string:
		payload = property
	case Payload:
		value, _ := json.Marshal(property)
		payload = string(value)
	}

	return payload
}

func colorToString(property any) (color string) {
	switch property := property.(type) {
	case string:
		color = property
	case Color:
		color = string(property)
	}

	return color
}
