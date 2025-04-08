// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputConfluentCloud struct {
	Ack                       types.Int64                                            `tfsdk:"ack"`
	AuthenticationTimeout     types.Float64                                          `tfsdk:"authentication_timeout"`
	BackoffRate               types.Float64                                          `tfsdk:"backoff_rate"`
	Brokers                   []types.String                                         `tfsdk:"brokers"`
	Compression               types.String                                           `tfsdk:"compression"`
	ConnectionTimeout         types.Float64                                          `tfsdk:"connection_timeout"`
	Description               types.String                                           `tfsdk:"description"`
	Environment               types.String                                           `tfsdk:"environment"`
	FlushEventCount           types.Float64                                          `tfsdk:"flush_event_count"`
	FlushPeriodSec            types.Float64                                          `tfsdk:"flush_period_sec"`
	Format                    types.String                                           `tfsdk:"format"`
	ID                        types.String                                           `tfsdk:"id"`
	InitialBackoff            types.Float64                                          `tfsdk:"initial_backoff"`
	KafkaSchemaRegistry       *OutputConfluentCloudKafkaSchemaRegistryAuthentication `tfsdk:"kafka_schema_registry"`
	MaxBackOff                types.Float64                                          `tfsdk:"max_back_off"`
	MaxRecordSizeKB           types.Float64                                          `tfsdk:"max_record_size_kb"`
	MaxRetries                types.Float64                                          `tfsdk:"max_retries"`
	OnBackpressure            types.String                                           `tfsdk:"on_backpressure"`
	Pipeline                  types.String                                           `tfsdk:"pipeline"`
	PqCompress                types.String                                           `tfsdk:"pq_compress"`
	PqControls                *OutputConfluentCloudPqControls                        `tfsdk:"pq_controls"`
	PqMaxFileSize             types.String                                           `tfsdk:"pq_max_file_size"`
	PqMaxSize                 types.String                                           `tfsdk:"pq_max_size"`
	PqMode                    types.String                                           `tfsdk:"pq_mode"`
	PqOnBackpressure          types.String                                           `tfsdk:"pq_on_backpressure"`
	PqPath                    types.String                                           `tfsdk:"pq_path"`
	ProtobufLibraryID         types.String                                           `tfsdk:"protobuf_library_id"`
	ReauthenticationThreshold types.Float64                                          `tfsdk:"reauthentication_threshold"`
	RequestTimeout            types.Float64                                          `tfsdk:"request_timeout"`
	Sasl                      *OutputConfluentCloudAuthentication                    `tfsdk:"sasl"`
	Streamtags                []types.String                                         `tfsdk:"streamtags"`
	SystemFields              []types.String                                         `tfsdk:"system_fields"`
	TLS                       *OutputConfluentCloudTLSSettingsClientSide             `tfsdk:"tls"`
	Topic                     types.String                                           `tfsdk:"topic"`
	Type                      types.String                                           `tfsdk:"type"`
}
