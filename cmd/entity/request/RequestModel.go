package request

type ApplicationAddRequest struct {
	ApplicationName           string `json:"applicationName,omitempty"`
	ApplicationType           string `json:"applicationType,omitempty"`
	ApplicationAdministrators string `json:"applicationAdministrators,omitempty"`
}

type DocumentAddRequest struct {
	Application string     `json:"application,omitempty"`
	Type        string     `json:"type,omitempty"`
	Path        string     `json:"path,omitempty"`
	Key         string     `json:"key,omitempty"`
	Document    []Document `json:"document,omitempty"`
}

type Document struct {
	CountryCode   string `json:"countryCode,omitempty"`
	LanguageValue string `json:"languageValue,omitempty"`
}
