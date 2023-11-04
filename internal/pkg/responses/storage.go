package responses

type StorageKeys struct {
	Error    *Error   `json:"error"`
	Response []string `json:"response,omitempty"`
}

type Storage struct {
	Error    *Error   `json:"error"`
	Response Storages `json:"response,omitempty"`
}

type Storages []struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type SetStorage struct {
	Error    *Error `json:"error"`
	Response any    `json:"response,omitempty"`
}
