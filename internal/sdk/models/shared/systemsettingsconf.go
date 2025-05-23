// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Headers struct {
}

type Ssl struct {
	CaPath      *string `json:"caPath,omitempty"`
	CertPath    string  `json:"certPath"`
	Disabled    bool    `json:"disabled"`
	Passphrase  string  `json:"passphrase"`
	PrivKeyPath string  `json:"privKeyPath"`
}

func (o *Ssl) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *Ssl) GetCertPath() string {
	if o == nil {
		return ""
	}
	return o.CertPath
}

func (o *Ssl) GetDisabled() bool {
	if o == nil {
		return false
	}
	return o.Disabled
}

func (o *Ssl) GetPassphrase() string {
	if o == nil {
		return ""
	}
	return o.Passphrase
}

func (o *Ssl) GetPrivKeyPath() string {
	if o == nil {
		return ""
	}
	return o.PrivKeyPath
}

type API struct {
	BaseURL            *string  `json:"baseUrl,omitempty"`
	DisableAPICache    *bool    `json:"disableApiCache,omitempty"`
	Disabled           bool     `json:"disabled"`
	Headers            *Headers `json:"headers,omitempty"`
	Host               string   `json:"host"`
	IdleSessionTTL     *float64 `json:"idleSessionTTL,omitempty"`
	ListenOnPort       *bool    `json:"listenOnPort,omitempty"`
	LoginRateLimit     *string  `json:"loginRateLimit,omitempty"`
	Port               float64  `json:"port"`
	Protocol           string   `json:"protocol"`
	Scripts            *bool    `json:"scripts,omitempty"`
	SensitiveFields    []string `json:"sensitiveFields,omitempty"`
	Ssl                Ssl      `json:"ssl"`
	SsoRateLimit       *string  `json:"ssoRateLimit,omitempty"`
	WorkerRemoteAccess bool     `json:"workerRemoteAccess"`
}

func (o *API) GetBaseURL() *string {
	if o == nil {
		return nil
	}
	return o.BaseURL
}

func (o *API) GetDisableAPICache() *bool {
	if o == nil {
		return nil
	}
	return o.DisableAPICache
}

func (o *API) GetDisabled() bool {
	if o == nil {
		return false
	}
	return o.Disabled
}

func (o *API) GetHeaders() *Headers {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *API) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *API) GetIdleSessionTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.IdleSessionTTL
}

func (o *API) GetListenOnPort() *bool {
	if o == nil {
		return nil
	}
	return o.ListenOnPort
}

func (o *API) GetLoginRateLimit() *string {
	if o == nil {
		return nil
	}
	return o.LoginRateLimit
}

func (o *API) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *API) GetProtocol() string {
	if o == nil {
		return ""
	}
	return o.Protocol
}

func (o *API) GetScripts() *bool {
	if o == nil {
		return nil
	}
	return o.Scripts
}

func (o *API) GetSensitiveFields() []string {
	if o == nil {
		return nil
	}
	return o.SensitiveFields
}

func (o *API) GetSsl() Ssl {
	if o == nil {
		return Ssl{}
	}
	return o.Ssl
}

func (o *API) GetSsoRateLimit() *string {
	if o == nil {
		return nil
	}
	return o.SsoRateLimit
}

func (o *API) GetWorkerRemoteAccess() bool {
	if o == nil {
		return false
	}
	return o.WorkerRemoteAccess
}

type Backups struct {
	BackupPersistence string `json:"backupPersistence"`
	BackupsDirectory  string `json:"backupsDirectory"`
}

func (o *Backups) GetBackupPersistence() string {
	if o == nil {
		return ""
	}
	return o.BackupPersistence
}

func (o *Backups) GetBackupsDirectory() string {
	if o == nil {
		return ""
	}
	return o.BackupsDirectory
}

type CustomLogo struct {
	Enabled         bool   `json:"enabled"`
	LogoDescription string `json:"logoDescription"`
	LogoImage       string `json:"logoImage"`
}

func (o *CustomLogo) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *CustomLogo) GetLogoDescription() string {
	if o == nil {
		return ""
	}
	return o.LogoDescription
}

func (o *CustomLogo) GetLogoImage() string {
	if o == nil {
		return ""
	}
	return o.LogoImage
}

type Pii struct {
	EnablePiiDetection bool `json:"enablePiiDetection"`
}

func (o *Pii) GetEnablePiiDetection() bool {
	if o == nil {
		return false
	}
	return o.EnablePiiDetection
}

type Proxy struct {
	UseEnvVars bool `json:"useEnvVars"`
}

func (o *Proxy) GetUseEnvVars() bool {
	if o == nil {
		return false
	}
	return o.UseEnvVars
}

type Rollback struct {
	RollbackEnabled bool     `json:"rollbackEnabled"`
	RollbackRetries *float64 `json:"rollbackRetries,omitempty"`
	RollbackTimeout *float64 `json:"rollbackTimeout,omitempty"`
}

func (o *Rollback) GetRollbackEnabled() bool {
	if o == nil {
		return false
	}
	return o.RollbackEnabled
}

func (o *Rollback) GetRollbackRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.RollbackRetries
}

func (o *Rollback) GetRollbackTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RollbackTimeout
}

type Shutdown struct {
	DrainTimeout float64 `json:"drainTimeout"`
}

func (o *Shutdown) GetDrainTimeout() float64 {
	if o == nil {
		return 0.0
	}
	return o.DrainTimeout
}

type Sni struct {
	DisableSNIRouting bool `json:"disableSNIRouting"`
}

func (o *Sni) GetDisableSNIRouting() bool {
	if o == nil {
		return false
	}
	return o.DisableSNIRouting
}

type Sockets struct {
	Directory *string `json:"directory,omitempty"`
}

func (o *Sockets) GetDirectory() *string {
	if o == nil {
		return nil
	}
	return o.Directory
}

type Upgrade string

const (
	UpgradeFalse Upgrade = "false"
	UpgradeAPI   Upgrade = "api"
)

func (e Upgrade) ToPointer() *Upgrade {
	return &e
}
func (e *Upgrade) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "false":
		fallthrough
	case "api":
		*e = Upgrade(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Upgrade: %v", v)
	}
}

type System struct {
	Intercom bool    `json:"intercom"`
	Upgrade  Upgrade `json:"upgrade"`
}

func (o *System) GetIntercom() bool {
	if o == nil {
		return false
	}
	return o.Intercom
}

func (o *System) GetUpgrade() Upgrade {
	if o == nil {
		return Upgrade("")
	}
	return o.Upgrade
}

type SystemSettingsConfTLS struct {
	DefaultCipherList  string `json:"defaultCipherList"`
	DefaultEcdhCurve   string `json:"defaultEcdhCurve"`
	MaxVersion         string `json:"maxVersion"`
	MinVersion         string `json:"minVersion"`
	RejectUnauthorized bool   `json:"rejectUnauthorized"`
}

func (o *SystemSettingsConfTLS) GetDefaultCipherList() string {
	if o == nil {
		return ""
	}
	return o.DefaultCipherList
}

func (o *SystemSettingsConfTLS) GetDefaultEcdhCurve() string {
	if o == nil {
		return ""
	}
	return o.DefaultEcdhCurve
}

func (o *SystemSettingsConfTLS) GetMaxVersion() string {
	if o == nil {
		return ""
	}
	return o.MaxVersion
}

func (o *SystemSettingsConfTLS) GetMinVersion() string {
	if o == nil {
		return ""
	}
	return o.MinVersion
}

func (o *SystemSettingsConfTLS) GetRejectUnauthorized() bool {
	if o == nil {
		return false
	}
	return o.RejectUnauthorized
}

type SystemSettingsConfWorkers struct {
	Count                  float64  `json:"count"`
	EnableHeapSnapshots    *bool    `json:"enableHeapSnapshots,omitempty"`
	LoadThrottlePerc       *float64 `json:"loadThrottlePerc,omitempty"`
	Memory                 float64  `json:"memory"`
	Minimum                float64  `json:"minimum"`
	StartupMaxConns        *float64 `json:"startupMaxConns,omitempty"`
	StartupThrottleTimeout *float64 `json:"startupThrottleTimeout,omitempty"`
	V8SingleThread         *bool    `json:"v8SingleThread,omitempty"`
}

func (o *SystemSettingsConfWorkers) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *SystemSettingsConfWorkers) GetEnableHeapSnapshots() *bool {
	if o == nil {
		return nil
	}
	return o.EnableHeapSnapshots
}

func (o *SystemSettingsConfWorkers) GetLoadThrottlePerc() *float64 {
	if o == nil {
		return nil
	}
	return o.LoadThrottlePerc
}

func (o *SystemSettingsConfWorkers) GetMemory() float64 {
	if o == nil {
		return 0.0
	}
	return o.Memory
}

func (o *SystemSettingsConfWorkers) GetMinimum() float64 {
	if o == nil {
		return 0.0
	}
	return o.Minimum
}

func (o *SystemSettingsConfWorkers) GetStartupMaxConns() *float64 {
	if o == nil {
		return nil
	}
	return o.StartupMaxConns
}

func (o *SystemSettingsConfWorkers) GetStartupThrottleTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.StartupThrottleTimeout
}

func (o *SystemSettingsConfWorkers) GetV8SingleThread() *bool {
	if o == nil {
		return nil
	}
	return o.V8SingleThread
}

type SystemSettingsConf struct {
	API                  API                       `json:"api"`
	Backups              Backups                   `json:"backups"`
	CustomLogo           CustomLogo                `json:"customLogo"`
	Pii                  Pii                       `json:"pii"`
	Proxy                Proxy                     `json:"proxy"`
	Rollback             Rollback                  `json:"rollback"`
	Shutdown             Shutdown                  `json:"shutdown"`
	Sni                  Sni                       `json:"sni"`
	Sockets              *Sockets                  `json:"sockets,omitempty"`
	System               System                    `json:"system"`
	TLS                  SystemSettingsConfTLS     `json:"tls"`
	UpgradeGroupSettings UpgradeGroupSettings      `json:"upgradeGroupSettings"`
	UpgradeSettings      UpgradeSettings           `json:"upgradeSettings"`
	Workers              SystemSettingsConfWorkers `json:"workers"`
}

func (o *SystemSettingsConf) GetAPI() API {
	if o == nil {
		return API{}
	}
	return o.API
}

func (o *SystemSettingsConf) GetBackups() Backups {
	if o == nil {
		return Backups{}
	}
	return o.Backups
}

func (o *SystemSettingsConf) GetCustomLogo() CustomLogo {
	if o == nil {
		return CustomLogo{}
	}
	return o.CustomLogo
}

func (o *SystemSettingsConf) GetPii() Pii {
	if o == nil {
		return Pii{}
	}
	return o.Pii
}

func (o *SystemSettingsConf) GetProxy() Proxy {
	if o == nil {
		return Proxy{}
	}
	return o.Proxy
}

func (o *SystemSettingsConf) GetRollback() Rollback {
	if o == nil {
		return Rollback{}
	}
	return o.Rollback
}

func (o *SystemSettingsConf) GetShutdown() Shutdown {
	if o == nil {
		return Shutdown{}
	}
	return o.Shutdown
}

func (o *SystemSettingsConf) GetSni() Sni {
	if o == nil {
		return Sni{}
	}
	return o.Sni
}

func (o *SystemSettingsConf) GetSockets() *Sockets {
	if o == nil {
		return nil
	}
	return o.Sockets
}

func (o *SystemSettingsConf) GetSystem() System {
	if o == nil {
		return System{}
	}
	return o.System
}

func (o *SystemSettingsConf) GetTLS() SystemSettingsConfTLS {
	if o == nil {
		return SystemSettingsConfTLS{}
	}
	return o.TLS
}

func (o *SystemSettingsConf) GetUpgradeGroupSettings() UpgradeGroupSettings {
	if o == nil {
		return UpgradeGroupSettings{}
	}
	return o.UpgradeGroupSettings
}

func (o *SystemSettingsConf) GetUpgradeSettings() UpgradeSettings {
	if o == nil {
		return UpgradeSettings{}
	}
	return o.UpgradeSettings
}

func (o *SystemSettingsConf) GetWorkers() SystemSettingsConfWorkers {
	if o == nil {
		return SystemSettingsConfWorkers{}
	}
	return o.Workers
}
