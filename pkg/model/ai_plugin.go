package model

// AiPlugin is the model of ai plugin.
type AiPlugin struct {
	SchemaVersion       string       `json:"schema_version,omitempty" yaml:"schema_version,omitempty"`
	NameForModel        string       `json:"name_for_model,omitempty" yaml:"name_for_model,omitempty"`
	NameForHuman        string       `json:"name_for_human,omitempty" yaml:"name_for_human,omitempty"`
	DescriptionForModel string       `json:"description_for_model,omitempty" yaml:"description_for_model,omitempty"`
	DescriptionForHuman string       `json:"description_for_human,omitempty" yaml:"description_for_human,omitempty"`
	Auth                ManifestAuth `json:"auth,omitempty" yaml:"auth,omitempty"`
	API                 APIObject    `json:"api,omitempty" yaml:"api,omitempty"`
	LogoURL             string       `json:"logo_url,omitempty" yaml:"logo_url,omitempty"`
	ContactEmail        string       `json:"contact_email,omitempty" yaml:"contact_email,omitempty"`
	LegalInfoURL        string       `json:"legal_info_url,omitempty" yaml:"legal_info_url,omitempty"`
}

type ManifestAuth struct {
	Type ManifestAuthType `json:"type,omitempty" yaml:"type,omitempty"`
}

type ManifestAuthType string

const (
	None        ManifestAuthType = "none"
	UserHTTP    ManifestAuthType = "user_http"
	ServiceHTTP ManifestAuthType = "service_http"
	OaAuth      ManifestAuthType = "oauth"
)

type HTTPAuthorizationType string

const (
	Bearer HTTPAuthorizationType = "bearer"
	Basic  HTTPAuthorizationType = "basic"
)

type APIObject struct {
	Type                string `json:"type,omitempty" yaml:"type,omitempty"`
	URL                 string `json:"url,omitempty" yaml:"url,omitempty"`
	IsUserAuthenticated bool   `json:"is_user_authenticated,omitempty" yaml:"is_user_authenticated,omitempty"`
}
