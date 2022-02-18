package entity

type CopywritingAddRequest struct {
	Application  string `json:"application,omitempty"`
	Type         string `json:"type,omitempty"`
	Path         string `json:"path,omitempty"`
	LanguageCode string `json:"language"`
	Key          string `json:"key,omitempty"`
	Copywriting  string `json:"copywriting,omitempty"`
}
