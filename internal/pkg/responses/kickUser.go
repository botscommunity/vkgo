package responses

type KickUser struct {
	Error *Error `json:"error"`
	Code  int    `json:"response"`
}
