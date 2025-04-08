// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputSplunkLb struct {
	AuthToken                    types.String                         `tfsdk:"auth_token"`
	AuthType                     types.String                         `tfsdk:"auth_type"`
	Compress                     types.String                         `tfsdk:"compress"`
	ConnectionTimeout            types.Float64                        `tfsdk:"connection_timeout"`
	Description                  types.String                         `tfsdk:"description"`
	DNSResolvePeriodSec          types.Float64                        `tfsdk:"dns_resolve_period_sec"`
	EnableACK                    types.Bool                           `tfsdk:"enable_ack"`
	EnableMultiMetrics           types.Bool                           `tfsdk:"enable_multi_metrics"`
	Environment                  types.String                         `tfsdk:"environment"`
	ExcludeSelf                  types.Bool                           `tfsdk:"exclude_self"`
	Hosts                        []Hosts                              `tfsdk:"hosts"`
	ID                           types.String                         `tfsdk:"id"`
	IndexerDiscovery             types.Bool                           `tfsdk:"indexer_discovery"`
	IndexerDiscoveryConfigs      *IndexerDiscoveryConfigs             `tfsdk:"indexer_discovery_configs"`
	LoadBalanceStatsPeriodSec    types.Float64                        `tfsdk:"load_balance_stats_period_sec"`
	LogFailedRequests            types.Bool                           `tfsdk:"log_failed_requests"`
	MaxConcurrentSenders         types.Float64                        `tfsdk:"max_concurrent_senders"`
	MaxFailedHealthChecks        types.Float64                        `tfsdk:"max_failed_health_checks"`
	MaxS2Sversion                types.String                         `tfsdk:"max_s2_sversion"`
	NestedFields                 types.String                         `tfsdk:"nested_fields"`
	OnBackpressure               types.String                         `tfsdk:"on_backpressure"`
	Pipeline                     types.String                         `tfsdk:"pipeline"`
	PqCompress                   types.String                         `tfsdk:"pq_compress"`
	PqControls                   *OutputSplunkLbPqControls            `tfsdk:"pq_controls"`
	PqMaxFileSize                types.String                         `tfsdk:"pq_max_file_size"`
	PqMaxSize                    types.String                         `tfsdk:"pq_max_size"`
	PqMode                       types.String                         `tfsdk:"pq_mode"`
	PqOnBackpressure             types.String                         `tfsdk:"pq_on_backpressure"`
	PqPath                       types.String                         `tfsdk:"pq_path"`
	SenderUnhealthyTimeAllowance types.Float64                        `tfsdk:"sender_unhealthy_time_allowance"`
	Streamtags                   []types.String                       `tfsdk:"streamtags"`
	SystemFields                 []types.String                       `tfsdk:"system_fields"`
	TextSecret                   types.String                         `tfsdk:"text_secret"`
	ThrottleRatePerSec           types.String                         `tfsdk:"throttle_rate_per_sec"`
	TLS                          *OutputSplunkLbTLSSettingsClientSide `tfsdk:"tls"`
	Type                         types.String                         `tfsdk:"type"`
	WriteTimeout                 types.Float64                        `tfsdk:"write_timeout"`
}
