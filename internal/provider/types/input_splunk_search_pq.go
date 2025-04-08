// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputSplunkSearchPq struct {
	CommitFrequency types.Float64 `tfsdk:"commit_frequency"`
	Compress        types.String  `tfsdk:"compress"`
	MaxBufferSize   types.Float64 `tfsdk:"max_buffer_size"`
	MaxFileSize     types.String  `tfsdk:"max_file_size"`
	MaxSize         types.String  `tfsdk:"max_size"`
	Mode            types.String  `tfsdk:"mode"`
	Path            types.String  `tfsdk:"path"`
}
