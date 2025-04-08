// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputModelDrivenTelemetry struct {
	Connections       []InputModelDrivenTelemetryConnections          `tfsdk:"connections"`
	Description       types.String                                    `tfsdk:"description"`
	Disabled          types.Bool                                      `tfsdk:"disabled"`
	Environment       types.String                                    `tfsdk:"environment"`
	Host              types.String                                    `tfsdk:"host"`
	ID                types.String                                    `tfsdk:"id"`
	MaxActiveCxn      types.Float64                                   `tfsdk:"max_active_cxn"`
	Metadata          []InputModelDrivenTelemetryMetadata             `tfsdk:"metadata"`
	Pipeline          types.String                                    `tfsdk:"pipeline"`
	Port              types.Float64                                   `tfsdk:"port"`
	Pq                *InputModelDrivenTelemetryPq                    `tfsdk:"pq"`
	PqEnabled         types.Bool                                      `tfsdk:"pq_enabled"`
	SendToRoutes      types.Bool                                      `tfsdk:"send_to_routes"`
	ShutdownTimeoutMs types.Float64                                   `tfsdk:"shutdown_timeout_ms"`
	Streamtags        []types.String                                  `tfsdk:"streamtags"`
	TLS               *InputModelDrivenTelemetryTLSSettingsServerSide `tfsdk:"tls"`
	Type              types.String                                    `tfsdk:"type"`
}
