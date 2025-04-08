// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputCriblHTTPType string

const (
	InputCriblHTTPTypeCriblHTTP InputCriblHTTPType = "cribl_http"
)

func (e InputCriblHTTPType) ToPointer() *InputCriblHTTPType {
	return &e
}
func (e *InputCriblHTTPType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cribl_http":
		*e = InputCriblHTTPType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblHTTPType: %v", v)
	}
}

type InputCriblHTTPConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputCriblHTTPConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputCriblHTTPConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputCriblHTTPMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputCriblHTTPMode string

const (
	InputCriblHTTPModeSmart  InputCriblHTTPMode = "smart"
	InputCriblHTTPModeAlways InputCriblHTTPMode = "always"
)

func (e InputCriblHTTPMode) ToPointer() *InputCriblHTTPMode {
	return &e
}
func (e *InputCriblHTTPMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputCriblHTTPMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblHTTPMode: %v", v)
	}
}

// InputCriblHTTPCompression - Codec to use to compress the persisted data
type InputCriblHTTPCompression string

const (
	InputCriblHTTPCompressionNone InputCriblHTTPCompression = "none"
	InputCriblHTTPCompressionGzip InputCriblHTTPCompression = "gzip"
)

func (e InputCriblHTTPCompression) ToPointer() *InputCriblHTTPCompression {
	return &e
}
func (e *InputCriblHTTPCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputCriblHTTPCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblHTTPCompression: %v", v)
	}
}

type InputCriblHTTPPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputCriblHTTPMode `default:"always" json:"mode"`
	// The maximum number of events to hold in memory before writing the events to disk
	MaxBufferSize *float64 `default:"1000" json:"maxBufferSize"`
	// The number of events to send downstream before committing that Stream has read them
	CommitFrequency *float64 `default:"42" json:"commitFrequency"`
	// The maximum size to store in each queue file before closing and optionally compressing. Enter a numeral with units of KB, MB, etc.
	MaxFileSize *string `default:"1 MB" json:"maxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `default:"5GB" json:"maxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>
	Path *string `default:"\\$CRIBL_HOME/state/queues" json:"path"`
	// Codec to use to compress the persisted data
	Compress *InputCriblHTTPCompression `default:"none" json:"compress"`
}

func (i InputCriblHTTPPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputCriblHTTPPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputCriblHTTPPq) GetMode() *InputCriblHTTPMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputCriblHTTPPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputCriblHTTPPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputCriblHTTPPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputCriblHTTPPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputCriblHTTPPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputCriblHTTPPq) GetCompress() *InputCriblHTTPCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputCriblHTTPMinimumTLSVersion - Minimum TLS version to accept from connections
type InputCriblHTTPMinimumTLSVersion string

const (
	InputCriblHTTPMinimumTLSVersionTlSv1  InputCriblHTTPMinimumTLSVersion = "TLSv1"
	InputCriblHTTPMinimumTLSVersionTlSv11 InputCriblHTTPMinimumTLSVersion = "TLSv1.1"
	InputCriblHTTPMinimumTLSVersionTlSv12 InputCriblHTTPMinimumTLSVersion = "TLSv1.2"
	InputCriblHTTPMinimumTLSVersionTlSv13 InputCriblHTTPMinimumTLSVersion = "TLSv1.3"
)

func (e InputCriblHTTPMinimumTLSVersion) ToPointer() *InputCriblHTTPMinimumTLSVersion {
	return &e
}
func (e *InputCriblHTTPMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputCriblHTTPMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblHTTPMinimumTLSVersion: %v", v)
	}
}

// InputCriblHTTPMaximumTLSVersion - Maximum TLS version to accept from connections
type InputCriblHTTPMaximumTLSVersion string

const (
	InputCriblHTTPMaximumTLSVersionTlSv1  InputCriblHTTPMaximumTLSVersion = "TLSv1"
	InputCriblHTTPMaximumTLSVersionTlSv11 InputCriblHTTPMaximumTLSVersion = "TLSv1.1"
	InputCriblHTTPMaximumTLSVersionTlSv12 InputCriblHTTPMaximumTLSVersion = "TLSv1.2"
	InputCriblHTTPMaximumTLSVersionTlSv13 InputCriblHTTPMaximumTLSVersion = "TLSv1.3"
)

func (e InputCriblHTTPMaximumTLSVersion) ToPointer() *InputCriblHTTPMaximumTLSVersion {
	return &e
}
func (e *InputCriblHTTPMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputCriblHTTPMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblHTTPMaximumTLSVersion: %v", v)
	}
}

type InputCriblHTTPTLSSettingsServerSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// The name of the predefined certificate
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Passphrase to use to decrypt private key
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert        *bool `default:"false" json:"requestCert"`
	RejectUnauthorized any   `json:"rejectUnauthorized,omitempty"`
	CommonNameRegex    any   `json:"commonNameRegex,omitempty"`
	// Minimum TLS version to accept from connections
	MinVersion *InputCriblHTTPMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to accept from connections
	MaxVersion *InputCriblHTTPMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (i InputCriblHTTPTLSSettingsServerSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputCriblHTTPTLSSettingsServerSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetRequestCert() *bool {
	if o == nil {
		return nil
	}
	return o.RequestCert
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetRejectUnauthorized() any {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetCommonNameRegex() any {
	if o == nil {
		return nil
	}
	return o.CommonNameRegex
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetMinVersion() *InputCriblHTTPMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputCriblHTTPTLSSettingsServerSide) GetMaxVersion() *InputCriblHTTPMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type InputCriblHTTPMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputCriblHTTPMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputCriblHTTPMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputCriblHTTP struct {
	// Unique ID for this input
	ID       *string             `json:"id,omitempty"`
	Type     *InputCriblHTTPType `json:"type,omitempty"`
	Disabled *bool               `default:"false" json:"disabled"`
	// Pipeline to process data from this Source before sending it through the Routes
	Pipeline *string `json:"pipeline,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `default:"true" json:"sendToRoutes"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Use a disk queue to minimize data loss when connected services block. See [Cribl Docs](https://docs.cribl.io/stream/persistent-queues) for PQ defaults (Cribl-managed Cloud Workers) and configuration options (on-prem and hybrid Workers).
	PqEnabled *bool `default:"false" json:"pqEnabled"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Direct connections to Destinations, and optionally via a Pipeline or a Pack
	Connections []InputCriblHTTPConnections `json:"connections,omitempty"`
	Pq          *InputCriblHTTPPq           `json:"pq,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host *string `default:"0.0.0.0" json:"host"`
	// Port to listen on
	Port float64 `json:"port"`
	// Shared secrets to be provided by any client (Authorization: <token>). If empty, unauthorized access is permitted.
	AuthTokens []string                             `json:"authTokens,omitempty"`
	TLS        *InputCriblHTTPTLSSettingsServerSide `json:"tls,omitempty"`
	// Maximum number of active requests per Worker Process. Use 0 for unlimited.
	MaxActiveReq *float64 `default:"256" json:"maxActiveReq"`
	// Maximum number of requests per socket before @{product} instructs the client to close the connection. Default is 0 (unlimited).
	MaxRequestsPerSocket *int64 `default:"0" json:"maxRequestsPerSocket"`
	// Enable when clients are connecting through a proxy that supports the x-forwarded-for header to keep the client's original IP address on the event instead of the proxy's IP address
	EnableProxyHeader *bool `default:"false" json:"enableProxyHeader"`
	// Add request headers to events, in the __headers field
	CaptureHeaders *bool `default:"false" json:"captureHeaders"`
	// How often request activity is logged at the `info` level. A value of 1 would log every request, 10 every 10th request, etc.
	ActivityLogSampleRate *float64 `default:"100" json:"activityLogSampleRate"`
	// How long to wait for an incoming request to complete before aborting it. Use 0 to disable.
	RequestTimeout *float64 `default:"0" json:"requestTimeout"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. To wait forever, set to 0.
	SocketTimeout *float64 `default:"0" json:"socketTimeout"`
	// After the last response is sent, @{product} will wait this long for additional data before closing the socket connection. Minimum 1 sec.; maximum 600 sec. (10 min.).
	KeepAliveTimeout *float64 `default:"5" json:"keepAliveTimeout"`
	// Enable to expose the /cribl_health endpoint, which returns 200 OK when this Source is healthy
	EnableHealthCheck *bool `default:"false" json:"enableHealthCheck"`
	// Messages from matched IP addresses will be processed, unless also matched by the denylist
	IPAllowlistRegex *string `default:"/.*/" json:"ipAllowlistRegex"`
	// Messages from matched IP addresses will be ignored. This takes precedence over the allowlist.
	IPDenylistRegex *string `default:"/^\\$/" json:"ipDenylistRegex"`
	// Fields to add to events from this input
	Metadata    []InputCriblHTTPMetadata `json:"metadata,omitempty"`
	Description *string                  `json:"description,omitempty"`
}

func (i InputCriblHTTP) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputCriblHTTP) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputCriblHTTP) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputCriblHTTP) GetType() *InputCriblHTTPType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputCriblHTTP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputCriblHTTP) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputCriblHTTP) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputCriblHTTP) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputCriblHTTP) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputCriblHTTP) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputCriblHTTP) GetConnections() []InputCriblHTTPConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputCriblHTTP) GetPq() *InputCriblHTTPPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputCriblHTTP) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputCriblHTTP) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *InputCriblHTTP) GetAuthTokens() []string {
	if o == nil {
		return nil
	}
	return o.AuthTokens
}

func (o *InputCriblHTTP) GetTLS() *InputCriblHTTPTLSSettingsServerSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputCriblHTTP) GetMaxActiveReq() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxActiveReq
}

func (o *InputCriblHTTP) GetMaxRequestsPerSocket() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestsPerSocket
}

func (o *InputCriblHTTP) GetEnableProxyHeader() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxyHeader
}

func (o *InputCriblHTTP) GetCaptureHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.CaptureHeaders
}

func (o *InputCriblHTTP) GetActivityLogSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.ActivityLogSampleRate
}

func (o *InputCriblHTTP) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputCriblHTTP) GetSocketTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketTimeout
}

func (o *InputCriblHTTP) GetKeepAliveTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTimeout
}

func (o *InputCriblHTTP) GetEnableHealthCheck() *bool {
	if o == nil {
		return nil
	}
	return o.EnableHealthCheck
}

func (o *InputCriblHTTP) GetIPAllowlistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPAllowlistRegex
}

func (o *InputCriblHTTP) GetIPDenylistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPDenylistRegex
}

func (o *InputCriblHTTP) GetMetadata() []InputCriblHTTPMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputCriblHTTP) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
