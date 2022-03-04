package request

type CopywritingAddRequest struct {
	Application string        `json:"application,omitempty"`
	Type        string        `json:"type,omitempty"`
	Path        string        `json:"path,omitempty"`
	Key         string        `json:"key,omitempty"`
	Copywriting []Copywriting `json:"copywriting,omitempty"`
}

type Copywriting struct {
	LanguageCode  string `json:"languageCode,omitempty"`
	LanguageValue string `json:"languageValue,omitempty"`
}
