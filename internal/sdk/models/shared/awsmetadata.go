// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AwsMetadataTags struct {
}

type AwsMetadata struct {
	Hostname       *string          `json:"hostname,omitempty"`
	Identity       map[string]any   `json:"identity,omitempty"`
	PublicIpv4     *string          `json:"public_ipv4,omitempty"`
	Roles          []string         `json:"roles,omitempty"`
	SecurityGroups []string         `json:"security_groups,omitempty"`
	Tags           *AwsMetadataTags `json:"tags,omitempty"`
	Version        float64          `json:"version"`
}

func (o *AwsMetadata) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *AwsMetadata) GetIdentity() map[string]any {
	if o == nil {
		return nil
	}
	return o.Identity
}

func (o *AwsMetadata) GetPublicIpv4() *string {
	if o == nil {
		return nil
	}
	return o.PublicIpv4
}

func (o *AwsMetadata) GetRoles() []string {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *AwsMetadata) GetSecurityGroups() []string {
	if o == nil {
		return nil
	}
	return o.SecurityGroups
}

func (o *AwsMetadata) GetTags() *AwsMetadataTags {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AwsMetadata) GetVersion() float64 {
	if o == nil {
		return 0.0
	}
	return o.Version
}
