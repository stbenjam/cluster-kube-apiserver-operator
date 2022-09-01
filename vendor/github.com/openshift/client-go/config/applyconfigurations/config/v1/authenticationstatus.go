// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AuthenticationStatusApplyConfiguration represents an declarative configuration of the AuthenticationStatus type for use
// with apply.
type AuthenticationStatusApplyConfiguration struct {
	IntegratedOAuthMetadata *ConfigMapNameReferenceApplyConfiguration `json:"integratedOAuthMetadata,omitempty"`
}

// AuthenticationStatusApplyConfiguration constructs an declarative configuration of the AuthenticationStatus type for use with
// apply.
func AuthenticationStatus() *AuthenticationStatusApplyConfiguration {
	return &AuthenticationStatusApplyConfiguration{}
}

// WithIntegratedOAuthMetadata sets the IntegratedOAuthMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IntegratedOAuthMetadata field is set to the value of the last call.
func (b *AuthenticationStatusApplyConfiguration) WithIntegratedOAuthMetadata(value *ConfigMapNameReferenceApplyConfiguration) *AuthenticationStatusApplyConfiguration {
	b.IntegratedOAuthMetadata = value
	return b
}
