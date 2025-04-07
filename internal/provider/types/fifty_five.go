// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type FiftyFive struct {
	Connections          []Input55Connections `tfsdk:"connections"`
	Description          types.String         `tfsdk:"description"`
	Disabled             types.Bool           `tfsdk:"disabled"`
	EnablePassThrough    types.Bool           `tfsdk:"enable_pass_through"`
	Environment          types.String         `tfsdk:"environment"`
	Host                 types.String         `tfsdk:"host"`
	ID                   types.String         `tfsdk:"id"`
	IPAllowlistRegex     types.String         `tfsdk:"ip_allowlist_regex"`
	IPDenylistRegex      types.String         `tfsdk:"ip_denylist_regex"`
	IpfixEnabled         types.Bool           `tfsdk:"ipfix_enabled"`
	Metadata             []Input55Metadata    `tfsdk:"metadata"`
	Pipeline             types.String         `tfsdk:"pipeline"`
	Port                 types.Number         `tfsdk:"port"`
	Pq                   *Input55Pq           `tfsdk:"pq"`
	PqEnabled            types.Bool           `tfsdk:"pq_enabled"`
	SendToRoutes         types.Bool           `tfsdk:"send_to_routes"`
	Status               *Input55Status       `tfsdk:"status"`
	Streamtags           []types.String       `tfsdk:"streamtags"`
	TemplateCacheMinutes types.Number         `tfsdk:"template_cache_minutes"`
	Type                 types.String         `tfsdk:"type"`
	UDPSocketRxBufSize   types.Number         `tfsdk:"udp_socket_rx_buf_size"`
	V5Enabled            types.Bool           `tfsdk:"v5_enabled"`
	V9Enabled            types.Bool           `tfsdk:"v9_enabled"`
}
