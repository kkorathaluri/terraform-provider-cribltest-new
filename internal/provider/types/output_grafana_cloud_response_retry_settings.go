// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputGrafanaCloudResponseRetrySettings struct {
	BackoffRate    types.Float64 `tfsdk:"backoff_rate"`
	HTTPStatus     types.Float64 `tfsdk:"http_status"`
	InitialBackoff types.Float64 `tfsdk:"initial_backoff"`
	MaxBackoff     types.Float64 `tfsdk:"max_backoff"`
}
