package template

import (
	"fmt"

	"github.com/botscommunity/vkgo/keyboard"
)

// NewCarousel 🐸 creates a carousel object based on the Template structure
// 🍄 This is the easiest way to work with carousels
// 🌷 Does not require manual placement of elements, specifying the type, etc.
// 🌈 Pre-created elements can be used as arguments
// 🦌 Always returns a pointer to a Template object 🦋.
func NewCarousel(properties ...any) *Template {
	properties = append(properties, "carousel")
	return New(properties...)
}

// Title - 🍥 element header.
type Title string

// Description - 🍜 element description.
type Description string

// PhotoID - 🦊 VK photo ID format «#owner_id_#photo_id».
type PhotoID string

// Photo - ⛩️ VK photo ID format «#owner_id_#photo_id».
type Photo string

// Element - adding a new element to the template
// Accepted values:
// String - 1: Title, 2 - Description, 3 - PhotoID
// Element - element of the element to be added
// Action - action when clicking on the Action element
// keyboard.Buttons - vertical set of buttons
// *keyboard.Buttons - the same, but via the pointer
// Also available are custom types which are listed in the String argument
// Returns modified Template object with a new element.
func (template *Template) Element(properties ...any) *Template {
	element := Element{}

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			switch {
			case element.Title == "":
				element.Title = property
			case element.Description == "":
				element.Description = property
			default:
				element.PhotoID = property
			}
		case Element:
			element = property
		case Action:
			element.Action = property
		case keyboard.Buttons:
			value := property
			element.Buttons = &value
		case *keyboard.Buttons:
			element.Buttons = property
		case Title:
			element.Title = string(property)
		case Description:
			element.Description = string(property)
		case Photo:
			element.PhotoID = string(property)
		case PhotoID:
			element.PhotoID = string(property)
		}
	}

	return template.Add(element)
}

func (template *Template) Title(properties ...any) *Template {
	var title string

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			title = property
		case Title:
			title = string(property)
		}
	}

	if count := len(template.Elements); count > 0 {
		element := &template.Elements[count-1]
		element.Title = title
	}

	return template
}

func (template *Template) Description(properties ...any) *Template {
	var description string

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			description = property
		case Description:
			description = string(property)
		}
	}

	if count := len(template.Elements); count > 0 {
		element := &template.Elements[count-1]
		element.Description = description
	}

	return template
}

func (template *Template) Photo(properties ...any) *Template {
	var photo string

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			photo = property
		case Title:
			photo = string(property)
		}
	}

	if count := len(template.Elements); count > 0 {
		element := &template.Elements[count-1]
		element.PhotoID = photo
	}

	return template
}

// Add append an element to the template Elements.
func (template *Template) Add(element Element) *Template {
	template.Elements = append(template.Elements, element)
	return template
}

const (
	LINK  = "open_link"
	PHOTO = "open_photo"
)

// Action - 🌴 method of creating an action with a template element
// 🏖️ Accepts the following arguments:
// 🍹 Action Structure
// 😎 String - 1: Type, 2: Link
// 🚎 And the custom types described in the String argument
// 🦩 Returns the modified pattern by setting the Action to the last element.
func (template *Template) Action(properties ...any) *Template {
	action := Action{}

	for _, property := range properties {
		if typ := fmt.Sprintf("%T", property); typ == "string" {
			if action.Type == "" {
				action.Type = property.(string)
			} else {
				action.Link = property.(string)
			}
		}
	}

	if count := len(template.Elements); count > 0 {
		element := &template.Elements[count-1]
		element.Action = action
	}

	return template
}

// Buttons - 🤡 add buttons to the last created carousel element
// 🍔 Accepts arguments: * Keyboard.Buttons
// 🍟 Returns the modified template.
func (template *Template) Buttons(properties ...any) *Template {
	buttons := &keyboard.Buttons{}

	for _, property := range properties {
		if property, ok := property.(*keyboard.Buttons); ok {
			buttons = property
		}
	}

	if count := len(template.Elements); count > 0 {
		element := &template.Elements[count-1]
		element.Buttons = buttons
	}

	return template
}
