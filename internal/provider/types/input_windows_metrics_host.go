// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type InputWindowsMetricsHost struct {
	Custom *InputWindowsMetricsCustom `tfsdk:"custom"`
	Mode   types.String               `tfsdk:"mode"`
}
