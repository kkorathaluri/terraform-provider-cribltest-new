// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// CreatePacksResponseBody - a list of PackInstallInfo objects
type CreatePacksResponseBody struct {
	// number of items present in the items array
	CountTest *int64                   `json:"count_test,omitempty"`
	Items     []shared.PackInstallInfo `json:"items,omitempty"`
}

func (o *CreatePacksResponseBody) GetCountTest() *int64 {
	if o == nil {
		return nil
	}
	return o.CountTest
}

func (o *CreatePacksResponseBody) GetItems() []shared.PackInstallInfo {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreatePacksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of PackInstallInfo objects
	Object *CreatePacksResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *CreatePacksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePacksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePacksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePacksResponse) GetObject() *CreatePacksResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *CreatePacksResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
