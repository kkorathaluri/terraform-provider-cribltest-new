// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputAuthTokens struct {
	AllowedIndexesAtToken []types.String              `tfsdk:"allowed_indexes_at_token"`
	AuthType              types.String                `tfsdk:"auth_type"`
	Description           types.String                `tfsdk:"description"`
	Enabled               types.Bool                  `tfsdk:"enabled"`
	Metadata              []Input57AuthTokensMetadata `tfsdk:"metadata"`
	Token                 types.String                `tfsdk:"token"`
	TokenSecret           types.String                `tfsdk:"token_secret"`
}
