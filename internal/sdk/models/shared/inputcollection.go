// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputCollectionType string

const (
	InputCollectionTypeCollection InputCollectionType = "collection"
)

func (e InputCollectionType) ToPointer() *InputCollectionType {
	return &e
}
func (e *InputCollectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "collection":
		*e = InputCollectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCollectionType: %v", v)
	}
}

type Connections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *Connections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *Connections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputCollectionMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputCollectionMode string

const (
	InputCollectionModeSmart  InputCollectionMode = "smart"
	InputCollectionModeAlways InputCollectionMode = "always"
)

func (e InputCollectionMode) ToPointer() *InputCollectionMode {
	return &e
}
func (e *InputCollectionMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputCollectionMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCollectionMode: %v", v)
	}
}

// InputCollectionCompression - Codec to use to compress the persisted data
type InputCollectionCompression string

const (
	InputCollectionCompressionNone InputCollectionCompression = "none"
	InputCollectionCompressionGzip InputCollectionCompression = "gzip"
)

func (e InputCollectionCompression) ToPointer() *InputCollectionCompression {
	return &e
}
func (e *InputCollectionCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputCollectionCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCollectionCompression: %v", v)
	}
}

type Pq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputCollectionMode `default:"always" json:"mode"`
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
	Compress *InputCollectionCompression `default:"none" json:"compress"`
}

func (p Pq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Pq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Pq) GetMode() *InputCollectionMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *Pq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *Pq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *Pq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *Pq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *Pq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *Pq) GetCompress() *InputCollectionCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type Preprocess struct {
	Disabled *bool `default:"true" json:"disabled"`
	// Command to feed the data through (via stdin) and process its output (stdout)
	Command *string `json:"command,omitempty"`
	// Arguments to be added to the custom command
	Args []string `json:"args,omitempty"`
}

func (p Preprocess) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Preprocess) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Preprocess) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *Preprocess) GetCommand() *string {
	if o == nil {
		return nil
	}
	return o.Command
}

func (o *Preprocess) GetArgs() []string {
	if o == nil {
		return nil
	}
	return o.Args
}

type InputCollectionMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputCollectionMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputCollectionMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputCollection struct {
	// Unique ID for this input
	ID       string               `json:"id"`
	Type     *InputCollectionType `default:"collection" json:"type"`
	Disabled *bool                `default:"false" json:"disabled"`
	// Pipeline to process results
	Pipeline *string `json:"pipeline,omitempty"`
	// Send events to normal routing and event processing. Disable to select a specific Pipeline/Destination combination.
	SendToRoutes *bool `default:"true" json:"sendToRoutes"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Use a disk queue to minimize data loss when connected services block. See [Cribl Docs](https://docs.cribl.io/stream/persistent-queues) for PQ defaults (Cribl-managed Cloud Workers) and configuration options (on-prem and hybrid Workers).
	PqEnabled *bool `default:"false" json:"pqEnabled"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Direct connections to Destinations, and optionally via a Pipeline or a Pack
	Connections []Connections `json:"connections,omitempty"`
	Pq          *Pq           `json:"pq,omitempty"`
	// A list of event-breaking rulesets that will be applied, in order, to the input data stream
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// How long (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel before flushing the data stream out, as is, to the Pipelines
	StaleChannelFlushMs *float64    `default:"10000" json:"staleChannelFlushMs"`
	Preprocess          *Preprocess `json:"preprocess,omitempty"`
	// Rate (in bytes per second) to throttle while writing to an output. Accepts values with multiple-byte units, such as KB, MB, and GB. (Example: 42 MB) Default value of 0 specifies no throttling.
	ThrottleRatePerSec *string `default:"0" json:"throttleRatePerSec"`
	// Fields to add to events from this input
	Metadata []InputCollectionMetadata `json:"metadata,omitempty"`
	// Destination to send results to
	Output *string `json:"output,omitempty"`
}

func (i InputCollection) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputCollection) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputCollection) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InputCollection) GetType() *InputCollectionType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputCollection) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputCollection) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputCollection) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputCollection) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputCollection) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputCollection) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputCollection) GetConnections() []Connections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputCollection) GetPq() *Pq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputCollection) GetBreakerRulesets() []string {
	if o == nil {
		return nil
	}
	return o.BreakerRulesets
}

func (o *InputCollection) GetStaleChannelFlushMs() *float64 {
	if o == nil {
		return nil
	}
	return o.StaleChannelFlushMs
}

func (o *InputCollection) GetPreprocess() *Preprocess {
	if o == nil {
		return nil
	}
	return o.Preprocess
}

func (o *InputCollection) GetThrottleRatePerSec() *string {
	if o == nil {
		return nil
	}
	return o.ThrottleRatePerSec
}

func (o *InputCollection) GetMetadata() []InputCollectionMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputCollection) GetOutput() *string {
	if o == nil {
		return nil
	}
	return o.Output
}
