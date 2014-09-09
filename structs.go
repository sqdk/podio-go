package podiogo

type AuthMessage struct {
	AccessToken  string         `json:"access_token"`
	TokenType    string         `json:"token_type"`
	ExpiresIn    int            `json:"expires_in"`
	RefreshToken string         `json:"refresh_token"`
	Reference    PodioReference `json:"ref"`
}

type PodioReference struct {
	Type string `json:"type"`
	Id   int    `json:"id"`
}

type PodioResponse struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

type PodioId struct {
	ProfileId int64 `json:"profile_id"`
}

type PodioItemCreateResponse struct {
	ItemId    int64  `json:"item_id"`
	ItemTitle string `json:"title"`
}

type PodioItemUpdateResponse struct {
	ItemRevision int64  `json:"revision"`
	ItemTitle    string `json:"title"`
}

type PodioEmbed struct {
	EmbedId int64 `json:"embed_id"`
}
