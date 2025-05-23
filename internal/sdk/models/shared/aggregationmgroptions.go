// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AggregationMgrOptions struct {
	Aggregations             []string       `json:"aggregations"`
	Cumulative               bool           `json:"cumulative"`
	DotsAsLiterals           *bool          `json:"dotsAsLiterals,omitempty"`
	FlushEventLimit          float64        `json:"flushEventLimit"`
	FlushMemLimit            float64        `json:"flushMemLimit"`
	Hostname                 string         `json:"hostname"`
	IdleTimeLimitSeconds     *float64       `json:"idleTimeLimitSeconds,omitempty"`
	LagToleranceSeconds      *float64       `json:"lagToleranceSeconds,omitempty"`
	MetricsMode              bool           `json:"metricsMode"`
	Prefix                   *string        `json:"prefix,omitempty"`
	PreserveSplitByStructure *bool          `json:"preserveSplitByStructure,omitempty"`
	SearchAggMode            *SearchAggMode `json:"searchAggMode,omitempty"`
	SplitBys                 []string       `json:"splitBys,omitempty"`
	SufficientStatsOnly      bool           `json:"sufficientStatsOnly"`
	TimeWindowSeconds        float64        `json:"timeWindowSeconds"`
}

func (o *AggregationMgrOptions) GetAggregations() []string {
	if o == nil {
		return []string{}
	}
	return o.Aggregations
}

func (o *AggregationMgrOptions) GetCumulative() bool {
	if o == nil {
		return false
	}
	return o.Cumulative
}

func (o *AggregationMgrOptions) GetDotsAsLiterals() *bool {
	if o == nil {
		return nil
	}
	return o.DotsAsLiterals
}

func (o *AggregationMgrOptions) GetFlushEventLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.FlushEventLimit
}

func (o *AggregationMgrOptions) GetFlushMemLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.FlushMemLimit
}

func (o *AggregationMgrOptions) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *AggregationMgrOptions) GetIdleTimeLimitSeconds() *float64 {
	if o == nil {
		return nil
	}
	return o.IdleTimeLimitSeconds
}

func (o *AggregationMgrOptions) GetLagToleranceSeconds() *float64 {
	if o == nil {
		return nil
	}
	return o.LagToleranceSeconds
}

func (o *AggregationMgrOptions) GetMetricsMode() bool {
	if o == nil {
		return false
	}
	return o.MetricsMode
}

func (o *AggregationMgrOptions) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *AggregationMgrOptions) GetPreserveSplitByStructure() *bool {
	if o == nil {
		return nil
	}
	return o.PreserveSplitByStructure
}

func (o *AggregationMgrOptions) GetSearchAggMode() *SearchAggMode {
	if o == nil {
		return nil
	}
	return o.SearchAggMode
}

func (o *AggregationMgrOptions) GetSplitBys() []string {
	if o == nil {
		return nil
	}
	return o.SplitBys
}

func (o *AggregationMgrOptions) GetSufficientStatsOnly() bool {
	if o == nil {
		return false
	}
	return o.SufficientStatsOnly
}

func (o *AggregationMgrOptions) GetTimeWindowSeconds() float64 {
	if o == nil {
		return 0.0
	}
	return o.TimeWindowSeconds
}
