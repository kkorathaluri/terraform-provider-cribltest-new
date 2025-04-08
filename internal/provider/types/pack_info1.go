// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type PackInfo1 struct {
	Author              types.String            `tfsdk:"author"`
	Description         types.String            `tfsdk:"description"`
	DisplayName         types.String            `tfsdk:"display_name"`
	Exports             []types.String          `tfsdk:"exports"`
	ID                  types.String            `tfsdk:"id"`
	IsDisabled          types.Bool              `tfsdk:"is_disabled"`
	MinLogStreamVersion types.String            `tfsdk:"min_log_stream_version"`
	Settings            map[string]types.String `tfsdk:"settings"`
	Source              types.String            `tfsdk:"source"`
	Spec                types.String            `tfsdk:"spec"`
	Tags                *PackInfoTags           `tfsdk:"tags"`
	Version             types.String            `tfsdk:"version"`
}
