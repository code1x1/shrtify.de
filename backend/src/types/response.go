package types

type CreateShortcodeResponsePayload struct {
	Meta     Meta                    `json:"meta"`
	Response CreateShortcodeResponse `json:"response"`
}
type CreateShortcodeResponse struct {
	ShortUrl string `json:"shortUrl"`
}

type GetShortcodesResponsePayload struct {
	Meta     Meta                    `json:"meta"`
	Response []GetShortcodesResponse `json:"response"`
}
type GetShortcodesResponse struct {
	ShortUrl    string `json:"shortUrl"`
	OriginalUrl string `json:"originalUrl"`
}

type GetShortcodesMeta struct {
	Pages int64 `json:"pages"`
}
