// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputSystemMetricsType string

const (
	InputSystemMetricsTypeSystemMetrics InputSystemMetricsType = "system_metrics"
)

func (e InputSystemMetricsType) ToPointer() *InputSystemMetricsType {
	return &e
}
func (e *InputSystemMetricsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system_metrics":
		*e = InputSystemMetricsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsType: %v", v)
	}
}

type InputSystemMetricsConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputSystemMetricsConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSystemMetricsConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputSystemMetricsPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputSystemMetricsPqMode string

const (
	InputSystemMetricsPqModeSmart  InputSystemMetricsPqMode = "smart"
	InputSystemMetricsPqModeAlways InputSystemMetricsPqMode = "always"
)

func (e InputSystemMetricsPqMode) ToPointer() *InputSystemMetricsPqMode {
	return &e
}
func (e *InputSystemMetricsPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputSystemMetricsPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsPqMode: %v", v)
	}
}

// InputSystemMetricsCompression - Codec to use to compress the persisted data
type InputSystemMetricsCompression string

const (
	InputSystemMetricsCompressionNone InputSystemMetricsCompression = "none"
	InputSystemMetricsCompressionGzip InputSystemMetricsCompression = "gzip"
)

func (e InputSystemMetricsCompression) ToPointer() *InputSystemMetricsCompression {
	return &e
}
func (e *InputSystemMetricsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSystemMetricsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsCompression: %v", v)
	}
}

type InputSystemMetricsPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputSystemMetricsPqMode `default:"always" json:"mode"`
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
	Compress *InputSystemMetricsCompression `default:"none" json:"compress"`
}

func (i InputSystemMetricsPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemMetricsPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemMetricsPq) GetMode() *InputSystemMetricsPqMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSystemMetricsPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputSystemMetricsPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputSystemMetricsPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputSystemMetricsPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputSystemMetricsPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputSystemMetricsPq) GetCompress() *InputSystemMetricsCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputSystemMetricsMode - Select level of detail for host metrics
type InputSystemMetricsMode string

const (
	InputSystemMetricsModeBasic    InputSystemMetricsMode = "basic"
	InputSystemMetricsModeAll      InputSystemMetricsMode = "all"
	InputSystemMetricsModeCustom   InputSystemMetricsMode = "custom"
	InputSystemMetricsModeDisabled InputSystemMetricsMode = "disabled"
)

func (e InputSystemMetricsMode) ToPointer() *InputSystemMetricsMode {
	return &e
}
func (e *InputSystemMetricsMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputSystemMetricsMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsMode: %v", v)
	}
}

// InputSystemMetricsHostMode - Select the level of detail for system metrics
type InputSystemMetricsHostMode string

const (
	InputSystemMetricsHostModeBasic    InputSystemMetricsHostMode = "basic"
	InputSystemMetricsHostModeAll      InputSystemMetricsHostMode = "all"
	InputSystemMetricsHostModeCustom   InputSystemMetricsHostMode = "custom"
	InputSystemMetricsHostModeDisabled InputSystemMetricsHostMode = "disabled"
)

func (e InputSystemMetricsHostMode) ToPointer() *InputSystemMetricsHostMode {
	return &e
}
func (e *InputSystemMetricsHostMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputSystemMetricsHostMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsHostMode: %v", v)
	}
}

type InputSystemMetricsSystem struct {
	// Select the level of detail for system metrics
	Mode *InputSystemMetricsHostMode `default:"basic" json:"mode"`
	// Generate metrics for the numbers of processes in various states
	Processes *bool `default:"false" json:"processes"`
}

func (i InputSystemMetricsSystem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemMetricsSystem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemMetricsSystem) GetMode() *InputSystemMetricsHostMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSystemMetricsSystem) GetProcesses() *bool {
	if o == nil {
		return nil
	}
	return o.Processes
}

// InputSystemMetricsHostCustomMode - Select the level of detail for CPU metrics
type InputSystemMetricsHostCustomMode string

const (
	InputSystemMetricsHostCustomModeBasic    InputSystemMetricsHostCustomMode = "basic"
	InputSystemMetricsHostCustomModeAll      InputSystemMetricsHostCustomMode = "all"
	InputSystemMetricsHostCustomModeCustom   InputSystemMetricsHostCustomMode = "custom"
	InputSystemMetricsHostCustomModeDisabled InputSystemMetricsHostCustomMode = "disabled"
)

func (e InputSystemMetricsHostCustomMode) ToPointer() *InputSystemMetricsHostCustomMode {
	return &e
}
func (e *InputSystemMetricsHostCustomMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputSystemMetricsHostCustomMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsHostCustomMode: %v", v)
	}
}

type CPU struct {
	// Select the level of detail for CPU metrics
	Mode *InputSystemMetricsHostCustomMode `default:"basic" json:"mode"`
	// Generate metrics for each CPU
	PerCPU *bool `default:"false" json:"perCpu"`
	// Generate metrics for all CPU states
	Detail *bool `default:"false" json:"detail"`
	// Generate raw, monotonic CPU time counters
	Time *bool `default:"false" json:"time"`
}

func (c CPU) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CPU) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CPU) GetMode() *InputSystemMetricsHostCustomMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *CPU) GetPerCPU() *bool {
	if o == nil {
		return nil
	}
	return o.PerCPU
}

func (o *CPU) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *CPU) GetTime() *bool {
	if o == nil {
		return nil
	}
	return o.Time
}

// InputSystemMetricsHostCustomMemoryMode - Select the level of detail for memory metrics
type InputSystemMetricsHostCustomMemoryMode string

const (
	InputSystemMetricsHostCustomMemoryModeBasic    InputSystemMetricsHostCustomMemoryMode = "basic"
	InputSystemMetricsHostCustomMemoryModeAll      InputSystemMetricsHostCustomMemoryMode = "all"
	InputSystemMetricsHostCustomMemoryModeCustom   InputSystemMetricsHostCustomMemoryMode = "custom"
	InputSystemMetricsHostCustomMemoryModeDisabled InputSystemMetricsHostCustomMemoryMode = "disabled"
)

func (e InputSystemMetricsHostCustomMemoryMode) ToPointer() *InputSystemMetricsHostCustomMemoryMode {
	return &e
}
func (e *InputSystemMetricsHostCustomMemoryMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputSystemMetricsHostCustomMemoryMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsHostCustomMemoryMode: %v", v)
	}
}

type InputSystemMetricsMemory struct {
	// Select the level of detail for memory metrics
	Mode *InputSystemMetricsHostCustomMemoryMode `default:"basic" json:"mode"`
	// Generate metrics for all memory states
	Detail *bool `default:"false" json:"detail"`
}

func (i InputSystemMetricsMemory) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemMetricsMemory) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemMetricsMemory) GetMode() *InputSystemMetricsHostCustomMemoryMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSystemMetricsMemory) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

// InputSystemMetricsHostCustomNetworkMode - Select the level of detail for network metrics
type InputSystemMetricsHostCustomNetworkMode string

const (
	InputSystemMetricsHostCustomNetworkModeBasic    InputSystemMetricsHostCustomNetworkMode = "basic"
	InputSystemMetricsHostCustomNetworkModeAll      InputSystemMetricsHostCustomNetworkMode = "all"
	InputSystemMetricsHostCustomNetworkModeCustom   InputSystemMetricsHostCustomNetworkMode = "custom"
	InputSystemMetricsHostCustomNetworkModeDisabled InputSystemMetricsHostCustomNetworkMode = "disabled"
)

func (e InputSystemMetricsHostCustomNetworkMode) ToPointer() *InputSystemMetricsHostCustomNetworkMode {
	return &e
}
func (e *InputSystemMetricsHostCustomNetworkMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputSystemMetricsHostCustomNetworkMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsHostCustomNetworkMode: %v", v)
	}
}

type Network struct {
	// Select the level of detail for network metrics
	Mode *InputSystemMetricsHostCustomNetworkMode `default:"basic" json:"mode"`
	// Network interfaces to include/exclude. Examples: eth0, !lo. All interfaces are included if this list is empty.
	Devices []string `json:"devices,omitempty"`
	// Generate separate metrics for each interface
	PerInterface *bool `default:"false" json:"perInterface"`
	// Generate full network metrics
	Detail *bool `default:"false" json:"detail"`
}

func (n Network) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *Network) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Network) GetMode() *InputSystemMetricsHostCustomNetworkMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *Network) GetDevices() []string {
	if o == nil {
		return nil
	}
	return o.Devices
}

func (o *Network) GetPerInterface() *bool {
	if o == nil {
		return nil
	}
	return o.PerInterface
}

func (o *Network) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

// InputSystemMetricsHostCustomDiskMode - Select the level of detail for disk metrics
type InputSystemMetricsHostCustomDiskMode string

const (
	InputSystemMetricsHostCustomDiskModeBasic    InputSystemMetricsHostCustomDiskMode = "basic"
	InputSystemMetricsHostCustomDiskModeAll      InputSystemMetricsHostCustomDiskMode = "all"
	InputSystemMetricsHostCustomDiskModeCustom   InputSystemMetricsHostCustomDiskMode = "custom"
	InputSystemMetricsHostCustomDiskModeDisabled InputSystemMetricsHostCustomDiskMode = "disabled"
)

func (e InputSystemMetricsHostCustomDiskMode) ToPointer() *InputSystemMetricsHostCustomDiskMode {
	return &e
}
func (e *InputSystemMetricsHostCustomDiskMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputSystemMetricsHostCustomDiskMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsHostCustomDiskMode: %v", v)
	}
}

type Disk struct {
	// Select the level of detail for disk metrics
	Mode *InputSystemMetricsHostCustomDiskMode `default:"basic" json:"mode"`
	// Block devices to include/exclude. Examples: sda*, !loop*. Wildcards and ! (not) operators are supported. All devices are included if this list is empty.
	Devices []string `json:"devices,omitempty"`
	// Filesystem mountpoints to include/exclude. Examples: /, /home, !/proc*, !/tmp. Wildcards and ! (not) operators are supported. All mountpoints are included if this list is empty.
	Mountpoints []string `json:"mountpoints,omitempty"`
	// Filesystem types to include/exclude. Examples: ext4, !*tmpfs, !squashfs. Wildcards and ! (not) operators are supported. All types are included if this list is empty.
	Fstypes []string `json:"fstypes,omitempty"`
	// Generate separate metrics for each device
	PerDevice *bool `default:"false" json:"perDevice"`
	// Generate full disk metrics
	Detail *bool `default:"false" json:"detail"`
}

func (d Disk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Disk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Disk) GetMode() *InputSystemMetricsHostCustomDiskMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *Disk) GetDevices() []string {
	if o == nil {
		return nil
	}
	return o.Devices
}

func (o *Disk) GetMountpoints() []string {
	if o == nil {
		return nil
	}
	return o.Mountpoints
}

func (o *Disk) GetFstypes() []string {
	if o == nil {
		return nil
	}
	return o.Fstypes
}

func (o *Disk) GetPerDevice() *bool {
	if o == nil {
		return nil
	}
	return o.PerDevice
}

func (o *Disk) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

type Custom struct {
	System  *InputSystemMetricsSystem `json:"system,omitempty"`
	CPU     *CPU                      `json:"cpu,omitempty"`
	Memory  *InputSystemMetricsMemory `json:"memory,omitempty"`
	Network *Network                  `json:"network,omitempty"`
	Disk    *Disk                     `json:"disk,omitempty"`
}

func (o *Custom) GetSystem() *InputSystemMetricsSystem {
	if o == nil {
		return nil
	}
	return o.System
}

func (o *Custom) GetCPU() *CPU {
	if o == nil {
		return nil
	}
	return o.CPU
}

func (o *Custom) GetMemory() *InputSystemMetricsMemory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *Custom) GetNetwork() *Network {
	if o == nil {
		return nil
	}
	return o.Network
}

func (o *Custom) GetDisk() *Disk {
	if o == nil {
		return nil
	}
	return o.Disk
}

type Host struct {
	// Select level of detail for host metrics
	Mode   *InputSystemMetricsMode `default:"basic" json:"mode"`
	Custom *Custom                 `json:"custom,omitempty"`
}

func (h Host) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *Host) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Host) GetMode() *InputSystemMetricsMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *Host) GetCustom() *Custom {
	if o == nil {
		return nil
	}
	return o.Custom
}

type Sets struct {
	Name            string `json:"name"`
	Filter          string `json:"filter"`
	IncludeChildren *bool  `default:"false" json:"includeChildren"`
}

func (s Sets) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Sets) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Sets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Sets) GetFilter() string {
	if o == nil {
		return ""
	}
	return o.Filter
}

func (o *Sets) GetIncludeChildren() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeChildren
}

type InputSystemMetricsProcess struct {
	// Configure sets to collect process metrics
	Sets []Sets `json:"sets,omitempty"`
}

func (o *InputSystemMetricsProcess) GetSets() []Sets {
	if o == nil {
		return nil
	}
	return o.Sets
}

// InputSystemMetricsContainerMode - Select the level of detail for container metrics
type InputSystemMetricsContainerMode string

const (
	InputSystemMetricsContainerModeBasic    InputSystemMetricsContainerMode = "basic"
	InputSystemMetricsContainerModeAll      InputSystemMetricsContainerMode = "all"
	InputSystemMetricsContainerModeCustom   InputSystemMetricsContainerMode = "custom"
	InputSystemMetricsContainerModeDisabled InputSystemMetricsContainerMode = "disabled"
)

func (e InputSystemMetricsContainerMode) ToPointer() *InputSystemMetricsContainerMode {
	return &e
}
func (e *InputSystemMetricsContainerMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputSystemMetricsContainerMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsContainerMode: %v", v)
	}
}

type Filters struct {
	Expr string `json:"expr"`
}

func (o *Filters) GetExpr() string {
	if o == nil {
		return ""
	}
	return o.Expr
}

type InputSystemMetricsContainer struct {
	// Select the level of detail for container metrics
	Mode *InputSystemMetricsContainerMode `default:"basic" json:"mode"`
	// Full paths for Docker's UNIX-domain socket
	DockerSocket []string `json:"dockerSocket,omitempty"`
	// Timeout, in seconds, for the Docker API
	DockerTimeout *float64 `default:"5" json:"dockerTimeout"`
	// Containers matching any of these will be included. All are included if no filters are added.
	Filters []Filters `json:"filters,omitempty"`
	// Include stopped and paused containers
	AllContainers *bool `default:"false" json:"allContainers"`
	// Generate separate metrics for each device
	PerDevice *bool `default:"false" json:"perDevice"`
	// Generate full container metrics
	Detail *bool `default:"false" json:"detail"`
}

func (i InputSystemMetricsContainer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemMetricsContainer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemMetricsContainer) GetMode() *InputSystemMetricsContainerMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSystemMetricsContainer) GetDockerSocket() []string {
	if o == nil {
		return nil
	}
	return o.DockerSocket
}

func (o *InputSystemMetricsContainer) GetDockerTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.DockerTimeout
}

func (o *InputSystemMetricsContainer) GetFilters() []Filters {
	if o == nil {
		return nil
	}
	return o.Filters
}

func (o *InputSystemMetricsContainer) GetAllContainers() *bool {
	if o == nil {
		return nil
	}
	return o.AllContainers
}

func (o *InputSystemMetricsContainer) GetPerDevice() *bool {
	if o == nil {
		return nil
	}
	return o.PerDevice
}

func (o *InputSystemMetricsContainer) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

type InputSystemMetricsMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputSystemMetricsMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputSystemMetricsMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputSystemMetricsDataCompressionFormat string

const (
	InputSystemMetricsDataCompressionFormatNone InputSystemMetricsDataCompressionFormat = "none"
	InputSystemMetricsDataCompressionFormatGzip InputSystemMetricsDataCompressionFormat = "gzip"
)

func (e InputSystemMetricsDataCompressionFormat) ToPointer() *InputSystemMetricsDataCompressionFormat {
	return &e
}
func (e *InputSystemMetricsDataCompressionFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSystemMetricsDataCompressionFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemMetricsDataCompressionFormat: %v", v)
	}
}

type InputSystemMetricsPersistence struct {
	// Spool metrics to disk for Cribl Edge and Search
	Enable *bool `default:"false" json:"enable"`
	// Time span for each file bucket
	TimeWindow *string `default:"10m" json:"timeWindow"`
	// Maximum disk space allowed to be consumed (examples: 420MB, 4GB). When limit is reached, older data will be deleted.
	MaxDataSize *string `default:"1GB" json:"maxDataSize"`
	// Maximum amount of time to retain data (examples: 2h, 4d). When limit is reached, older data will be deleted.
	MaxDataTime *string                                  `default:"24h" json:"maxDataTime"`
	Compress    *InputSystemMetricsDataCompressionFormat `default:"gzip" json:"compress"`
	// Path to use to write metrics. Defaults to $CRIBL_HOME/state/system_metrics
	DestPath *string `default:"\\$CRIBL_HOME/state/system_metrics" json:"destPath"`
}

func (i InputSystemMetricsPersistence) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemMetricsPersistence) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemMetricsPersistence) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

func (o *InputSystemMetricsPersistence) GetTimeWindow() *string {
	if o == nil {
		return nil
	}
	return o.TimeWindow
}

func (o *InputSystemMetricsPersistence) GetMaxDataSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataSize
}

func (o *InputSystemMetricsPersistence) GetMaxDataTime() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataTime
}

func (o *InputSystemMetricsPersistence) GetCompress() *InputSystemMetricsDataCompressionFormat {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *InputSystemMetricsPersistence) GetDestPath() *string {
	if o == nil {
		return nil
	}
	return o.DestPath
}

type InputSystemMetrics struct {
	// Unique ID for this input
	ID       string                 `json:"id"`
	Type     InputSystemMetricsType `json:"type"`
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
	Connections []InputSystemMetricsConnections `json:"connections,omitempty"`
	Pq          *InputSystemMetricsPq           `json:"pq,omitempty"`
	// Time, in seconds, between consecutive metric collections. Default is 10 seconds.
	Interval  *float64                     `default:"10" json:"interval"`
	Host      *Host                        `json:"host,omitempty"`
	Process   *InputSystemMetricsProcess   `json:"process,omitempty"`
	Container *InputSystemMetricsContainer `json:"container,omitempty"`
	// Fields to add to events from this input
	Metadata    []InputSystemMetricsMetadata   `json:"metadata,omitempty"`
	Persistence *InputSystemMetricsPersistence `json:"persistence,omitempty"`
	Description *string                        `json:"description,omitempty"`
}

func (i InputSystemMetrics) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemMetrics) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemMetrics) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InputSystemMetrics) GetType() InputSystemMetricsType {
	if o == nil {
		return InputSystemMetricsType("")
	}
	return o.Type
}

func (o *InputSystemMetrics) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputSystemMetrics) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSystemMetrics) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputSystemMetrics) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputSystemMetrics) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputSystemMetrics) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputSystemMetrics) GetConnections() []InputSystemMetricsConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputSystemMetrics) GetPq() *InputSystemMetricsPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputSystemMetrics) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *InputSystemMetrics) GetHost() *Host {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputSystemMetrics) GetProcess() *InputSystemMetricsProcess {
	if o == nil {
		return nil
	}
	return o.Process
}

func (o *InputSystemMetrics) GetContainer() *InputSystemMetricsContainer {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *InputSystemMetrics) GetMetadata() []InputSystemMetricsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputSystemMetrics) GetPersistence() *InputSystemMetricsPersistence {
	if o == nil {
		return nil
	}
	return o.Persistence
}

func (o *InputSystemMetrics) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
