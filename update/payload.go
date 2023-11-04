package update

import (
	"encoding/json"
	"fmt"
)

type rawPayload any

func DecodePayload(rawPayload rawPayload) (map[string]string, error) {
	payload := map[string]string{}

	switch rp := rawPayload.(type) {
	case string:
		object := map[string]any{}

		if err := json.Unmarshal([]byte(rp), &object); err != nil {
			return payload, err
		}

		for name, value := range object {
			payload[name] = fmt.Sprintf("%v", value)
		}

		return payload, nil
	case map[string]any:
		var (
			object    = map[string]any{}
			body, err = json.Marshal(rp)
		)

		if err != nil {
			return payload, err
		}

		if err := json.Unmarshal(body, &object); err != nil {
			return payload, err
		}

		for name, value := range object {
			payload[name] = fmt.Sprintf("%v", value)
		}
	}

	return payload, nil
}
