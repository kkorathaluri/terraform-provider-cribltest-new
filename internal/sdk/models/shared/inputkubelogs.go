// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputKubeLogsType string

const (
	InputKubeLogsTypeKubeLogs InputKubeLogsType = "kube_logs"
)

func (e InputKubeLogsType) ToPointer() *InputKubeLogsType {
	return &e
}
func (e *InputKubeLogsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kube_logs":
		*e = InputKubeLogsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputKubeLogsType: %v", v)
	}
}

type InputKubeLogsConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputKubeLogsConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputKubeLogsConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputKubeLogsMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputKubeLogsMode string

const (
	InputKubeLogsModeSmart  InputKubeLogsMode = "smart"
	InputKubeLogsModeAlways InputKubeLogsMode = "always"
)

func (e InputKubeLogsMode) ToPointer() *InputKubeLogsMode {
	return &e
}
func (e *InputKubeLogsMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputKubeLogsMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputKubeLogsMode: %v", v)
	}
}

// InputKubeLogsCompression - Codec to use to compress the persisted data
type InputKubeLogsCompression string

const (
	InputKubeLogsCompressionNone InputKubeLogsCompression = "none"
	InputKubeLogsCompressionGzip InputKubeLogsCompression = "gzip"
)

func (e InputKubeLogsCompression) ToPointer() *InputKubeLogsCompression {
	return &e
}
func (e *InputKubeLogsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputKubeLogsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputKubeLogsCompression: %v", v)
	}
}

type InputKubeLogsPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputKubeLogsMode `default:"always" json:"mode"`
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
	Compress *InputKubeLogsCompression `default:"none" json:"compress"`
}

func (i InputKubeLogsPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputKubeLogsPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputKubeLogsPq) GetMode() *InputKubeLogsMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputKubeLogsPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputKubeLogsPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputKubeLogsPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputKubeLogsPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputKubeLogsPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputKubeLogsPq) GetCompress() *InputKubeLogsCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type InputKubeLogsRules struct {
	// JavaScript expression applied to Pod objects. Return 'true' to include it.
	Filter string `json:"filter"`
	// Optional description of this rule's purpose
	Description *string `json:"description,omitempty"`
}

func (o *InputKubeLogsRules) GetFilter() string {
	if o == nil {
		return ""
	}
	return o.Filter
}

func (o *InputKubeLogsRules) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type InputKubeLogsMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputKubeLogsMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputKubeLogsMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// InputKubeLogsPersistenceCompression - Data compression format. Default is gzip.
type InputKubeLogsPersistenceCompression string

const (
	InputKubeLogsPersistenceCompressionNone InputKubeLogsPersistenceCompression = "none"
	InputKubeLogsPersistenceCompressionGzip InputKubeLogsPersistenceCompression = "gzip"
)

func (e InputKubeLogsPersistenceCompression) ToPointer() *InputKubeLogsPersistenceCompression {
	return &e
}
func (e *InputKubeLogsPersistenceCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputKubeLogsPersistenceCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputKubeLogsPersistenceCompression: %v", v)
	}
}

type InputKubeLogsDiskSpooling struct {
	// Spool events on disk for Cribl Edge and Search. Default is disabled.
	Enable *bool `default:"false" json:"enable"`
	// Time period for grouping spooled events. Default is 10m.
	TimeWindow *string `default:"10m" json:"timeWindow"`
	// Maximum disk space that can be consumed before older buckets are deleted. Examples: 420MB, 4GB. Default is 1GB.
	MaxDataSize *string `default:"1GB" json:"maxDataSize"`
	// Maximum amount of time to retain data before older buckets are deleted. Examples: 2h, 4d. Default is 24h.
	MaxDataTime *string `default:"24h" json:"maxDataTime"`
	// Data compression format. Default is gzip.
	Compress *InputKubeLogsPersistenceCompression `default:"gzip" json:"compress"`
}

func (i InputKubeLogsDiskSpooling) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputKubeLogsDiskSpooling) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputKubeLogsDiskSpooling) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

func (o *InputKubeLogsDiskSpooling) GetTimeWindow() *string {
	if o == nil {
		return nil
	}
	return o.TimeWindow
}

func (o *InputKubeLogsDiskSpooling) GetMaxDataSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataSize
}

func (o *InputKubeLogsDiskSpooling) GetMaxDataTime() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataTime
}

func (o *InputKubeLogsDiskSpooling) GetCompress() *InputKubeLogsPersistenceCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type InputKubeLogs struct {
	// Unique ID for this input
	ID       string            `json:"id"`
	Type     InputKubeLogsType `json:"type"`
	Disabled *bool             `default:"false" json:"disabled"`
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
	Connections []InputKubeLogsConnections `json:"connections,omitempty"`
	Pq          *InputKubeLogsPq           `json:"pq,omitempty"`
	// Time, in seconds, between checks for new containers. Default is 15 secs.
	Interval *float64 `default:"15" json:"interval"`
	// Add rules to decide which Pods to collect logs from. Logs are collected if no rules are given or if all the rules' expressions evaluate to true.
	Rules []InputKubeLogsRules `json:"rules,omitempty"`
	// For use when containers do not emit a timestamp, prefix each line of output with a timestamp. If you enable this setting, you can use the Kubernetes Logs Event Breaker and the kubernetes_logs Pre-processing Pipeline to remove them from the events after the timestamps are extracted.
	Timestamps *bool `default:"false" json:"timestamps"`
	// Fields to add to events from this input
	Metadata    []InputKubeLogsMetadata    `json:"metadata,omitempty"`
	Persistence *InputKubeLogsDiskSpooling `json:"persistence,omitempty"`
	// A list of event-breaking rulesets that will be applied, in order, to the input data stream
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// How long (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel before flushing the data stream out, as is, to the Pipelines
	StaleChannelFlushMs *float64 `default:"10000" json:"staleChannelFlushMs"`
	// Load balance traffic across all Worker Processes
	EnableLoadBalancing *bool   `default:"false" json:"enableLoadBalancing"`
	Description         *string `json:"description,omitempty"`
}

func (i InputKubeLogs) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputKubeLogs) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputKubeLogs) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InputKubeLogs) GetType() InputKubeLogsType {
	if o == nil {
		return InputKubeLogsType("")
	}
	return o.Type
}

func (o *InputKubeLogs) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputKubeLogs) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputKubeLogs) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputKubeLogs) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputKubeLogs) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputKubeLogs) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputKubeLogs) GetConnections() []InputKubeLogsConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputKubeLogs) GetPq() *InputKubeLogsPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputKubeLogs) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *InputKubeLogs) GetRules() []InputKubeLogsRules {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *InputKubeLogs) GetTimestamps() *bool {
	if o == nil {
		return nil
	}
	return o.Timestamps
}

func (o *InputKubeLogs) GetMetadata() []InputKubeLogsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputKubeLogs) GetPersistence() *InputKubeLogsDiskSpooling {
	if o == nil {
		return nil
	}
	return o.Persistence
}

func (o *InputKubeLogs) GetBreakerRulesets() []string {
	if o == nil {
		return nil
	}
	return o.BreakerRulesets
}

func (o *InputKubeLogs) GetStaleChannelFlushMs() *float64 {
	if o == nil {
		return nil
	}
	return o.StaleChannelFlushMs
}

func (o *InputKubeLogs) GetEnableLoadBalancing() *bool {
	if o == nil {
		return nil
	}
	return o.EnableLoadBalancing
}

func (o *InputKubeLogs) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
