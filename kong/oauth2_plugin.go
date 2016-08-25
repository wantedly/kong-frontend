package kong

type OAuth2PluginConfig struct {
	AcceptHTTPIfAlreadyTerminated bool   `form:"accept_http_if_already_terminated" json:"accept_http_if_already_terminated,omitempty"`
	EnableAuthorizationCode       bool   `form:"enable_authorization_code" json:"enable_authorization_code,omitempty"`
	EnableClientCredentials       bool   `form:"enable_client_credentials" json:"enable_client_credentials,omitempty"`
	EnableImplicitGrant           bool   `form:"enable_implicit_grant" json:"enable_implicit_grant,omitempty"`
	EnablePasswordGrant           bool   `form:"enable_password_grant" json:"enable_password_grant,omitempty"`
	HideCredentials               bool   `form:"hide_credentials" json:"hide_credentials,omitempty"`
	MandatoryScope                bool   `form:"mandatory_scope" json:"mandatory_scope,omitempty"`
	ProvisionKey                  string `form:"provision_key" json:"provision_key,omitempty"`
	TokenExpiration               int    `form:"token_expiration" json:"token_expiration,omitempty"`
	Scopes                        string `form:"scopes" json:"scopes,omitempty"`
}
