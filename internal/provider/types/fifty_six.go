// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type FiftySix struct {
	AssumeRoleArn               types.String         `tfsdk:"assume_role_arn"`
	AssumeRoleExternalID        types.String         `tfsdk:"assume_role_external_id"`
	AwsAccountID                types.String         `tfsdk:"aws_account_id"`
	AwsAPIKey                   types.String         `tfsdk:"aws_api_key"`
	AwsAuthenticationMethod     types.String         `tfsdk:"aws_authentication_method"`
	AwsSecret                   types.String         `tfsdk:"aws_secret"`
	AwsSecretKey                types.String         `tfsdk:"aws_secret_key"`
	BreakerRulesets             []types.String       `tfsdk:"breaker_rulesets"`
	Checkpointing               *InputCheckpointing  `tfsdk:"checkpointing"`
	Connections                 []Input56Connections `tfsdk:"connections"`
	Description                 types.String         `tfsdk:"description"`
	Disabled                    types.Bool           `tfsdk:"disabled"`
	DurationSeconds             types.Number         `tfsdk:"duration_seconds"`
	EnableAssumeRole            types.Bool           `tfsdk:"enable_assume_role"`
	EnableSQSAssumeRole         types.Bool           `tfsdk:"enable_sqs_assume_role"`
	Encoding                    types.String         `tfsdk:"encoding"`
	Endpoint                    types.String         `tfsdk:"endpoint"`
	Environment                 types.String         `tfsdk:"environment"`
	FileFilter                  types.String         `tfsdk:"file_filter"`
	ID                          types.String         `tfsdk:"id"`
	MaxMessages                 types.Number         `tfsdk:"max_messages"`
	Metadata                    []Input56Metadata    `tfsdk:"metadata"`
	NumReceivers                types.Number         `tfsdk:"num_receivers"`
	ParquetChunkDownloadTimeout types.Number         `tfsdk:"parquet_chunk_download_timeout"`
	ParquetChunkSizeMB          types.Number         `tfsdk:"parquet_chunk_size_mb"`
	Pipeline                    types.String         `tfsdk:"pipeline"`
	PollTimeout                 types.Number         `tfsdk:"poll_timeout"`
	Pq                          *Input56Pq           `tfsdk:"pq"`
	PqEnabled                   types.Bool           `tfsdk:"pq_enabled"`
	Preprocess                  *Input56Preprocess   `tfsdk:"preprocess"`
	QueueName                   types.String         `tfsdk:"queue_name"`
	Region                      types.String         `tfsdk:"region"`
	RejectUnauthorized          types.Bool           `tfsdk:"reject_unauthorized"`
	ReuseConnections            types.Bool           `tfsdk:"reuse_connections"`
	SendToRoutes                types.Bool           `tfsdk:"send_to_routes"`
	SignatureVersion            types.String         `tfsdk:"signature_version"`
	SkipOnError                 types.Bool           `tfsdk:"skip_on_error"`
	SocketTimeout               types.Number         `tfsdk:"socket_timeout"`
	StaleChannelFlushMs         types.Number         `tfsdk:"stale_channel_flush_ms"`
	Status                      *Input56Status       `tfsdk:"status"`
	Streamtags                  []types.String       `tfsdk:"streamtags"`
	Type                        types.String         `tfsdk:"type"`
	VisibilityTimeout           types.Number         `tfsdk:"visibility_timeout"`
}
