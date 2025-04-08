// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ContentConfig struct {
	ContentType types.String  `tfsdk:"content_type"`
	Description types.String  `tfsdk:"description"`
	Enabled     types.Bool    `tfsdk:"enabled"`
	Interval    types.Float64 `tfsdk:"interval"`
	LogLevel    types.String  `tfsdk:"log_level"`
}
