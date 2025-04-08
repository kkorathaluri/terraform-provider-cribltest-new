// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputAzureDataExplorer struct {
	AddIDToStagePath              types.Bool                                     `tfsdk:"add_id_to_stage_path"`
	AdditionalProperties          []AdditionalProperties                         `tfsdk:"additional_properties"`
	Certificate                   *OutputAzureDataExplorerCertificate            `tfsdk:"certificate"`
	ClientID                      types.String                                   `tfsdk:"client_id"`
	ClientSecret                  types.String                                   `tfsdk:"client_secret"`
	ClusterURL                    types.String                                   `tfsdk:"cluster_url"`
	Compress                      types.String                                   `tfsdk:"compress"`
	Concurrency                   types.Float64                                  `tfsdk:"concurrency"`
	Database                      types.String                                   `tfsdk:"database"`
	DeadletterEnabled             types.Bool                                     `tfsdk:"deadletter_enabled"`
	Description                   types.String                                   `tfsdk:"description"`
	Environment                   types.String                                   `tfsdk:"environment"`
	ExtentTags                    []ExtentTags                                   `tfsdk:"extent_tags"`
	FileNameSuffix                types.String                                   `tfsdk:"file_name_suffix"`
	FlushImmediately              types.Bool                                     `tfsdk:"flush_immediately"`
	FlushPeriodSec                types.Float64                                  `tfsdk:"flush_period_sec"`
	Format                        types.String                                   `tfsdk:"format"`
	ID                            types.String                                   `tfsdk:"id"`
	IngestIfNotExists             []IngestIfNotExists                            `tfsdk:"ingest_if_not_exists"`
	IngestMode                    types.String                                   `tfsdk:"ingest_mode"`
	IngestURL                     types.String                                   `tfsdk:"ingest_url"`
	IsMappingObj                  types.Bool                                     `tfsdk:"is_mapping_obj"`
	KeepAlive                     types.Bool                                     `tfsdk:"keep_alive"`
	MappingRef                    types.String                                   `tfsdk:"mapping_ref"`
	MaxConcurrentFileParts        types.Float64                                  `tfsdk:"max_concurrent_file_parts"`
	MaxFileIdleTimeSec            types.Float64                                  `tfsdk:"max_file_idle_time_sec"`
	MaxFileOpenTimeSec            types.Float64                                  `tfsdk:"max_file_open_time_sec"`
	MaxFileSizeMB                 types.Float64                                  `tfsdk:"max_file_size_mb"`
	MaxOpenFiles                  types.Float64                                  `tfsdk:"max_open_files"`
	MaxPayloadEvents              types.Float64                                  `tfsdk:"max_payload_events"`
	MaxPayloadSizeKB              types.Float64                                  `tfsdk:"max_payload_size_kb"`
	OauthEndpoint                 types.String                                   `tfsdk:"oauth_endpoint"`
	OauthType                     types.String                                   `tfsdk:"oauth_type"`
	OnBackpressure                types.String                                   `tfsdk:"on_backpressure"`
	OnDiskFullBackpressure        types.String                                   `tfsdk:"on_disk_full_backpressure"`
	Pipeline                      types.String                                   `tfsdk:"pipeline"`
	PqCompress                    types.String                                   `tfsdk:"pq_compress"`
	PqControls                    *OutputAzureDataExplorerPqControls             `tfsdk:"pq_controls"`
	PqMaxFileSize                 types.String                                   `tfsdk:"pq_max_file_size"`
	PqMaxSize                     types.String                                   `tfsdk:"pq_max_size"`
	PqMode                        types.String                                   `tfsdk:"pq_mode"`
	PqOnBackpressure              types.String                                   `tfsdk:"pq_on_backpressure"`
	PqPath                        types.String                                   `tfsdk:"pq_path"`
	RejectUnauthorized            types.Bool                                     `tfsdk:"reject_unauthorized"`
	RemoveEmptyDirs               types.Bool                                     `tfsdk:"remove_empty_dirs"`
	ReportLevel                   types.String                                   `tfsdk:"report_level"`
	ReportMethod                  types.String                                   `tfsdk:"report_method"`
	ResponseHonorRetryAfterHeader types.Bool                                     `tfsdk:"response_honor_retry_after_header"`
	ResponseRetrySettings         []OutputAzureDataExplorerResponseRetrySettings `tfsdk:"response_retry_settings"`
	RetainBlobOnSuccess           types.Bool                                     `tfsdk:"retain_blob_on_success"`
	Scope                         types.String                                   `tfsdk:"scope"`
	StagePath                     types.String                                   `tfsdk:"stage_path"`
	Streamtags                    []types.String                                 `tfsdk:"streamtags"`
	SystemFields                  []types.String                                 `tfsdk:"system_fields"`
	Table                         types.String                                   `tfsdk:"table"`
	TenantID                      types.String                                   `tfsdk:"tenant_id"`
	TextSecret                    types.String                                   `tfsdk:"text_secret"`
	TimeoutRetrySettings          *OutputAzureDataExplorerTimeoutRetrySettings   `tfsdk:"timeout_retry_settings"`
	TimeoutSec                    types.Float64                                  `tfsdk:"timeout_sec"`
	Type                          types.String                                   `tfsdk:"type"`
	UseRoundRobinDNS              types.Bool                                     `tfsdk:"use_round_robin_dns"`
	ValidateDatabaseSettings      types.Bool                                     `tfsdk:"validate_database_settings"`
}
