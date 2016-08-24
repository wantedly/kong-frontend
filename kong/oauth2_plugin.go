package kong

type OAuth2PluginConfig struct {
	AcceptHTTPIfAlreadyTerminated bool   `json:"accept_http_if_already_terminated,omitempty"`
	EnableAuthorizationCode       bool   `json:"enable_authorization_code,omitempty"`
	EnableClientCredentials       bool   `json:"enable_client_credentials,omitempty"`
	EnableImplicitGrant           bool   `json:"enable_implicit_grant,omitempty"`
	EnablePasswordGrant           bool   `json:"enable_password_grant,omitempty"`
	HideCredentials               bool   `json:"hide_credentials,omitempty"`
	MandatoryScope                bool   `json:"mandatory_scope,omitempty"`
	ProvisionKey                  string `json:"provision_key,omitempty"`
	TokenExpiration               int    `json:"token_expiration,omitempty"`
}
