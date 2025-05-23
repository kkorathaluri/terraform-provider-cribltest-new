// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OutputClickHouseTLSSettingsClientSide struct {
	CaPath          types.String `tfsdk:"ca_path"`
	CertificateName types.String `tfsdk:"certificate_name"`
	CertPath        types.String `tfsdk:"cert_path"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	MaxVersion      types.String `tfsdk:"max_version"`
	MinVersion      types.String `tfsdk:"min_version"`
	Passphrase      types.String `tfsdk:"passphrase"`
	PrivKeyPath     types.String `tfsdk:"priv_key_path"`
	Servername      types.String `tfsdk:"servername"`
}
