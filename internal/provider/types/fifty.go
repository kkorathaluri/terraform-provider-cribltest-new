// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Fifty struct {
	AllowMachineIDMismatch types.Bool           `tfsdk:"allow_machine_id_mismatch"`
	AuthMethod             types.String         `tfsdk:"auth_method"`
	CaFingerprint          types.String         `tfsdk:"ca_fingerprint"`
	CaptureHeaders         types.Bool           `tfsdk:"capture_headers"`
	Connections            []Input50Connections `tfsdk:"connections"`
	Description            types.String         `tfsdk:"description"`
	Disabled               types.Bool           `tfsdk:"disabled"`
	EnableHealthCheck      types.Bool           `tfsdk:"enable_health_check"`
	EnableProxyHeader      types.Bool           `tfsdk:"enable_proxy_header"`
	Environment            types.String         `tfsdk:"environment"`
	Host                   types.String         `tfsdk:"host"`
	ID                     types.String         `tfsdk:"id"`
	IPAllowlistRegex       types.String         `tfsdk:"ip_allowlist_regex"`
	IPDenylistRegex        types.String         `tfsdk:"ip_denylist_regex"`
	KeepAliveTimeout       types.Number         `tfsdk:"keep_alive_timeout"`
	Keytab                 types.String         `tfsdk:"keytab"`
	MaxActiveReq           types.Number         `tfsdk:"max_active_req"`
	MaxRequestsPerSocket   types.Int64          `tfsdk:"max_requests_per_socket"`
	Metadata               []Input50Metadata    `tfsdk:"metadata"`
	Pipeline               types.String         `tfsdk:"pipeline"`
	Port                   types.Number         `tfsdk:"port"`
	Pq                     *Input50Pq           `tfsdk:"pq"`
	PqEnabled              types.Bool           `tfsdk:"pq_enabled"`
	Principal              types.String         `tfsdk:"principal"`
	SendToRoutes           types.Bool           `tfsdk:"send_to_routes"`
	SocketTimeout          types.Number         `tfsdk:"socket_timeout"`
	Status                 *Input50Status       `tfsdk:"status"`
	Streamtags             []types.String       `tfsdk:"streamtags"`
	Subscriptions          []Subscriptions      `tfsdk:"subscriptions"`
	TLS                    *MTLSSettings        `tfsdk:"tls"`
	Type                   types.String         `tfsdk:"type"`
}
