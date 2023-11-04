package keyboard

import (
	"encoding/json"
)

func TextButton(properties ...any) *Buttons {
	return NewText(properties...)
}
func NewText(properties ...any) *Buttons {
	keyboard := New().Text(properties...)
	if len(keyboard.Buttons) > 0 {
		return &keyboard.Buttons[0]
	}

	return &Buttons{}
}

func (buttons *Buttons) TextButton(properties ...any) *Buttons {
	return buttons.Text(properties...)
}
func (buttons *Buttons) Text(properties ...any) *Buttons {
	keyboard := New().Text(properties...)
	if len(keyboard.Buttons) > 0 {
		_ = append(*buttons, keyboard.Buttons[0][0])
		return buttons
	}

	return buttons
}

func CallbackButton(properties ...any) *Buttons {
	return NewCallback(properties...)
}
func NewCallback(properties ...any) *Buttons {
	keyboard := New().Callback(properties...)
	if len(keyboard.Buttons) > 0 {
		return &keyboard.Buttons[0]
	}

	return &Buttons{}
}

func (buttons *Buttons) CallbackButton(properties ...any) *Buttons {
	return buttons.Callback(properties...)
}
func (buttons *Buttons) Callback(properties ...any) *Buttons {
	keyboard := New().Callback(properties...)
	if len(keyboard.Buttons) > 0 {
		_ = append(*buttons, keyboard.Buttons[0][0])
		return buttons
	}

	return buttons
}

func LinkButton(properties ...any) *Buttons {
	return NewLink(properties...)
}
func NewLink(properties ...any) *Buttons {
	keyboard := New().Link(properties...)
	if len(keyboard.Buttons) > 0 {
		return &keyboard.Buttons[0]
	}

	return &Buttons{}
}

func (buttons *Buttons) LinkButton(properties ...any) *Buttons {
	return buttons.Link(properties...)
}
func (buttons *Buttons) Link(properties ...any) *Buttons {
	keyboard := New().Link(properties...)
	if len(keyboard.Buttons) > 0 {
		_ = append(*buttons, keyboard.Buttons[0][0])
		return buttons
	}

	return buttons
}

func ApplicationButton(properties ...any) *Buttons {
	return NewApp(properties...)
}
func AppButton(properties ...any) *Buttons {
	return NewApp(properties...)
}

func NewApplication(properties ...any) *Buttons {
	return NewApp(properties...)
}
func NewApp(properties ...any) *Buttons {
	keyboard := New().App(properties...)
	if len(keyboard.Buttons) > 0 {
		return &keyboard.Buttons[0]
	}

	return &Buttons{}
}

func (buttons *Buttons) ApplicationButton(properties ...any) *Buttons {
	return buttons.App(properties...)
}
func (buttons *Buttons) AppButton(properties ...any) *Buttons {
	return buttons.App(properties...)
}

func (buttons *Buttons) Application(properties ...any) *Buttons {
	return buttons.App(properties...)
}
func (buttons *Buttons) App(properties ...any) *Buttons {
	keyboard := New().App(properties...)
	if len(keyboard.Buttons) > 0 {
		_ = append(*buttons, keyboard.Buttons[0][0])
		return buttons
	}

	return buttons
}

func LocationButton(properties ...any) *Buttons {
	return NewLocation(properties...)
}
func NewLocation(properties ...any) *Buttons {
	keyboard := New().Location(properties...)
	if len(keyboard.Buttons) > 0 {
		return &keyboard.Buttons[0]
	}

	return &Buttons{}
}

func (buttons *Buttons) LocationButton(properties ...any) *Buttons {
	return buttons.Location(properties...)
}
func (buttons *Buttons) Location(properties ...any) *Buttons {
	keyboard := New().Location(properties...)
	if len(keyboard.Buttons) > 0 {
		_ = append(*buttons, keyboard.Buttons[0][0])
		return buttons
	}

	return buttons
}

func (buttons *Buttons) JSON() string {
	if data, err := json.Marshal(buttons); err == nil {
		return string(data)
	}

	return ""
}
