package config

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

const OIDCDefaultScopes = "openid email profile"

// OIDCEnabled checks if login via OpenID Connect (OIDC) is enabled.
func (c *Config) OIDCEnabled() bool {
	return c.options.OIDCUri != "" && c.options.OIDCClient != "" && c.options.OIDCSecret != ""
}

// OIDCUri returns the OpenID Connect issuer URI as *url.URL for single sign-on via OIDC.
func (c *Config) OIDCUri() *url.URL {
	if uri := c.options.OIDCUri; uri == "" {
		return &url.URL{}
	} else if result, err := url.Parse(uri); err != nil {
		log.Errorf("oidc: failed to parse issuer URI (%s)", err)
		return &url.URL{}
	} else {
		return result
	}
}

// OIDCInsecure checks if OIDC issuer SSL/TLS certificate verification should be skipped.
func (c *Config) OIDCInsecure() bool {
	return c.options.OIDCInsecure
}

// OIDCClient returns the Client ID for single sign-on via OIDC.
func (c *Config) OIDCClient() string {
	return c.options.OIDCClient
}

// OIDCSecret returns the Client ID for single sign-on via OIDC.
func (c *Config) OIDCSecret() string {
	return c.options.OIDCSecret
}

// OIDCScopes returns the user information scopes for single sign-on via OIDC.
func (c *Config) OIDCScopes() string {
	if c.options.OIDCScopes == "" {
		return OIDCDefaultScopes
	}

	return c.options.OIDCScopes
}

// OIDCRegister checks if new accounts may be created via OIDC.
func (c *Config) OIDCRegister() bool {
	return c.options.OIDCRegister
}

// OIDCRedirect checks if unauthenticated users should automatically be redirected to the OIDC login page.
func (c *Config) OIDCRedirect() bool {
	return c.options.OIDCRedirect
}

// OIDCProvider returns the OIDC provider name, if any.
func (c *Config) OIDCProvider() string {
	return c.options.OIDCProvider
}

// OIDCProviderIcon returns the OIDC provider icon URI, if any.
func (c *Config) OIDCProviderIcon() string {
	return c.options.OIDCProviderIcon
}

// OIDCReport returns the OpenID Connect config values as a table for reporting.
func (c *Config) OIDCReport() (rows [][]string, cols []string) {
	cols = []string{"Name", "Value"}

	rows = [][]string{
		{"oidc-uri", c.OIDCUri().String()},
		{"oidc-insecure", fmt.Sprintf("%t", c.OIDCInsecure())},
		{"oidc-client", c.OIDCClient()},
		{"oidc-secret", strings.Repeat("*", utf8.RuneCountInString(c.OIDCSecret()))},
		{"oidc-scopes", c.OIDCScopes()},
		{"oidc-register", fmt.Sprintf("%t", c.OIDCRegister())},
		{"oidc-redirect", fmt.Sprintf("%t", c.OIDCRedirect())},
		{"oidc-provider", c.OIDCProvider()},
		{"oidc-provider-icon", c.OIDCProviderIcon()},
	}

	return rows, cols
}
