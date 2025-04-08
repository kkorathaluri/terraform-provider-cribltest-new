// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type SavedJobScheduledSearchJobType string

const (
	SavedJobScheduledSearchJobTypeCollection      SavedJobScheduledSearchJobType = "collection"
	SavedJobScheduledSearchJobTypeExecutor        SavedJobScheduledSearchJobType = "executor"
	SavedJobScheduledSearchJobTypeScheduledSearch SavedJobScheduledSearchJobType = "scheduledSearch"
)

func (e SavedJobScheduledSearchJobType) ToPointer() *SavedJobScheduledSearchJobType {
	return &e
}
func (e *SavedJobScheduledSearchJobType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "collection":
		fallthrough
	case "executor":
		fallthrough
	case "scheduledSearch":
		*e = SavedJobScheduledSearchJobType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobScheduledSearchJobType: %v", v)
	}
}

type SavedJobScheduledSearchType string

const (
	SavedJobScheduledSearchTypeCollection SavedJobScheduledSearchType = "collection"
)

func (e SavedJobScheduledSearchType) ToPointer() *SavedJobScheduledSearchType {
	return &e
}
func (e *SavedJobScheduledSearchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "collection":
		*e = SavedJobScheduledSearchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobScheduledSearchType: %v", v)
	}
}

// SavedJobScheduledSearchLogLevel - Level at which to set task logging
type SavedJobScheduledSearchLogLevel string

const (
	SavedJobScheduledSearchLogLevelError SavedJobScheduledSearchLogLevel = "error"
	SavedJobScheduledSearchLogLevelWarn  SavedJobScheduledSearchLogLevel = "warn"
	SavedJobScheduledSearchLogLevelInfo  SavedJobScheduledSearchLogLevel = "info"
	SavedJobScheduledSearchLogLevelDebug SavedJobScheduledSearchLogLevel = "debug"
	SavedJobScheduledSearchLogLevelSilly SavedJobScheduledSearchLogLevel = "silly"
)

func (e SavedJobScheduledSearchLogLevel) ToPointer() *SavedJobScheduledSearchLogLevel {
	return &e
}
func (e *SavedJobScheduledSearchLogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "warn":
		fallthrough
	case "info":
		fallthrough
	case "debug":
		fallthrough
	case "silly":
		*e = SavedJobScheduledSearchLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedJobScheduledSearchLogLevel: %v", v)
	}
}

type SavedJobScheduledSearchTimeWarning struct {
}

type SavedJobScheduledSearchRunSettings struct {
	Type *SavedJobScheduledSearchType `json:"type,omitempty"`
	// Reschedule tasks that failed with non-fatal errors
	RescheduleDroppedTasks *bool `default:"true" json:"rescheduleDroppedTasks"`
	// Maximum number of times a task can be rescheduled
	MaxTaskReschedule *float64 `default:"1" json:"maxTaskReschedule"`
	// Level at which to set task logging
	LogLevel *SavedJobScheduledSearchLogLevel `default:"info" json:"logLevel"`
	// Maximum time the job is allowed to run. Time unit defaults to seconds if not specified (examples: 30, 45s, 15m). Enter 0 for unlimited time.
	JobTimeout *string `default:"0" json:"jobTimeout"`
	// Job run mode. Preview will either return up to N matching results, or will run until capture time T is reached. Discovery will gather the list of files to turn into streaming tasks, without running the data collection job. Full Run will run the collection job.
	Mode          *string `default:"list" json:"mode"`
	TimeRangeType *string `default:"relative" json:"timeRangeType"`
	// Earliest time to collect data for the selected timezone
	Earliest *float64 `json:"earliest,omitempty"`
	// Latest time to collect data for the selected timezone
	Latest            *float64                            `json:"latest,omitempty"`
	TimestampTimezone any                                 `json:"timestampTimezone,omitempty"`
	TimeWarning       *SavedJobScheduledSearchTimeWarning `json:"timeWarning,omitempty"`
	// A filter for tokens in the provided collect path and/or the events being collected
	Expression *string `default:"true" json:"expression"`
	// Limits the bundle size for small tasks. For example,
	//
	//
	//
	//
	//
	//         if your lower bundle size is 1MB, you can bundle up to five 200KB files into one task.
	MinTaskSize *string `default:"1MB" json:"minTaskSize"`
	// Limits the bundle size for files above the lower task bundle size. For example, if your upper bundle size is 10MB,
	//
	//
	//
	//
	//
	//         you can bundle up to five 2MB files into one task. Files greater than this size will be assigned to individual tasks.
	MaxTaskSize *string `default:"10MB" json:"maxTaskSize"`
}

func (s SavedJobScheduledSearchRunSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SavedJobScheduledSearchRunSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SavedJobScheduledSearchRunSettings) GetType() *SavedJobScheduledSearchType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SavedJobScheduledSearchRunSettings) GetRescheduleDroppedTasks() *bool {
	if o == nil {
		return nil
	}
	return o.RescheduleDroppedTasks
}

func (o *SavedJobScheduledSearchRunSettings) GetMaxTaskReschedule() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxTaskReschedule
}

func (o *SavedJobScheduledSearchRunSettings) GetLogLevel() *SavedJobScheduledSearchLogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *SavedJobScheduledSearchRunSettings) GetJobTimeout() *string {
	if o == nil {
		return nil
	}
	return o.JobTimeout
}

func (o *SavedJobScheduledSearchRunSettings) GetMode() *string {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *SavedJobScheduledSearchRunSettings) GetTimeRangeType() *string {
	if o == nil {
		return nil
	}
	return o.TimeRangeType
}

func (o *SavedJobScheduledSearchRunSettings) GetEarliest() *float64 {
	if o == nil {
		return nil
	}
	return o.Earliest
}

func (o *SavedJobScheduledSearchRunSettings) GetLatest() *float64 {
	if o == nil {
		return nil
	}
	return o.Latest
}

func (o *SavedJobScheduledSearchRunSettings) GetTimestampTimezone() any {
	if o == nil {
		return nil
	}
	return o.TimestampTimezone
}

func (o *SavedJobScheduledSearchRunSettings) GetTimeWarning() *SavedJobScheduledSearchTimeWarning {
	if o == nil {
		return nil
	}
	return o.TimeWarning
}

func (o *SavedJobScheduledSearchRunSettings) GetExpression() *string {
	if o == nil {
		return nil
	}
	return o.Expression
}

func (o *SavedJobScheduledSearchRunSettings) GetMinTaskSize() *string {
	if o == nil {
		return nil
	}
	return o.MinTaskSize
}

func (o *SavedJobScheduledSearchRunSettings) GetMaxTaskSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxTaskSize
}

// SavedJobScheduledSearchSchedule - Configuration for a scheduled job
type SavedJobScheduledSearchSchedule struct {
	// Enable to configure scheduling for this Collector
	Enabled *bool `json:"enabled,omitempty"`
	// A cron schedule on which to run this job
	CronSchedule *string `default:"*/5 * * * *" json:"cronSchedule"`
	// The maximum number of instances of this scheduled job that may be running at any time
	MaxConcurrentRuns *float64 `default:"1" json:"maxConcurrentRuns"`
	// Skippable jobs can be delayed, up to their next run time, if the system is hitting concurrency limits
	Skippable    *bool                               `default:"true" json:"skippable"`
	ResumeMissed any                                 `json:"resumeMissed,omitempty"`
	Run          *SavedJobScheduledSearchRunSettings `json:"run,omitempty"`
}

func (s SavedJobScheduledSearchSchedule) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SavedJobScheduledSearchSchedule) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SavedJobScheduledSearchSchedule) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SavedJobScheduledSearchSchedule) GetCronSchedule() *string {
	if o == nil {
		return nil
	}
	return o.CronSchedule
}

func (o *SavedJobScheduledSearchSchedule) GetMaxConcurrentRuns() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxConcurrentRuns
}

func (o *SavedJobScheduledSearchSchedule) GetSkippable() *bool {
	if o == nil {
		return nil
	}
	return o.Skippable
}

func (o *SavedJobScheduledSearchSchedule) GetResumeMissed() any {
	if o == nil {
		return nil
	}
	return o.ResumeMissed
}

func (o *SavedJobScheduledSearchSchedule) GetRun() *SavedJobScheduledSearchRunSettings {
	if o == nil {
		return nil
	}
	return o.Run
}

type SavedJobScheduledSearch struct {
	// Unique ID for this Job
	ID          *string                        `json:"id,omitempty"`
	Description *string                        `json:"description,omitempty"`
	Type        SavedJobScheduledSearchJobType `json:"type"`
	// Time to keep the job's artifacts on disk after job completion. This also affects how long a job is listed in the Job Inspector.
	TTL *string `default:"4h" json:"ttl"`
	// List of fields to remove from Discover results. Wildcards (for example, aws*) are allowed. This is useful when discovery returns sensitive fields that should not be exposed in the Jobs user interface.
	RemoveFields []string `json:"removeFields,omitempty"`
	// Resume the ad hoc job if a failure condition causes Stream to restart during job execution
	ResumeOnBoot *bool `default:"false" json:"resumeOnBoot"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Configuration for a scheduled job
	Schedule *SavedJobScheduledSearchSchedule `json:"schedule,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Identifies which search query to run
	SavedQueryID string `json:"savedQueryId"`
}

func (s SavedJobScheduledSearch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SavedJobScheduledSearch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SavedJobScheduledSearch) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SavedJobScheduledSearch) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SavedJobScheduledSearch) GetType() SavedJobScheduledSearchJobType {
	if o == nil {
		return SavedJobScheduledSearchJobType("")
	}
	return o.Type
}

func (o *SavedJobScheduledSearch) GetTTL() *string {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *SavedJobScheduledSearch) GetRemoveFields() []string {
	if o == nil {
		return nil
	}
	return o.RemoveFields
}

func (o *SavedJobScheduledSearch) GetResumeOnBoot() *bool {
	if o == nil {
		return nil
	}
	return o.ResumeOnBoot
}

func (o *SavedJobScheduledSearch) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *SavedJobScheduledSearch) GetSchedule() *SavedJobScheduledSearchSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *SavedJobScheduledSearch) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *SavedJobScheduledSearch) GetSavedQueryID() string {
	if o == nil {
		return ""
	}
	return o.SavedQueryID
}
