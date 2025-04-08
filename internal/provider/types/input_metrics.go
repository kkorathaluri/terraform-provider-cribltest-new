// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputMetrics struct {
	Connections        []InputMetricsConnections          `tfsdk:"connections"`
	Description        types.String                       `tfsdk:"description"`
	Disabled           types.Bool                         `tfsdk:"disabled"`
	EnableProxyHeader  types.Bool                         `tfsdk:"enable_proxy_header"`
	Environment        types.String                       `tfsdk:"environment"`
	Host               types.String                       `tfsdk:"host"`
	ID                 types.String                       `tfsdk:"id"`
	IPWhitelistRegex   types.String                       `tfsdk:"ip_whitelist_regex"`
	MaxBufferSize      types.Float64                      `tfsdk:"max_buffer_size"`
	Metadata           []InputMetricsMetadata             `tfsdk:"metadata"`
	Pipeline           types.String                       `tfsdk:"pipeline"`
	Pq                 *InputMetricsPq                    `tfsdk:"pq"`
	PqEnabled          types.Bool                         `tfsdk:"pq_enabled"`
	SendToRoutes       types.Bool                         `tfsdk:"send_to_routes"`
	Streamtags         []types.String                     `tfsdk:"streamtags"`
	TCPPort            types.Float64                      `tfsdk:"tcp_port"`
	TLS                *InputMetricsTLSSettingsServerSide `tfsdk:"tls"`
	Type               types.String                       `tfsdk:"type"`
	UDPPort            types.Float64                      `tfsdk:"udp_port"`
	UDPSocketRxBufSize types.Float64                      `tfsdk:"udp_socket_rx_buf_size"`
}
