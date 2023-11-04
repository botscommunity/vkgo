package responses

type HealthReply struct {
	Error    *Error `json:"error"`
	Response Health `json:"response"`
}

type Health struct {
	Error    *Error         `json:"error"`
	Statuses []HealthStatus `json:"health_statuses,omitempty"`
}

type HealthStatus struct {
	AverageTime int    `json:"average_response_time,omitempty"`
	Category    string `json:"category,omitempty"`
	Status      string `json:"status,omitempty"`
	Uptime      any    `json:"uptime,omitempty"`
}
