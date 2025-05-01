package hooks

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	// Default environment values
	defaultEnvironmentID = "cribl-playground.cloud"
	defaultWorkerGroupID = "default"

	// URL path constants
	workerGroupsPath = "/workerGroups"
	systemProjectsPath = "/system/projects"
)

// ResourceType represents the type of resource being accessed
type ResourceType string

const (
	// Resource types
	WorkerGroupResource ResourceType = "workerGroup"
	SystemResource      ResourceType = "system"
	ProjectResource     ResourceType = "project"
	OtherResource       ResourceType = "other"
)

// workerGroupResources defines paths that should be routed to worker group URLs
var workerGroupResources = []string{
	"/packs",
	"/routes",
	"/pipeline",
	"/functions",
	"/lookups",
	"/jobs",
	"/collectors",
	"/destinations",
	"/sources",
	"/system",
}

// CriblTerraformURLHook implements URL routing for Cribl Terraform API
type CriblTerraformURLHook struct {
	defaultBase *url.URL
	wgBase      *url.URL
}

var (
	_ sdkInitHook       = (*CriblTerraformURLHook)(nil)
	_ beforeRequestHook = (*CriblTerraformURLHook)(nil)
)

// getEnvironmentValue gets an environment variable with a default fallback
func getEnvironmentValue(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// constructURL builds a URL with the given components
func constructURL(envID, orgID, workspaceID, workerGroupID string) string {
	return fmt.Sprintf("https://app.%s/organizations/%s/workspaces/%s/app/api/v1/m/%s",
		envID,
		orgID,
		workspaceID,
		workerGroupID,
	)
}

// NewCriblTerraformURLHook creates a new URL hook for Cribl Terraform
func NewCriblTerraformURLHook() *CriblTerraformURLHook {
	envID := getEnvironmentValue("CRIBL_ENVIRONMENT_ID", defaultEnvironmentID)
	orgID := getEnvironmentValue("CRIBL_ORGANIZATION_ID", "")
	workspaceID := getEnvironmentValue("CRIBL_WORKSPACE_ID", "")
	workerGroupID := getEnvironmentValue("CRIBL_WORKER_GROUP_ID", defaultWorkerGroupID)

	// Construct the default URL
	defaultURL := constructURL(envID, orgID, workspaceID, workerGroupID)

	// Construct the worker group URL (without worker group ID)
	wgURL := fmt.Sprintf("https://app.%s/organizations/%s/workspaces/%s/app/api/v1",
		envID,
		orgID,
		workspaceID,
	)

	defaultBase, _ := url.Parse(defaultURL)
	wgBase, _ := url.Parse(wgURL)

	return &CriblTerraformURLHook{
		defaultBase: defaultBase,
		wgBase:      wgBase,
	}
}

// SDKInit implements the sdkInitHook interface
// It sets the base URL to the constructed default URL
func (o *CriblTerraformURLHook) SDKInit(baseURL string, client HTTPClient) (string, HTTPClient) {
	return o.defaultBase.String(), client
}

// getResourceType determines the type of resource being accessed
func getResourceType(path string) ResourceType {
	if strings.HasPrefix(path, workerGroupsPath) {
		return WorkerGroupResource
	}
	if strings.Contains(path, systemProjectsPath) {
		return ProjectResource
	}
	for _, resource := range workerGroupResources {
		if strings.HasPrefix(path, resource) {
			return SystemResource
		}
	}
	return OtherResource
}

// BeforeRequest implements the beforeRequestHook interface
// It modifies the request URL based on the path and resource type
func (o *CriblTerraformURLHook) BeforeRequest(ctx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	path := req.URL.Path
	resourceType := getResourceType(path)

	switch resourceType {
	case WorkerGroupResource:
		// For worker group operations, use the base API URL
		req.URL = o.wgBase.ResolveReference(req.URL)
	case ProjectResource, SystemResource:
		// For system/projects resources and worker group resources, use the worker group URL
		req.URL = o.defaultBase.ResolveReference(req.URL)
	default:
		// For all other paths, use the base API URL
		req.URL = o.wgBase.ResolveReference(req.URL)
	}

	return req, nil
}