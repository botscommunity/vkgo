package responses

type ResolveDomainReply struct {
	Error    *Error        `json:"error"`
	Response ResolveDomain `json:"response"`
}

type ResolveDomain struct {
	Error    *Error `json:"error"`
	Type     string `json:"type,omitempty"`
	ObjectID int    `json:"object_id,omitempty"`
}
