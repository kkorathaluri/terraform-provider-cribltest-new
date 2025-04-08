// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputDataset struct {
	APIKey                        types.String                         `tfsdk:"api_key"`
	AuthType                      types.String                         `tfsdk:"auth_type"`
	Compress                      types.Bool                           `tfsdk:"compress"`
	Concurrency                   types.Float64                        `tfsdk:"concurrency"`
	CustomURL                     types.String                         `tfsdk:"custom_url"`
	DefaultSeverity               types.String                         `tfsdk:"default_severity"`
	Description                   types.String                         `tfsdk:"description"`
	Environment                   types.String                         `tfsdk:"environment"`
	ExcludeFields                 []types.String                       `tfsdk:"exclude_fields"`
	ExtraHTTPHeaders              []OutputDatasetExtraHTTPHeaders      `tfsdk:"extra_http_headers"`
	FailedRequestLoggingMode      types.String                         `tfsdk:"failed_request_logging_mode"`
	FlushPeriodSec                types.Float64                        `tfsdk:"flush_period_sec"`
	ID                            types.String                         `tfsdk:"id"`
	MaxPayloadEvents              types.Float64                        `tfsdk:"max_payload_events"`
	MaxPayloadSizeKB              types.Float64                        `tfsdk:"max_payload_size_kb"`
	MessageField                  types.String                         `tfsdk:"message_field"`
	OnBackpressure                types.String                         `tfsdk:"on_backpressure"`
	Pipeline                      types.String                         `tfsdk:"pipeline"`
	PqCompress                    types.String                         `tfsdk:"pq_compress"`
	PqControls                    *OutputDatasetPqControls             `tfsdk:"pq_controls"`
	PqMaxFileSize                 types.String                         `tfsdk:"pq_max_file_size"`
	PqMaxSize                     types.String                         `tfsdk:"pq_max_size"`
	PqMode                        types.String                         `tfsdk:"pq_mode"`
	PqOnBackpressure              types.String                         `tfsdk:"pq_on_backpressure"`
	PqPath                        types.String                         `tfsdk:"pq_path"`
	RejectUnauthorized            types.Bool                           `tfsdk:"reject_unauthorized"`
	ResponseHonorRetryAfterHeader types.Bool                           `tfsdk:"response_honor_retry_after_header"`
	ResponseRetrySettings         []OutputDatasetResponseRetrySettings `tfsdk:"response_retry_settings"`
	SafeHeaders                   []types.String                       `tfsdk:"safe_headers"`
	ServerHostField               types.String                         `tfsdk:"server_host_field"`
	Site                          types.String                         `tfsdk:"site"`
	Streamtags                    []types.String                       `tfsdk:"streamtags"`
	SystemFields                  []types.String                       `tfsdk:"system_fields"`
	TextSecret                    types.String                         `tfsdk:"text_secret"`
	TimeoutRetrySettings          *OutputDatasetTimeoutRetrySettings   `tfsdk:"timeout_retry_settings"`
	TimeoutSec                    types.Float64                        `tfsdk:"timeout_sec"`
	TimestampField                types.String                         `tfsdk:"timestamp_field"`
	TotalMemoryLimitKB            types.Float64                        `tfsdk:"total_memory_limit_kb"`
	Type                          types.String                         `tfsdk:"type"`
	UseRoundRobinDNS              types.Bool                           `tfsdk:"use_round_robin_dns"`
}
