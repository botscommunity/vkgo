package responses

type OffersReply struct {
	Error    *Error `json:"error"`
	Response Offers `json:"response"`
}

type Offers struct {
	Error  *Error `json:"error"`
	Count  int    `json:"count,omitempty"`
	Offers []int  `json:"items,omitempty"`
}
