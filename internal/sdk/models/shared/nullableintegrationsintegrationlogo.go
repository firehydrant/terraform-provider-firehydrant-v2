// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NullableIntegrationsIntegrationLogo struct {
	LogoURL *string `json:"logo_url,omitempty"`
}

func (o *NullableIntegrationsIntegrationLogo) GetLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoURL
}
