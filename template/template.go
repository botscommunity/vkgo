package template

import (
	"encoding/json"

	"github.com/botscommunity/vkgo/keyboard"
)

// Template 🍂 presents VK API template objects
// ☠️  Type specifies the type of template to create. Available Type: «carousel»
// 🎃 Elements - the list of Element objects arranged horizontally in the carousel.
type Template struct {
	Type     string    `json:"type"`
	Elements []Element `json:"elements"`
}

// Element 💐 is a carousel object that is displayed horizontally
// 🌈 Title - element header. Limited to 80 characters
// 🐰 Description - element description. Limited to 80 characters
// 🌷 PhotoID - VK photo ID format «#owner_id_#photo_id»
// 🐇 Action - the action that should happen when you click on the element
// 🐣 Buttons - pointer to the buttons from the «Keyboard» package, which are arranged vertically.
type Element struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	PhotoID     string            `json:"photo_id"`
	Action      Action            `json:"action"`
	Buttons     *keyboard.Buttons `json:"buttons"`
}

// Action - 🌲 action when clicking on a carousel element
// 🌱 Type - action type. Available: open_link or open_photo
// 🌳 OpenLink opens a link, OpenPhoto opens a photo of the object
// 👩‍🌾 Link - link with the type «open_link». Important: not «open_photo».
type Action struct {
	Type string `json:"type"`
	Link string `json:"link,omitempty"`
}

// New 🌷 is responsible for creating the template
// 🌈 It takes String - type Template and optional type []Element
// 🌺 Returns a pointer to the Template object.
func New(properties ...any) *Template {
	template := Template{Elements: []Element{}}

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			template.Type = property
		case []Element:
			template.Elements = property
		}
	}

	return &template
}

// JSON 🍷🧣🍁 turns an object into a JSON format string.
func (template *Template) JSON() string {
	data, err := json.Marshal(template)
	if err != nil {
		return ""
	}

	return string(data)
}
