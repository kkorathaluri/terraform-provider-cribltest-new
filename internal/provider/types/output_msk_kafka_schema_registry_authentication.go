// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputMskKafkaSchemaRegistryAuthentication struct {
	Auth                 *OutputMskAuth                                     `tfsdk:"auth"`
	ConnectionTimeout    types.Float64                                      `tfsdk:"connection_timeout"`
	DefaultKeySchemaID   types.Float64                                      `tfsdk:"default_key_schema_id"`
	DefaultValueSchemaID types.Float64                                      `tfsdk:"default_value_schema_id"`
	Disabled             types.Bool                                         `tfsdk:"disabled"`
	MaxRetries           types.Float64                                      `tfsdk:"max_retries"`
	RequestTimeout       types.Float64                                      `tfsdk:"request_timeout"`
	SchemaRegistryURL    types.String                                       `tfsdk:"schema_registry_url"`
	TLS                  *OutputMskKafkaSchemaRegistryTLSSettingsClientSide `tfsdk:"tls"`
}
