// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputJournalFilesType string

const (
	InputJournalFilesTypeJournalFiles InputJournalFilesType = "journal_files"
)

func (e InputJournalFilesType) ToPointer() *InputJournalFilesType {
	return &e
}
func (e *InputJournalFilesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "journal_files":
		*e = InputJournalFilesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputJournalFilesType: %v", v)
	}
}

type InputJournalFilesConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputJournalFilesConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputJournalFilesConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputJournalFilesMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputJournalFilesMode string

const (
	InputJournalFilesModeSmart  InputJournalFilesMode = "smart"
	InputJournalFilesModeAlways InputJournalFilesMode = "always"
)

func (e InputJournalFilesMode) ToPointer() *InputJournalFilesMode {
	return &e
}
func (e *InputJournalFilesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputJournalFilesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputJournalFilesMode: %v", v)
	}
}

// InputJournalFilesCompression - Codec to use to compress the persisted data
type InputJournalFilesCompression string

const (
	InputJournalFilesCompressionNone InputJournalFilesCompression = "none"
	InputJournalFilesCompressionGzip InputJournalFilesCompression = "gzip"
)

func (e InputJournalFilesCompression) ToPointer() *InputJournalFilesCompression {
	return &e
}
func (e *InputJournalFilesCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputJournalFilesCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputJournalFilesCompression: %v", v)
	}
}

type InputJournalFilesPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputJournalFilesMode `default:"always" json:"mode"`
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
	Compress *InputJournalFilesCompression `default:"none" json:"compress"`
}

func (i InputJournalFilesPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputJournalFilesPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputJournalFilesPq) GetMode() *InputJournalFilesMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputJournalFilesPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputJournalFilesPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputJournalFilesPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputJournalFilesPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputJournalFilesPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputJournalFilesPq) GetCompress() *InputJournalFilesCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type InputJournalFilesRules struct {
	// JavaScript expression applied to Journal objects. Return 'true' to include it.
	Filter string `json:"filter"`
	// Optional description of this rule's purpose
	Description *string `json:"description,omitempty"`
}

func (o *InputJournalFilesRules) GetFilter() string {
	if o == nil {
		return ""
	}
	return o.Filter
}

func (o *InputJournalFilesRules) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type InputJournalFilesMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputJournalFilesMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputJournalFilesMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputJournalFiles struct {
	// Unique ID for this input
	ID       *string                `json:"id,omitempty"`
	Type     *InputJournalFilesType `json:"type,omitempty"`
	Disabled *bool                  `default:"false" json:"disabled"`
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
	Connections []InputJournalFilesConnections `json:"connections,omitempty"`
	Pq          *InputJournalFilesPq           `json:"pq,omitempty"`
	// Directory path to search for journals. Environment variables will be resolved, e.g. $CRIBL_EDGE_FS_ROOT/var/log/journal/$MACHINE_ID.
	Path string `json:"path"`
	// Time, in seconds, between scanning for journals.
	Interval *float64 `default:"10" json:"interval"`
	// The full path of discovered journals are matched against this wildcard list.
	Journals []string `json:"journals,omitempty"`
	// Add rules to decide which journal objects to allow. Events are generated if no rules are given or if all the rules' expressions evaluate to true.
	Rules []InputJournalFilesRules `json:"rules,omitempty"`
	// Skip log messages that are not part of the current boot session.
	CurrentBoot *bool `default:"false" json:"currentBoot"`
	// The maximum log message age, in duration form (e.g,: 60s, 4h, 3d, 1w).  Default of no value will apply no max age filters.
	MaxAgeDur *string `json:"maxAgeDur,omitempty"`
	// Fields to add to events from this input
	Metadata    []InputJournalFilesMetadata `json:"metadata,omitempty"`
	Description *string                     `json:"description,omitempty"`
}

func (i InputJournalFiles) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputJournalFiles) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputJournalFiles) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputJournalFiles) GetType() *InputJournalFilesType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputJournalFiles) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputJournalFiles) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputJournalFiles) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputJournalFiles) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputJournalFiles) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputJournalFiles) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputJournalFiles) GetConnections() []InputJournalFilesConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputJournalFiles) GetPq() *InputJournalFilesPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputJournalFiles) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *InputJournalFiles) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *InputJournalFiles) GetJournals() []string {
	if o == nil {
		return nil
	}
	return o.Journals
}

func (o *InputJournalFiles) GetRules() []InputJournalFilesRules {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *InputJournalFiles) GetCurrentBoot() *bool {
	if o == nil {
		return nil
	}
	return o.CurrentBoot
}

func (o *InputJournalFiles) GetMaxAgeDur() *string {
	if o == nil {
		return nil
	}
	return o.MaxAgeDur
}

func (o *InputJournalFiles) GetMetadata() []InputJournalFilesMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputJournalFiles) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
