// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputConfluentCloudType string

const (
	InputConfluentCloudTypeConfluentCloud InputConfluentCloudType = "confluent_cloud"
)

func (e InputConfluentCloudType) ToPointer() *InputConfluentCloudType {
	return &e
}
func (e *InputConfluentCloudType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "confluent_cloud":
		*e = InputConfluentCloudType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudType: %v", v)
	}
}

type InputConfluentCloudConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputConfluentCloudConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputConfluentCloudConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputConfluentCloudMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputConfluentCloudMode string

const (
	InputConfluentCloudModeSmart  InputConfluentCloudMode = "smart"
	InputConfluentCloudModeAlways InputConfluentCloudMode = "always"
)

func (e InputConfluentCloudMode) ToPointer() *InputConfluentCloudMode {
	return &e
}
func (e *InputConfluentCloudMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputConfluentCloudMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudMode: %v", v)
	}
}

// InputConfluentCloudCompression - Codec to use to compress the persisted data
type InputConfluentCloudCompression string

const (
	InputConfluentCloudCompressionNone InputConfluentCloudCompression = "none"
	InputConfluentCloudCompressionGzip InputConfluentCloudCompression = "gzip"
)

func (e InputConfluentCloudCompression) ToPointer() *InputConfluentCloudCompression {
	return &e
}
func (e *InputConfluentCloudCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputConfluentCloudCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudCompression: %v", v)
	}
}

type InputConfluentCloudPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputConfluentCloudMode `default:"always" json:"mode"`
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
	Compress *InputConfluentCloudCompression `default:"none" json:"compress"`
}

func (i InputConfluentCloudPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputConfluentCloudPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputConfluentCloudPq) GetMode() *InputConfluentCloudMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputConfluentCloudPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputConfluentCloudPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputConfluentCloudPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputConfluentCloudPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputConfluentCloudPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputConfluentCloudPq) GetCompress() *InputConfluentCloudCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputConfluentCloudMinimumTLSVersion - Minimum TLS version to use when connecting
type InputConfluentCloudMinimumTLSVersion string

const (
	InputConfluentCloudMinimumTLSVersionTlSv1  InputConfluentCloudMinimumTLSVersion = "TLSv1"
	InputConfluentCloudMinimumTLSVersionTlSv11 InputConfluentCloudMinimumTLSVersion = "TLSv1.1"
	InputConfluentCloudMinimumTLSVersionTlSv12 InputConfluentCloudMinimumTLSVersion = "TLSv1.2"
	InputConfluentCloudMinimumTLSVersionTlSv13 InputConfluentCloudMinimumTLSVersion = "TLSv1.3"
)

func (e InputConfluentCloudMinimumTLSVersion) ToPointer() *InputConfluentCloudMinimumTLSVersion {
	return &e
}
func (e *InputConfluentCloudMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputConfluentCloudMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudMinimumTLSVersion: %v", v)
	}
}

// InputConfluentCloudMaximumTLSVersion - Maximum TLS version to use when connecting
type InputConfluentCloudMaximumTLSVersion string

const (
	InputConfluentCloudMaximumTLSVersionTlSv1  InputConfluentCloudMaximumTLSVersion = "TLSv1"
	InputConfluentCloudMaximumTLSVersionTlSv11 InputConfluentCloudMaximumTLSVersion = "TLSv1.1"
	InputConfluentCloudMaximumTLSVersionTlSv12 InputConfluentCloudMaximumTLSVersion = "TLSv1.2"
	InputConfluentCloudMaximumTLSVersionTlSv13 InputConfluentCloudMaximumTLSVersion = "TLSv1.3"
)

func (e InputConfluentCloudMaximumTLSVersion) ToPointer() *InputConfluentCloudMaximumTLSVersion {
	return &e
}
func (e *InputConfluentCloudMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputConfluentCloudMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudMaximumTLSVersion: %v", v)
	}
}

type InputConfluentCloudTLSSettingsClientSide struct {
	Disabled *bool `default:"false" json:"disabled"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another
	//                     trusted CA (e.g., the system's CA). Defaults to Yes. Overrides the toggle from Advanced Settings, when also present.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
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
	MinVersion *InputConfluentCloudMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *InputConfluentCloudMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (i InputConfluentCloudTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputConfluentCloudTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetMinVersion() *InputConfluentCloudMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputConfluentCloudTLSSettingsClientSide) GetMaxVersion() *InputConfluentCloudMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

// InputConfluentCloudAuth - Credentials to use when authenticating with the schema registry using basic HTTP authentication
type InputConfluentCloudAuth struct {
	// Enable authentication
	Disabled *bool `default:"true" json:"disabled"`
	// Select or create a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
}

func (i InputConfluentCloudAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputConfluentCloudAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputConfluentCloudAuth) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputConfluentCloudAuth) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

// InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion - Minimum TLS version to use when connecting
type InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion string

const (
	InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv1  InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1"
	InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv11 InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.1"
	InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv12 InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.2"
	InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv13 InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.3"
)

func (e InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion) ToPointer() *InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion {
	return &e
}
func (e *InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion: %v", v)
	}
}

// InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion - Maximum TLS version to use when connecting
type InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion string

const (
	InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv1  InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1"
	InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv11 InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.1"
	InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv12 InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.2"
	InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv13 InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.3"
)

func (e InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion) ToPointer() *InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion {
	return &e
}
func (e *InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion: %v", v)
	}
}

type InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another
	//                     trusted CA (e.g., the system's CA). Defaults to Yes. Overrides the toggle from Advanced Settings, when also present.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
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
	MinVersion *InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (i InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetMinVersion() *InputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetMaxVersion() *InputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type InputConfluentCloudKafkaSchemaRegistryAuthentication struct {
	// Enable Schema Registry
	Disabled *bool `default:"true" json:"disabled"`
	// URL for accessing the Confluent Schema Registry. Example: http://localhost:8081. To connect over TLS, use https instead of http.
	SchemaRegistryURL *string `default:"http://localhost:8081" json:"schemaRegistryURL"`
	// Maximum time to wait for a Schema Registry connection to complete successfully
	ConnectionTimeout *float64 `default:"30000" json:"connectionTimeout"`
	// Maximum time to wait for the Schema Registry to respond to a request
	RequestTimeout *float64 `default:"30000" json:"requestTimeout"`
	// Maximum number of times to try fetching schemas from the Schema Registry
	MaxRetries *float64 `default:"1" json:"maxRetries"`
	// Credentials to use when authenticating with the schema registry using basic HTTP authentication
	Auth *InputConfluentCloudAuth                                     `json:"auth,omitempty"`
	TLS  *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide `json:"tls,omitempty"`
}

func (i InputConfluentCloudKafkaSchemaRegistryAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputConfluentCloudKafkaSchemaRegistryAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputConfluentCloudKafkaSchemaRegistryAuthentication) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputConfluentCloudKafkaSchemaRegistryAuthentication) GetSchemaRegistryURL() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryURL
}

func (o *InputConfluentCloudKafkaSchemaRegistryAuthentication) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *InputConfluentCloudKafkaSchemaRegistryAuthentication) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputConfluentCloudKafkaSchemaRegistryAuthentication) GetMaxRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetries
}

func (o *InputConfluentCloudKafkaSchemaRegistryAuthentication) GetAuth() *InputConfluentCloudAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *InputConfluentCloudKafkaSchemaRegistryAuthentication) GetTLS() *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

// InputConfluentCloudSASLMechanism - SASL authentication mechanism to use.
type InputConfluentCloudSASLMechanism string

const (
	InputConfluentCloudSASLMechanismPlain       InputConfluentCloudSASLMechanism = "plain"
	InputConfluentCloudSASLMechanismScramSha256 InputConfluentCloudSASLMechanism = "scram-sha-256"
	InputConfluentCloudSASLMechanismScramSha512 InputConfluentCloudSASLMechanism = "scram-sha-512"
	InputConfluentCloudSASLMechanismKerberos    InputConfluentCloudSASLMechanism = "kerberos"
)

func (e InputConfluentCloudSASLMechanism) ToPointer() *InputConfluentCloudSASLMechanism {
	return &e
}
func (e *InputConfluentCloudSASLMechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plain":
		fallthrough
	case "scram-sha-256":
		fallthrough
	case "scram-sha-512":
		fallthrough
	case "kerberos":
		*e = InputConfluentCloudSASLMechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputConfluentCloudSASLMechanism: %v", v)
	}
}

// InputConfluentCloudAuthentication - Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
type InputConfluentCloudAuthentication struct {
	// Enable Authentication
	Disabled *bool `default:"true" json:"disabled"`
	// SASL authentication mechanism to use.
	Mechanism *InputConfluentCloudSASLMechanism `default:"plain" json:"mechanism"`
}

func (i InputConfluentCloudAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputConfluentCloudAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputConfluentCloudAuthentication) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputConfluentCloudAuthentication) GetMechanism() *InputConfluentCloudSASLMechanism {
	if o == nil {
		return nil
	}
	return o.Mechanism
}

type InputConfluentCloudMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputConfluentCloudMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputConfluentCloudMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputConfluentCloud struct {
	// Unique ID for this input
	ID       *string                  `json:"id,omitempty"`
	Type     *InputConfluentCloudType `json:"type,omitempty"`
	Disabled *bool                    `default:"false" json:"disabled"`
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
	Connections []InputConfluentCloudConnections `json:"connections,omitempty"`
	Pq          *InputConfluentCloudPq           `json:"pq,omitempty"`
	// List of Confluent Cloud bootstrap servers to use, such as yourAccount.confluent.cloud:9092
	Brokers []string                                  `json:"brokers"`
	TLS     *InputConfluentCloudTLSSettingsClientSide `json:"tls,omitempty"`
	// Topic to subscribe to. Warning: To optimize performance, Cribl suggests subscribing each Kafka Source to only a single topic.
	Topics []string `json:"topics,omitempty"`
	// Specifies the consumer group to which this instance belongs. Defaults to 'Cribl'.
	GroupID *string `default:"Cribl" json:"groupId"`
	// Leave toggled to 'Yes' if you want the Source, upon first subscribing to a topic, to read starting with the earliest available message
	FromBeginning       *bool                                                 `default:"true" json:"fromBeginning"`
	KafkaSchemaRegistry *InputConfluentCloudKafkaSchemaRegistryAuthentication `json:"kafkaSchemaRegistry,omitempty"`
	// Maximum time to wait for a connection to complete successfully
	ConnectionTimeout *float64 `default:"10000" json:"connectionTimeout"`
	// Maximum time to wait for Kafka to respond to a request
	RequestTimeout *float64 `default:"60000" json:"requestTimeout"`
	// If messages are failing, you can set the maximum number of retries as high as 100 to prevent loss of data
	MaxRetries *float64 `default:"5" json:"maxRetries"`
	// The maximum wait time for a retry, in milliseconds. Default (and minimum) is 30,000 ms (30 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackOff *float64 `default:"30000" json:"maxBackOff"`
	// Initial value used to calculate the retry, in milliseconds. Maximum is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"300" json:"initialBackoff"`
	// Set the backoff multiplier (2-20) to control the retry frequency for failed messages. For faster retries, use a lower multiplier. For slower retries with more delay between attempts, use a higher multiplier. The multiplier is used in an exponential backoff formula; see the Kafka [documentation](https://kafka.js.org/docs/retry-detailed) for details.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// Maximum time to wait for Kafka to respond to an authentication request
	AuthenticationTimeout *float64 `default:"10000" json:"authenticationTimeout"`
	// Specifies a time window during which @{product} can reauthenticate if needed. Creates the window measuring backwards from the moment when credentials are set to expire.
	ReauthenticationThreshold *float64 `default:"10000" json:"reauthenticationThreshold"`
	// Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
	Sasl *InputConfluentCloudAuthentication `json:"sasl,omitempty"`
	//       Timeout used to detect client failures when using Kafka's group management facilities.
	//       If the client sends the broker no heartbeats before this timeout expires,
	//       the broker will remove this client from the group, and will initiate a rebalance.
	//       Value must be between the broker's configured group.min.session.timeout.ms and group.max.session.timeout.ms.
	//       See details [here](https://kafka.apache.org/documentation/#consumerconfigs_session.timeout.ms).
	SessionTimeout *float64 `default:"30000" json:"sessionTimeout"`
	//       Maximum allowed time for each worker to join the group after a rebalance has begun.
	//       If the timeout is exceeded, the coordinator broker will remove the worker from the group.
	//       See details [here](https://kafka.apache.org/documentation/#connectconfigs_rebalance.timeout.ms).
	RebalanceTimeout *float64 `default:"60000" json:"rebalanceTimeout"`
	//       Expected time between heartbeats to the consumer coordinator when using Kafka's group management facilities.
	//       Value must be lower than sessionTimeout, and typically should not exceed 1/3 of the sessionTimeout value.
	//       See details [here](https://kafka.apache.org/documentation/#consumerconfigs_heartbeat.interval.ms).
	HeartbeatInterval *float64 `default:"3000" json:"heartbeatInterval"`
	// How often to commit offsets. If both this and Offset commit threshold are set, @{product} commits offsets when either condition is met. If both are empty, @{product} commits offsets after each batch.
	AutoCommitInterval *float64 `json:"autoCommitInterval,omitempty"`
	// How many events are needed to trigger an offset commit. If both this and Offset commit interval are set, @{product} commits offsets when either condition is met. If both are empty, @{product} commits offsets after each batch.
	AutoCommitThreshold *float64 `json:"autoCommitThreshold,omitempty"`
	// Maximum amount of data that Kafka will return per partition, per fetch request. Must equal or exceed the maximum message size (maxBytesPerPartition) that Kafka is configured to allow. Otherwise, @{product} can get stuck trying to retrieve messages. Defaults to 1048576 (1 MB).
	MaxBytesPerPartition *float64 `default:"1048576" json:"maxBytesPerPartition"`
	// Maximum number of bytes that Kafka will return per fetch request. Defaults to 10485760 (10 MB).
	MaxBytes *float64 `default:"10485760" json:"maxBytes"`
	// Maximum number of network errors before the consumer recreates a socket.
	MaxSocketErrors *float64 `default:"0" json:"maxSocketErrors"`
	// Fields to add to events from this input
	Metadata    []InputConfluentCloudMetadata `json:"metadata,omitempty"`
	Description *string                       `json:"description,omitempty"`
}

func (i InputConfluentCloud) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputConfluentCloud) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputConfluentCloud) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputConfluentCloud) GetType() *InputConfluentCloudType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputConfluentCloud) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputConfluentCloud) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputConfluentCloud) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputConfluentCloud) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputConfluentCloud) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputConfluentCloud) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputConfluentCloud) GetConnections() []InputConfluentCloudConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputConfluentCloud) GetPq() *InputConfluentCloudPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputConfluentCloud) GetBrokers() []string {
	if o == nil {
		return []string{}
	}
	return o.Brokers
}

func (o *InputConfluentCloud) GetTLS() *InputConfluentCloudTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputConfluentCloud) GetTopics() []string {
	if o == nil {
		return nil
	}
	return o.Topics
}

func (o *InputConfluentCloud) GetGroupID() *string {
	if o == nil {
		return nil
	}
	return o.GroupID
}

func (o *InputConfluentCloud) GetFromBeginning() *bool {
	if o == nil {
		return nil
	}
	return o.FromBeginning
}

func (o *InputConfluentCloud) GetKafkaSchemaRegistry() *InputConfluentCloudKafkaSchemaRegistryAuthentication {
	if o == nil {
		return nil
	}
	return o.KafkaSchemaRegistry
}

func (o *InputConfluentCloud) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *InputConfluentCloud) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputConfluentCloud) GetMaxRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetries
}

func (o *InputConfluentCloud) GetMaxBackOff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackOff
}

func (o *InputConfluentCloud) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *InputConfluentCloud) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *InputConfluentCloud) GetAuthenticationTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.AuthenticationTimeout
}

func (o *InputConfluentCloud) GetReauthenticationThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.ReauthenticationThreshold
}

func (o *InputConfluentCloud) GetSasl() *InputConfluentCloudAuthentication {
	if o == nil {
		return nil
	}
	return o.Sasl
}

func (o *InputConfluentCloud) GetSessionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionTimeout
}

func (o *InputConfluentCloud) GetRebalanceTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RebalanceTimeout
}

func (o *InputConfluentCloud) GetHeartbeatInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.HeartbeatInterval
}

func (o *InputConfluentCloud) GetAutoCommitInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.AutoCommitInterval
}

func (o *InputConfluentCloud) GetAutoCommitThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.AutoCommitThreshold
}

func (o *InputConfluentCloud) GetMaxBytesPerPartition() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBytesPerPartition
}

func (o *InputConfluentCloud) GetMaxBytes() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *InputConfluentCloud) GetMaxSocketErrors() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxSocketErrors
}

func (o *InputConfluentCloud) GetMetadata() []InputConfluentCloudMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputConfluentCloud) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
