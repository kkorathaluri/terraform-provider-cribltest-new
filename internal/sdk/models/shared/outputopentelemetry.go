// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputOpenTelemetryType string

const (
	OutputOpenTelemetryTypeOpenTelemetry OutputOpenTelemetryType = "open_telemetry"
)

func (e OutputOpenTelemetryType) ToPointer() *OutputOpenTelemetryType {
	return &e
}
func (e *OutputOpenTelemetryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open_telemetry":
		*e = OutputOpenTelemetryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryType: %v", v)
	}
}

// OutputOpenTelemetryProtocol - Select a transport option for OpenTelemetry
type OutputOpenTelemetryProtocol string

const (
	OutputOpenTelemetryProtocolGrpc OutputOpenTelemetryProtocol = "grpc"
	OutputOpenTelemetryProtocolHTTP OutputOpenTelemetryProtocol = "http"
)

func (e OutputOpenTelemetryProtocol) ToPointer() *OutputOpenTelemetryProtocol {
	return &e
}
func (e *OutputOpenTelemetryProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "http":
		*e = OutputOpenTelemetryProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryProtocol: %v", v)
	}
}

// OutputOpenTelemetryOTLPVersion - The version of OTLP Protobuf definitions to use when structuring data to send
type OutputOpenTelemetryOTLPVersion string

const (
	OutputOpenTelemetryOTLPVersionZeroDot10Dot0 OutputOpenTelemetryOTLPVersion = "0.10.0"
	OutputOpenTelemetryOTLPVersionOneDot3Dot1   OutputOpenTelemetryOTLPVersion = "1.3.1"
)

func (e OutputOpenTelemetryOTLPVersion) ToPointer() *OutputOpenTelemetryOTLPVersion {
	return &e
}
func (e *OutputOpenTelemetryOTLPVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "0.10.0":
		fallthrough
	case "1.3.1":
		*e = OutputOpenTelemetryOTLPVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryOTLPVersion: %v", v)
	}
}

// OutputOpenTelemetryCompression - Type of compression to apply to messages sent to the OpenTelemetry endpoint
type OutputOpenTelemetryCompression string

const (
	OutputOpenTelemetryCompressionNone    OutputOpenTelemetryCompression = "none"
	OutputOpenTelemetryCompressionDeflate OutputOpenTelemetryCompression = "deflate"
	OutputOpenTelemetryCompressionGzip    OutputOpenTelemetryCompression = "gzip"
)

func (e OutputOpenTelemetryCompression) ToPointer() *OutputOpenTelemetryCompression {
	return &e
}
func (e *OutputOpenTelemetryCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "deflate":
		fallthrough
	case "gzip":
		*e = OutputOpenTelemetryCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryCompression: %v", v)
	}
}

// OutputOpenTelemetryHTTPCompressCompression - Type of compression to apply to messages sent to the OpenTelemetry endpoint
type OutputOpenTelemetryHTTPCompressCompression string

const (
	OutputOpenTelemetryHTTPCompressCompressionNone OutputOpenTelemetryHTTPCompressCompression = "none"
	OutputOpenTelemetryHTTPCompressCompressionGzip OutputOpenTelemetryHTTPCompressCompression = "gzip"
)

func (e OutputOpenTelemetryHTTPCompressCompression) ToPointer() *OutputOpenTelemetryHTTPCompressCompression {
	return &e
}
func (e *OutputOpenTelemetryHTTPCompressCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputOpenTelemetryHTTPCompressCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryHTTPCompressCompression: %v", v)
	}
}

// OutputOpenTelemetryAuthenticationType - OpenTelemetry authentication type
type OutputOpenTelemetryAuthenticationType string

const (
	OutputOpenTelemetryAuthenticationTypeNone              OutputOpenTelemetryAuthenticationType = "none"
	OutputOpenTelemetryAuthenticationTypeBasic             OutputOpenTelemetryAuthenticationType = "basic"
	OutputOpenTelemetryAuthenticationTypeCredentialsSecret OutputOpenTelemetryAuthenticationType = "credentialsSecret"
	OutputOpenTelemetryAuthenticationTypeToken             OutputOpenTelemetryAuthenticationType = "token"
	OutputOpenTelemetryAuthenticationTypeTextSecret        OutputOpenTelemetryAuthenticationType = "textSecret"
	OutputOpenTelemetryAuthenticationTypeOauth             OutputOpenTelemetryAuthenticationType = "oauth"
)

func (e OutputOpenTelemetryAuthenticationType) ToPointer() *OutputOpenTelemetryAuthenticationType {
	return &e
}
func (e *OutputOpenTelemetryAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "basic":
		fallthrough
	case "credentialsSecret":
		fallthrough
	case "token":
		fallthrough
	case "textSecret":
		fallthrough
	case "oauth":
		*e = OutputOpenTelemetryAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryAuthenticationType: %v", v)
	}
}

type OutputOpenTelemetryMetadata struct {
	Key   *string `default:"" json:"key"`
	Value string  `json:"value"`
}

func (o OutputOpenTelemetryMetadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputOpenTelemetryMetadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputOpenTelemetryMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *OutputOpenTelemetryMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputOpenTelemetryFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputOpenTelemetryFailedRequestLoggingMode string

const (
	OutputOpenTelemetryFailedRequestLoggingModePayload           OutputOpenTelemetryFailedRequestLoggingMode = "payload"
	OutputOpenTelemetryFailedRequestLoggingModePayloadAndHeaders OutputOpenTelemetryFailedRequestLoggingMode = "payloadAndHeaders"
	OutputOpenTelemetryFailedRequestLoggingModeNone              OutputOpenTelemetryFailedRequestLoggingMode = "none"
)

func (e OutputOpenTelemetryFailedRequestLoggingMode) ToPointer() *OutputOpenTelemetryFailedRequestLoggingMode {
	return &e
}
func (e *OutputOpenTelemetryFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payload":
		fallthrough
	case "payloadAndHeaders":
		fallthrough
	case "none":
		*e = OutputOpenTelemetryFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryFailedRequestLoggingMode: %v", v)
	}
}

// OutputOpenTelemetryBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputOpenTelemetryBackpressureBehavior string

const (
	OutputOpenTelemetryBackpressureBehaviorBlock OutputOpenTelemetryBackpressureBehavior = "block"
	OutputOpenTelemetryBackpressureBehaviorDrop  OutputOpenTelemetryBackpressureBehavior = "drop"
	OutputOpenTelemetryBackpressureBehaviorQueue OutputOpenTelemetryBackpressureBehavior = "queue"
)

func (e OutputOpenTelemetryBackpressureBehavior) ToPointer() *OutputOpenTelemetryBackpressureBehavior {
	return &e
}
func (e *OutputOpenTelemetryBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		fallthrough
	case "queue":
		*e = OutputOpenTelemetryBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryBackpressureBehavior: %v", v)
	}
}

type OutputOpenTelemetryOauthParams struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

func (o *OutputOpenTelemetryOauthParams) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OutputOpenTelemetryOauthParams) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputOpenTelemetryOauthHeaders struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

func (o *OutputOpenTelemetryOauthHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OutputOpenTelemetryOauthHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputOpenTelemetryExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputOpenTelemetryExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputOpenTelemetryExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputOpenTelemetryResponseRetrySettings struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputOpenTelemetryResponseRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputOpenTelemetryResponseRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputOpenTelemetryResponseRetrySettings) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputOpenTelemetryResponseRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputOpenTelemetryResponseRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputOpenTelemetryResponseRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputOpenTelemetryTimeoutRetrySettings struct {
	// Enable to retry on request timeout
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputOpenTelemetryTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputOpenTelemetryTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputOpenTelemetryTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputOpenTelemetryTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputOpenTelemetryTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputOpenTelemetryTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputOpenTelemetryMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputOpenTelemetryMinimumTLSVersion string

const (
	OutputOpenTelemetryMinimumTLSVersionTlSv1  OutputOpenTelemetryMinimumTLSVersion = "TLSv1"
	OutputOpenTelemetryMinimumTLSVersionTlSv11 OutputOpenTelemetryMinimumTLSVersion = "TLSv1.1"
	OutputOpenTelemetryMinimumTLSVersionTlSv12 OutputOpenTelemetryMinimumTLSVersion = "TLSv1.2"
	OutputOpenTelemetryMinimumTLSVersionTlSv13 OutputOpenTelemetryMinimumTLSVersion = "TLSv1.3"
)

func (e OutputOpenTelemetryMinimumTLSVersion) ToPointer() *OutputOpenTelemetryMinimumTLSVersion {
	return &e
}
func (e *OutputOpenTelemetryMinimumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = OutputOpenTelemetryMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryMinimumTLSVersion: %v", v)
	}
}

// OutputOpenTelemetryMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputOpenTelemetryMaximumTLSVersion string

const (
	OutputOpenTelemetryMaximumTLSVersionTlSv1  OutputOpenTelemetryMaximumTLSVersion = "TLSv1"
	OutputOpenTelemetryMaximumTLSVersionTlSv11 OutputOpenTelemetryMaximumTLSVersion = "TLSv1.1"
	OutputOpenTelemetryMaximumTLSVersionTlSv12 OutputOpenTelemetryMaximumTLSVersion = "TLSv1.2"
	OutputOpenTelemetryMaximumTLSVersionTlSv13 OutputOpenTelemetryMaximumTLSVersion = "TLSv1.3"
)

func (e OutputOpenTelemetryMaximumTLSVersion) ToPointer() *OutputOpenTelemetryMaximumTLSVersion {
	return &e
}
func (e *OutputOpenTelemetryMaximumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = OutputOpenTelemetryMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryMaximumTLSVersion: %v", v)
	}
}

type OutputOpenTelemetryTLSSettingsClientSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another
	//                     trusted CA (e.g., the system's CA). Defaults to Yes. Overrides the toggle from Advanced Settings, when also present.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputOpenTelemetryMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputOpenTelemetryMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (o OutputOpenTelemetryTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetMinVersion() *OutputOpenTelemetryMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *OutputOpenTelemetryTLSSettingsClientSide) GetMaxVersion() *OutputOpenTelemetryMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

// OutputOpenTelemetryPqCompressCompression - Codec to use to compress the persisted data.
type OutputOpenTelemetryPqCompressCompression string

const (
	OutputOpenTelemetryPqCompressCompressionNone OutputOpenTelemetryPqCompressCompression = "none"
	OutputOpenTelemetryPqCompressCompressionGzip OutputOpenTelemetryPqCompressCompression = "gzip"
)

func (e OutputOpenTelemetryPqCompressCompression) ToPointer() *OutputOpenTelemetryPqCompressCompression {
	return &e
}
func (e *OutputOpenTelemetryPqCompressCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputOpenTelemetryPqCompressCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryPqCompressCompression: %v", v)
	}
}

// OutputOpenTelemetryQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputOpenTelemetryQueueFullBehavior string

const (
	OutputOpenTelemetryQueueFullBehaviorBlock OutputOpenTelemetryQueueFullBehavior = "block"
	OutputOpenTelemetryQueueFullBehaviorDrop  OutputOpenTelemetryQueueFullBehavior = "drop"
)

func (e OutputOpenTelemetryQueueFullBehavior) ToPointer() *OutputOpenTelemetryQueueFullBehavior {
	return &e
}
func (e *OutputOpenTelemetryQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputOpenTelemetryQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryQueueFullBehavior: %v", v)
	}
}

// OutputOpenTelemetryMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputOpenTelemetryMode string

const (
	OutputOpenTelemetryModeError        OutputOpenTelemetryMode = "error"
	OutputOpenTelemetryModeBackpressure OutputOpenTelemetryMode = "backpressure"
	OutputOpenTelemetryModeAlways       OutputOpenTelemetryMode = "always"
)

func (e OutputOpenTelemetryMode) ToPointer() *OutputOpenTelemetryMode {
	return &e
}
func (e *OutputOpenTelemetryMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "backpressure":
		fallthrough
	case "always":
		*e = OutputOpenTelemetryMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputOpenTelemetryMode: %v", v)
	}
}

type OutputOpenTelemetryPqControls struct {
}

type OutputOpenTelemetry struct {
	// Unique ID for this output
	ID   *string                 `json:"id,omitempty"`
	Type OutputOpenTelemetryType `json:"type"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Select a transport option for OpenTelemetry
	Protocol *OutputOpenTelemetryProtocol `default:"grpc" json:"protocol"`
	// The endpoint where OTel events will be sent. Enter any valid URL or an IP address (IPv4 or IPv6; enclose IPv6 addresses in square brackets). Unspecified ports will default to 4317, unless the endpoint is an HTTPS-based URL or TLS is enabled, in which case 443 will be used.
	Endpoint string `json:"endpoint"`
	// The version of OTLP Protobuf definitions to use when structuring data to send
	OtlpVersion *OutputOpenTelemetryOTLPVersion `default:"0.10.0" json:"otlpVersion"`
	// Type of compression to apply to messages sent to the OpenTelemetry endpoint
	Compress *OutputOpenTelemetryCompression `default:"gzip" json:"compress"`
	// Type of compression to apply to messages sent to the OpenTelemetry endpoint
	HTTPCompress *OutputOpenTelemetryHTTPCompressCompression `default:"gzip" json:"httpCompress"`
	// OpenTelemetry authentication type
	AuthType *OutputOpenTelemetryAuthenticationType `default:"none" json:"authType"`
	// If you want to send traces to the default `{endpoint}/v1/traces` endpoint, leave this field empty; otherwise, specify the desired endpoint
	HTTPTracesEndpointOverride *string `json:"httpTracesEndpointOverride,omitempty"`
	// If you want to send metrics to the default `{endpoint}/v1/metrics` endpoint, leave this field empty; otherwise, specify the desired endpoint
	HTTPMetricsEndpointOverride *string `json:"httpMetricsEndpointOverride,omitempty"`
	// If you want to send logs to the default `{endpoint}/v1/logs` endpoint, leave this field empty; otherwise, specify the desired endpoint
	HTTPLogsEndpointOverride *string `json:"httpLogsEndpointOverride,omitempty"`
	// List of key-value pairs to send with each gRPC request. Value supports JavaScript expressions that are evaluated just once, when the destination gets started. To pass credentials as metadata, use 'C.Secret'.
	Metadata []OutputOpenTelemetryMetadata `json:"metadata,omitempty"`
	// Maximum number of ongoing requests before blocking
	Concurrency *float64 `default:"5" json:"concurrency"`
	// Maximum size, in KB, of the request body
	MaxPayloadSizeKB *float64 `default:"4096" json:"maxPayloadSizeKB"`
	// Amount of time, in seconds, to wait for a request to complete before canceling it
	TimeoutSec *float64 `default:"30" json:"timeoutSec"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *float64 `default:"1" json:"flushPeriodSec"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputOpenTelemetryFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// Amount of time (milliseconds) to wait for the connection to establish before retrying
	ConnectionTimeout *float64 `default:"10000" json:"connectionTimeout"`
	// How often the sender should ping the peer to keep the connection open
	KeepAliveTime *float64 `default:"30" json:"keepAliveTime"`
	// Disable to close the connection immediately after sending the outgoing request
	KeepAlive *bool `default:"true" json:"keepAlive"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputOpenTelemetryBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                                  `json:"description,omitempty"`
	Username       *string                                  `json:"username,omitempty"`
	Password       *string                                  `json:"password,omitempty"`
	// Bearer token to include in the authorization header
	Token *string `json:"token,omitempty"`
	// Select or create a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Select or create a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// URL for OAuth
	LoginURL *string `json:"loginUrl,omitempty"`
	// Secret parameter name to pass in request body
	SecretParamName *string `json:"secretParamName,omitempty"`
	// Secret parameter value to pass in request body
	Secret *string `json:"secret,omitempty"`
	// Name of the auth token attribute in the OAuth response. Can be top-level (e.g., 'token'); or nested, using a period (e.g., 'data.token').
	TokenAttributeName *string `json:"tokenAttributeName,omitempty"`
	// JavaScript expression to compute the Authorization header value to pass in requests. The value `${token}` is used to reference the token obtained from authentication, e.g.: `Bearer ${token}`.
	AuthHeaderExpr *string `default:"Bearer \\${token}" json:"authHeaderExpr"`
	// How often the OAuth token should be refreshed.
	TokenTimeoutSecs *float64 `default:"3600" json:"tokenTimeoutSecs"`
	// Additional parameters to send in the OAuth login request. @{product} will combine the secret with these parameters, and will send the URL-encoded result in a POST request to the endpoint specified in the 'Login URL'. We'll automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthParams []OutputOpenTelemetryOauthParams `json:"oauthParams,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []OutputOpenTelemetryOauthHeaders `json:"oauthHeaders,omitempty"`
	// Reject certificates not authorized by a CA in the CA certificate path or by another trusted CA (such as the system's).
	//         Defaults to Yes. When this setting is also present in TLS Settings (Client Side),
	//         that value will take precedence.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Enables round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"false" json:"useRoundRobinDns"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputOpenTelemetryExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable).
	ResponseRetrySettings []OutputOpenTelemetryResponseRetrySettings `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputOpenTelemetryTimeoutRetrySettings   `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool                                     `default:"false" json:"responseHonorRetryAfterHeader"`
	TLS                           *OutputOpenTelemetryTLSSettingsClientSide `json:"tls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"\\$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputOpenTelemetryPqCompressCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputOpenTelemetryQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputOpenTelemetryMode       `default:"error" json:"pqMode"`
	PqControls *OutputOpenTelemetryPqControls `json:"pqControls,omitempty"`
}

func (o OutputOpenTelemetry) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputOpenTelemetry) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputOpenTelemetry) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputOpenTelemetry) GetType() OutputOpenTelemetryType {
	if o == nil {
		return OutputOpenTelemetryType("")
	}
	return o.Type
}

func (o *OutputOpenTelemetry) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputOpenTelemetry) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputOpenTelemetry) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputOpenTelemetry) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputOpenTelemetry) GetProtocol() *OutputOpenTelemetryProtocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *OutputOpenTelemetry) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *OutputOpenTelemetry) GetOtlpVersion() *OutputOpenTelemetryOTLPVersion {
	if o == nil {
		return nil
	}
	return o.OtlpVersion
}

func (o *OutputOpenTelemetry) GetCompress() *OutputOpenTelemetryCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputOpenTelemetry) GetHTTPCompress() *OutputOpenTelemetryHTTPCompressCompression {
	if o == nil {
		return nil
	}
	return o.HTTPCompress
}

func (o *OutputOpenTelemetry) GetAuthType() *OutputOpenTelemetryAuthenticationType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputOpenTelemetry) GetHTTPTracesEndpointOverride() *string {
	if o == nil {
		return nil
	}
	return o.HTTPTracesEndpointOverride
}

func (o *OutputOpenTelemetry) GetHTTPMetricsEndpointOverride() *string {
	if o == nil {
		return nil
	}
	return o.HTTPMetricsEndpointOverride
}

func (o *OutputOpenTelemetry) GetHTTPLogsEndpointOverride() *string {
	if o == nil {
		return nil
	}
	return o.HTTPLogsEndpointOverride
}

func (o *OutputOpenTelemetry) GetMetadata() []OutputOpenTelemetryMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *OutputOpenTelemetry) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputOpenTelemetry) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputOpenTelemetry) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputOpenTelemetry) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputOpenTelemetry) GetFailedRequestLoggingMode() *OutputOpenTelemetryFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputOpenTelemetry) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *OutputOpenTelemetry) GetKeepAliveTime() *float64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTime
}

func (o *OutputOpenTelemetry) GetKeepAlive() *bool {
	if o == nil {
		return nil
	}
	return o.KeepAlive
}

func (o *OutputOpenTelemetry) GetOnBackpressure() *OutputOpenTelemetryBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputOpenTelemetry) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputOpenTelemetry) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *OutputOpenTelemetry) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *OutputOpenTelemetry) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OutputOpenTelemetry) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

func (o *OutputOpenTelemetry) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputOpenTelemetry) GetLoginURL() *string {
	if o == nil {
		return nil
	}
	return o.LoginURL
}

func (o *OutputOpenTelemetry) GetSecretParamName() *string {
	if o == nil {
		return nil
	}
	return o.SecretParamName
}

func (o *OutputOpenTelemetry) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *OutputOpenTelemetry) GetTokenAttributeName() *string {
	if o == nil {
		return nil
	}
	return o.TokenAttributeName
}

func (o *OutputOpenTelemetry) GetAuthHeaderExpr() *string {
	if o == nil {
		return nil
	}
	return o.AuthHeaderExpr
}

func (o *OutputOpenTelemetry) GetTokenTimeoutSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TokenTimeoutSecs
}

func (o *OutputOpenTelemetry) GetOauthParams() []OutputOpenTelemetryOauthParams {
	if o == nil {
		return nil
	}
	return o.OauthParams
}

func (o *OutputOpenTelemetry) GetOauthHeaders() []OutputOpenTelemetryOauthHeaders {
	if o == nil {
		return nil
	}
	return o.OauthHeaders
}

func (o *OutputOpenTelemetry) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputOpenTelemetry) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputOpenTelemetry) GetExtraHTTPHeaders() []OutputOpenTelemetryExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputOpenTelemetry) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputOpenTelemetry) GetResponseRetrySettings() []OutputOpenTelemetryResponseRetrySettings {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputOpenTelemetry) GetTimeoutRetrySettings() *OutputOpenTelemetryTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputOpenTelemetry) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputOpenTelemetry) GetTLS() *OutputOpenTelemetryTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *OutputOpenTelemetry) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputOpenTelemetry) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputOpenTelemetry) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputOpenTelemetry) GetPqCompress() *OutputOpenTelemetryPqCompressCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputOpenTelemetry) GetPqOnBackpressure() *OutputOpenTelemetryQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputOpenTelemetry) GetPqMode() *OutputOpenTelemetryMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputOpenTelemetry) GetPqControls() *OutputOpenTelemetryPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
