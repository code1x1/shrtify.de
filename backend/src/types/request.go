package types

type Meta struct {
	Status int     `json:"status"`
	Error  *string `json:"error"`
	More   any     `json:"more"`
}

type CreateShortcodeRequestPayload struct {
	Request CreateShortcodeRequest `json:"request"`
}
type CreateShortcodeRequest struct {
	Url  string `json:"url"`
	Code string `json:"code"`
}
