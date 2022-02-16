package entity

import "medusa-globalization-copywriting-system/tools/language"

type CopywritingAddRequest struct {
	Application string            `json:"application,omitempty"`
	Type        string            `json:"type,omitempty"`
	Path        string            `json:"path,omitempty"`
	Language    language.Language `json:"language"`
	Key         string            `json:"key,omitempty"`
	Copywriting string            `json:"copywriting,omitempty"`
}
