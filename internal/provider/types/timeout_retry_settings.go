// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type TimeoutRetrySettings struct {
	BackoffRate    types.Float64 `tfsdk:"backoff_rate"`
	InitialBackoff types.Float64 `tfsdk:"initial_backoff"`
	MaxBackoff     types.Float64 `tfsdk:"max_backoff"`
	TimeoutRetry   types.Bool    `tfsdk:"timeout_retry"`
}
