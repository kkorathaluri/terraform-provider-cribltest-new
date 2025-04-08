// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

var GetCreditsServerList = []string{
	"https://api.cribl-staging.cloud",
}

type GetCreditsRequest struct {
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *GetCreditsRequest) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

type GetCreditsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse   *http.Response
	GetCreditsDTO *shared.GetCreditsDTO
}

func (o *GetCreditsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreditsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreditsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCreditsResponse) GetGetCreditsDTO() *shared.GetCreditsDTO {
	if o == nil {
		return nil
	}
	return o.GetCreditsDTO
}
